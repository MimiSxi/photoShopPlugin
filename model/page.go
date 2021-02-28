/**
 * @Author zhangfan
 * @create 2021/2/27 下午2:59
 * Description:
 */

package model

import (
	"errors"
	"github.com/graphql-go/graphql"
	"time"
)

type Page struct {
	ID        uint                     `gorm:"primary_key" gqlschema:"update!;delete!;query!;querys" description:"单页画布设计id"`
	RenderRes string                   `gorm:"Type:varchar(1000);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"画布的base64格式图片渲染图"`
	Status    ProJCommonStatusEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"update;querys" description:"模板页面状态"`
	Direction PageDirectionEnumType    `gorm:"DEFAULT:2;NOT NULL;" gqlschema:"create;update;querys" description:"画布页面放置方向"`
	PType     PageTypeEnumType         `gorm:"DEFAULT:4;NOT NULL;" gqlschema:"create;update;querys" description:"画布页面类型"`
	CreatedAt time.Time                `description:"创建时间" gqlschema:"querys"`
	UpdatedAt time.Time                `description:"更新时间" gqlschema:"querys"`
	DeletedAt *time.Time
	v2        int `gorm:"-" exclude:"true"`
	//canvasJson:String 画布json数据
	//font:String 该画布包含的字体json数据
}

type Pages struct {
	TotalCount int
	Edges      []Page
}

func (o Page) Query(params graphql.ResolveParams) (Page, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o Page) Querys(params graphql.ResolveParams) (Pages, error) {
	var result Pages

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

func (o Page) Create(params graphql.ResolveParams) (Page, error) {
	//		todo canvasJson:String 画布json数据
	//​		todo font:String 该画布包含的字体json数据
	p := params.Args
	if p["renderRes"] != nil {
		o.RenderRes = p["renderRes"].(string)
	}
	if p["direction"] != nil {
		o.Direction = p["direction"].(PageDirectionEnumType)
	}
	if p["pType"] != nil {
		o.PType = p["pType"].(PageTypeEnumType)
	}
	err := db.Create(&o).Error
	return o, err
}

func (o Page) Update(params graphql.ResolveParams) (Page, error) {
	v, ok := params.Source.(Page)
	if !ok {
		return o, errors.New("update param")
	}
	p := params.Args
	//		todo canvasJson:String 画布json数据
	//​		todo font:String 该画布包含的字体json数据
	if p["renderRes"] != nil {
		v.RenderRes = p["renderRes"].(string)
	}
	if p["status"] != nil {
		v.Status = p["status"].(ProJCommonStatusEnumType)
	}
	if p["direction"] != nil {
		v.Direction = p["direction"].(PageDirectionEnumType)
	}
	if p["pType"] != nil {
		v.PType = p["pType"].(PageTypeEnumType)
	}
	err := db.Save(&v).Error
	return v, err
}

func (o Page) Delete(params graphql.ResolveParams) (Page, error) {
	v, ok := params.Source.(Page)
	if !ok {
		return o, errors.New("delete param")
	}
	err := db.Delete(&v).Error
	return v, err
}
