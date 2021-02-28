/**
 * @Author zhangfan
 * @create 2021/2/27 下午2:58
 * Description: 设计器项目
 */

package model

import (
	"errors"
	"github.com/graphql-go/graphql"
	"time"
)

type ProJ struct {
	ID        uint                     `gorm:"primary_key" gqlschema:"update!;delete!;query!;querys" description:"设计器项目id"`
	UserId    uint                     `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create!;querys" description:"创建用户id" funservice:"employee"`
	Status    ProJCommonStatusEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"update;querys" description:"状态"`
	Name      string                   `gorm:"Type:varchar(64);DEFAULT:'';NOT NULL;" gqlschema:"create!;update;querys" description:"项目名称"`
	Cover     string                   `gorm:"Type:varchar(64);DEFAULT:'';NOT NULL;" gqlschema:"create!;update;querys" description:"封面"`
	CreatedAt time.Time                `description:"创建时间" gqlschema:"querys"`
	UpdatedAt time.Time                `description:"更新时间" gqlschema:"querys"`
	DeletedAt *time.Time
	v2        int `gorm:"-" exclude:"true"`
	//pages:[Page] 画布
}

// 设计器项目集合
type ProJs struct {
	TotalCount int
	Edges      []ProJ
}

func (o ProJ) Query(params graphql.ResolveParams) (ProJ, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o ProJ) Querys(params graphql.ResolveParams) (ProJs, error) {
	var result ProJs

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

func (o ProJ) Create(params graphql.ResolveParams) (ProJ, error) {
	p := params.Args
	if p["name"] != nil {
		o.Name = p["name"].(string)
	}
	//todo pages:[Page]!
	o.UserId = p["userId"].(uint)
	o.Cover = p["cover"].(string)
	err := db.Create(&o).Error
	return o, err
}

func (o ProJ) Update(params graphql.ResolveParams) (ProJ, error) {
	v, ok := params.Source.(ProJ)
	if !ok {
		return o, errors.New("update param")
	}
	p := params.Args
	// todo pages:[Page]
	if p["status"] != nil {
		v.Status = p["status"].(ProJCommonStatusEnumType)
	}
	if p["name"] != nil {
		v.Name = p["name"].(string)
	}
	if p["cover"] != nil {
		v.Cover = p["cover"].(string)
	}
	err := db.Save(&v).Error
	return v, err
}

func (o ProJ) Delete(params graphql.ResolveParams) (ProJ, error) {
	v, ok := params.Source.(ProJ)
	if !ok {
		return o, errors.New("delete param")
	}
	err := db.Delete(&v).Error
	return v, err
}
