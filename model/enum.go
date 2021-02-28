package model

import "github.com/Fiber-Man/funplugin"

// 设计器项目通用状态枚举类型
type ProJCommonStatusEnumType uint

const (
	P_ENABLE  ProJCommonStatusEnumType = 1 // 正常
	P_DISABLE ProJCommonStatusEnumType = 2 // 停用
	P_DELETE  ProJCommonStatusEnumType = 2 // 删除
)

func (s ProJCommonStatusEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"P_ENABLE": funplugin.EnumValue{
			Value:       P_ENABLE,
			Description: "正常",
		},
		"P_DISABLE": funplugin.EnumValue{
			Value:       P_DISABLE,
			Description: "停用",
		},
		"P_DELETE": funplugin.EnumValue{
			Value:       P_DELETE,
			Description: "删除",
		},
	}
}

// 画布页面种类枚举类型
type PageTypeEnumType uint

const (
	COVER       PageTypeEnumType = 1 // 封面
	BACK_COVER  PageTypeEnumType = 2 // 封底
	CERTIFICATE PageTypeEnumType = 3 // 证书
	NORMAL      PageTypeEnumType = 4 // 普通
	TITLE_PAGE  PageTypeEnumType = 5 // 扉页
)

func (s PageTypeEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"COVER": funplugin.EnumValue{
			Value:       COVER,
			Description: "封面",
		},
		"BACK_COVER": funplugin.EnumValue{
			Value:       BACK_COVER,
			Description: "封底",
		},
		"CERTIFICATE": funplugin.EnumValue{
			Value:       CERTIFICATE,
			Description: "证书",
		},
		"NORMAL": funplugin.EnumValue{
			Value:       NORMAL,
			Description: "普通",
		},
		"TITLE_PAGE": funplugin.EnumValue{
			Value:       TITLE_PAGE,
			Description: "扉页",
		},
	}
}

// 画布页面方向枚举类型
type PageDirectionEnumType uint

const (
	V PageDirectionEnumType = 1 // 垂直方向
	H PageDirectionEnumType = 2 // 水平方向
)

func (s PageDirectionEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"V": funplugin.EnumValue{
			Value:       V,
			Description: "垂直方向",
		},
		"H": funplugin.EnumValue{
			Value:       H,
			Description: "水平方向",
		},
	}
}
