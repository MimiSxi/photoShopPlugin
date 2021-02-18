package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Device struct {
	ID           uint                 `gorm:"primary_key" gqlschema:"update!;delete!;query!;querys" description:"ID"`
	DeviceId     string               `gorm:"Type:varchar(64);DEFAULT:'';NOT NULL;" gqlschema:"create!;update;querys" description:"设备编号"`  // 设备编号
	Name         string               `gorm:"Type:varchar(64);DEFAULT:'';NOT NULL;" gqlschema:"create!;update;querys" description:"设备名称"`  // 设备名称
	Status       DeviceStatusEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"create!;update;querys" description:"设备状态"`                    // 设备状态
	BorrowStatus BorrowStatusEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"create;querys" description:"借用状态"`                            // 借用状态
	Type         string               `gorm:"Type:varchar(64);DEFAULT:'';NOT NULL;" gqlschema:"create!;update;querys" description:"设备型号"`  // 设备型号
	ProductName  string               `gorm:"Type:varchar(64);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"设备出厂编号"` // 设备出厂编号
	ProductType  string               `gorm:"Type:varchar(64);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"设备出厂型号"` // 设备出厂型号
	Annex        AnnexJSON            `gorm:"Type:text;" gqlschema:"create;update" description:"设备附件"`                                     // 设备附件
	Remark       string               `gorm:"Type:varchar(256);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"备注"`    // 备注
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	v2           int `gorm:"-" exclude:"true"`
}

type Devices struct {
	TotalCount int
	Edges      []Device
	// Data       int      `gqlschema:"fields:create!;update; func:'a'; description:'设备出厂型号'; exclude:'true'`
}

type AnnexJSON map[string]interface{}

func (c AnnexJSON) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *AnnexJSON) Scan(input interface{}) error {
	v, ok := input.([]byte)
	if !ok { //sqlite
		v = []byte(input.(string))
	}
	err := json.Unmarshal(v, c)
	return err
}

func (o *Device) BeforeSave(scope *gorm.Scope) (err error) {
	// dbx := scope.DB()

	// if o.EmployeeJobID > 0 {
	// 	if err := dbx.Where("id = ?", o.EmployeeJobID).First(&EmployeeJob{}).Error; err != nil {
	// 		return err
	// 	}
	// }
	return
}

func (o *Device) QueryByID(id uint) (err error) {
	return db.Where("id = ?", id).First(&o).Error
}

func (o Device) Create(params graphql.ResolveParams) (Device, error) {
	p := params.Args

	o.Name = p["name"].(string)
	o.Type = p["type"].(string)
	o.DeviceId = p["deviceId"].(string)

	if p["productName"] != nil {
		o.ProductName = p["productName"].(string)
	}
	if p["productType"] != nil {
		o.ProductType = p["productType"].(string)
	}

	if p["status"] != nil {
		o.Status = p["status"].(DeviceStatusEnumType)
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

	err := db.Transaction(func(tx *gorm.DB) error {
		// if params.Args["orgid"] != nil {
		// 	// check OrgID
		// 	if err := tx.Where("id = ?", o.orgID).First(&Org{}).Error; err != nil {
		// 		return err
		// 	}
		// }

		// create Device
		err := tx.Create(&o).Error
		if err != nil {
			return err
		}
		return nil
	})

	return o, err
}

func (o Device) Query(params graphql.ResolveParams) (Device, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o Device) Querys(params graphql.ResolveParams) (Devices, error) {
	var result Devices

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

func (o Device) Update(params graphql.ResolveParams) (Device, error) {
	v, ok := params.Source.(Device)
	if !ok {
		return o, errors.New("update param")
	}

	// applyform := ApplyForm{}
	// if err := db.Where("device_id = ?", v.ID).First(&applyform).Error; err != gorm.ErrRecordNotFound {
	// 	return v, errors.New("存在借用记录，不允许修改内容")
	// }

	p := params.Args
	if p["name"] != nil {
		v.Name = p["name"].(string)
	}
	if p["type"] != nil {
		v.Type = p["type"].(string)
	}
	if p["deviceId"] != nil {
		v.DeviceId = p["deviceId"].(string)
	}
	if p["productName"] != nil {
		v.ProductName = p["productName"].(string)
	}
	if p["productType"] != nil {
		v.ProductType = p["productType"].(string)
	}
	if p["remark"] != nil {
		v.Remark = p["remark"].(string)
	}
	if p["annex"] != nil {
		config := p["annex"].(string)
		if err := json.Unmarshal([]byte(config), &v.Annex); err != nil {
			return o, err
		}
	}

	if p["status"] != nil {
		if v.BorrowStatus != BorrowStatus_NO_APPLY {
			return v, errors.New("申领状态时，不允许修改状态")
		}
		v.Status = p["status"].(DeviceStatusEnumType)
	}

	err := db.Save(&v).Error
	return v, err
}

func (o Device) Delete(params graphql.ResolveParams) (Device, error) {
	v, ok := params.Source.(Device)
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
