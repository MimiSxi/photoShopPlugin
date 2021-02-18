package model

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type ApplyForm struct {
	ID            uint                    `gorm:"primary_key" gqlschema:"giveoutdevice!;returndevice!;canceldevice!;query!;"`
	DeviceID      uint                    `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create!;querys" description:"设备编号" exclude:"true" funservice:"device"`   // 设备ID
	EmployeeID    uint                    `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create!;querys" description:"职员编号" exclude:"true" funservice:"employee"` // 职员ID
	Status        ApplyFormStatusEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"querys" description:"借用状态"`                                              // 借用状态
	BorrowMan     string                  `gorm:"Type:varchar(64);DEFAULT:'';NOT NULL;" gqlschema:"create;querys" description:"借用人"`                      // 借用人
	Annex         AnnexJSON               `gorm:"Type:text;" gqlschema:"create" description:"借用设备附件"`                                                     // 借用设备附件
	Remark        string                  `gorm:"Type:varchar(256);DEFAULT:'';NOT NULL;" gqlschema:"create;querys" description:"备注"`                      // 备注
	BorrowedAt    *time.Time              `description:"领用时间" gqlschema:"querys"`                                                                         // 领用时间
	DistributerID uint                    `gorm:"gorm:"DEFAULT:0;NOT NULL;" description:"发放人" exclude:"true" funservice:"employee"`                       // 发放人
	ReturnAt      *time.Time              `description:"归还时间" gqlschema:"querys"`                                                                         // 归还时间
	CreatedAt     time.Time               `description:"申请时间" gqlschema:"querys"`                                                                         // 归还时间
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	v2            int `gorm:"-" exclude:"true"`
}

type ApplyForms struct {
	TotalCount int
	Edges      []ApplyForm
}

func (o *ApplyForm) BeforeSave(scope *gorm.Scope) (err error) {
	// dbx := scope.DB()
	// if o.EmployeeJobID > 0 {
	// 	if err := dbx.Where("id = ?", o.EmployeeJobID).First(&EmployeeJob{}).Error; err != nil {
	// 		return err
	// 	}
	// }
	return
}

func (o *ApplyForm) QueryByID(id uint) (err error) {
	return db.Where("id = ?", id).First(&o).Error
}

func (o ApplyForm) Create(params graphql.ResolveParams) (ApplyForm, error) {
	p := params.Args

	o.Status = ApplyFormStatus_APPLIED
	o.DeviceID = p["deviceID"].(uint)
	o.EmployeeID = p["employeeID"].(uint)

	if p["borrowMan"] != nil {
		o.BorrowMan = p["borrowMan"].(string)
	}

	if p["remark"] != nil {
		o.Remark = p["remark"].(string)
	}

	if p["annex"] != nil {
		config := p["annex"].(string)
		if err := json.Unmarshal([]byte(config), &o.Annex); err != nil {
			return o, err
		}
	}
	o.BorrowedAt = nil
	o.ReturnAt = nil

	err := db.Transaction(func(tx *gorm.DB) error {
		//1. 检查设备状态是否正常
		//2. 检查设备状态是否可以被借用，是否已经被借用
		//3. 修改设备状态
		var device Device
		if err := tx.Where("id = ?", o.DeviceID).First(&device).Error; err != nil {
			return err
		}

		if device.Status != DeviceStatus_NORMAL {
			return errors.New("device status is not NORMAL")
		}

		if device.BorrowStatus != BorrowStatus_NO_APPLY {
			return errors.New("device BorrowStatus is not NO_APPLY")
		}

		device.BorrowStatus = BorrowStatus_TOBE_BORROWED
		tx.Save(&device)

		// create ApplyForm
		return tx.Create(&o).Error
	})

	return o, err
}

func (o ApplyForm) Query(params graphql.ResolveParams) (ApplyForm, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o ApplyForm) Querys(params graphql.ResolveParams) (ApplyForms, error) {
	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	if params.Source != nil {
		if obj, ok := params.Source.(Device); ok {
			dbselect = dbselect.Where("device_id = ?", obj.ID)
			dbcount = dbcount.Where("device_id = ?", obj.ID)
		} else {
			v := reflect.ValueOf(params.Source)
			if v.Type().Name() == "Employee" {
				var id uint
				{
					idx := v.FieldByName("ID")
					if !idx.IsValid() {
						panic("bad field in gorm.Model id")
					}
					id = uint(idx.Uint())
				}
				dbselect = dbselect.Where("employee_id = ?", id)
				dbcount = dbcount.Where("employee_id = ?", id)
			}
		}
	}

	var result ApplyForms
	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}

	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

//发放
func (o ApplyForm) Giveoutdevice(params graphql.ResolveParams) (ApplyForm, error) {
	v, ok := params.Source.(ApplyForm)
	if !ok {
		return o, errors.New("update param")
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		//1. 检查记录状态
		//2. 修改设备状态

		if v.Status != ApplyFormStatus_APPLIED {
			return errors.New("applyform status is not APPLIED")
		}

		var device Device
		if err := tx.Where("id = ?", v.DeviceID).First(&device).Error; err != nil {
			return err
		}
		if device.BorrowStatus != BorrowStatus_TOBE_BORROWED {
			return errors.New("device status is not TOBE_BORROWED")
		}
		device.BorrowStatus = BorrowStatus_TOBE_RETURNED
		tx.Save(&device)

		v.Status = ApplyFormStatus_BORROWED
		v.BorrowedAt = &[]time.Time{time.Now()}[0]
		//@todo 记录发放人

		{
			rootValue := params.Info.RootValue.(map[string]interface{})
			if rootValue["id"] == nil {
				return errors.New("user is not login")
			}
			userId, ok := rootValue["id"].(uint)
			if !ok {
				return errors.New("获取登陆用户失败")
			}
			v.DistributerID = userId
		}

		return tx.Save(&v).Error
	})

	return v, err
}

//归还
func (o ApplyForm) Returndevice(params graphql.ResolveParams) (ApplyForm, error) {
	v, ok := params.Source.(ApplyForm)
	if !ok {
		return o, errors.New("update param")
	}

	//
	err := db.Transaction(func(tx *gorm.DB) error {
		//1. 检查记录状态
		//2. 修改设备状态
		if v.Status != ApplyFormStatus_BORROWED {
			return errors.New("applyform status is not BORROWED")
		}

		var device Device
		if err := tx.Where("id = ?", v.DeviceID).First(&device).Error; err != nil {
			return err
		}
		if device.BorrowStatus != BorrowStatus_TOBE_RETURNED {
			return errors.New("device status is not TOBE_RETURNED")
		}
		device.BorrowStatus = BorrowStatus_NO_APPLY
		tx.Save(&device)

		//@todo 检查记录状态
		v.Status = ApplyFormStatus_RETURNED
		v.ReturnAt = &[]time.Time{time.Now()}[0]
		return tx.Save(&v).Error
	})

	return v, err
}

//取消
func (o ApplyForm) Canceldevice(params graphql.ResolveParams) (ApplyForm, error) {
	v, ok := params.Source.(ApplyForm)
	if !ok {
		return o, errors.New("update param")
	}

	//
	err := db.Transaction(func(tx *gorm.DB) error {
		//1. 检查记录状态
		//2. 修改设备状态
		if v.Status != ApplyFormStatus_APPLIED {
			return errors.New("applyform status is not APPLIED")
		}

		var device Device
		if err := tx.Where("id = ?", v.DeviceID).First(&device).Error; err != nil {
			return err
		}
		if device.BorrowStatus != BorrowStatus_TOBE_BORROWED {
			return errors.New("device status is not TOBE_BORROWED")
		}
		device.BorrowStatus = BorrowStatus_NO_APPLY
		tx.Save(&device)

		//@todo 检查记录状态
		v.Status = ApplyFormStatus_CANCEL
		return tx.Save(&v).Error
	})

	return v, err
}

/*
func (o ApplyForm) Update(params graphql.ResolveParams) (ApplyForm, error) {
	v, ok := params.Source.(ApplyForm)
	if !ok {
		return o, errors.New("update param")
	}

	p := params.Args
	// id := p["id"].(uint)
	// err := db.Where("id = ?", id).First(&o).Error
	// if err != nil {
	// 	return o, err
	// }
	// delete(p, "id")

	v.Name = p["name"].(string)
	v.Type = p["type"].(string)

	if p["deviceid"] != nil {
		v.DeviceId = p["deviceid"].(string)
	}

	if p["productname"] != nil {
		v.ProductName = p["productname"].(string)
	}
	if p["producttype"] != nil {
		v.ProductType = p["producttype"].(string)
	}
	if p["remark"] != nil {
		v.Remark = p["remark"].(string)
	}
	if p["annex"] != nil {
		config := p["annex"].(string)
		var annex AnnexJSON
		if err := json.Unmarshal([]byte(config), &annex); err != nil {
			return o, err
		}
		p["annex"] = annex
	}
	v.Status = p["status"].(DeviceStatusEnumType)

	err := db.Save(&v).Error
	return v, err
}

func (o ApplyForm) Delete(params graphql.ResolveParams) (ApplyForm, error) {
	v, ok := params.Source.(ApplyForm)
	if !ok {
		return o, errors.New("update param")
	}
	err := db.Delete(&v).Error
	return v, err

	// p := params.Args
	// err := db.Where(p).First(&o).Error
	// if err != nil {
	// 	return o, err
	// }
	// //
	// err = db.Transaction(func(tx *gorm.DB) error {
	// 	// if err := db.Delete(OrgEmployee{}, "employee_id = ?", o.ID).Error; err != nil {
	// 	// 	return err
	// 	// }
	// 	return db.Delete(&o).Error
	// })
	// return o, err
}
*/
