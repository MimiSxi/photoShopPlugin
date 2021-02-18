package model

import "github.com/Fiber-Man/funplugin"

type DeviceStatusEnumType uint

const (
	DeviceStatus_NONE     DeviceStatusEnumType = 0 // 未知
	DeviceStatus_NORMAL   DeviceStatusEnumType = 1 // 正常
	DeviceStatus_SHUTDOWN DeviceStatusEnumType = 2 // 停用
	DeviceStatus_VERIFY   DeviceStatusEnumType = 3 // 校验
	DeviceStatus_RETURN   DeviceStatusEnumType = 4 //  归还
)

func (s DeviceStatusEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"NONE": funplugin.EnumValue{
			Value:       DeviceStatus_NONE,
			Description: "未知",
		},
		"NORMAL": funplugin.EnumValue{
			Value:       DeviceStatus_NORMAL,
			Description: "正常",
		},
		"SHUTDOWN": funplugin.EnumValue{
			Value:       DeviceStatus_SHUTDOWN,
			Description: "停用",
		},
		"VERIFY": funplugin.EnumValue{
			Value:       DeviceStatus_VERIFY,
			Description: "校验",
		},
		"RETURN": funplugin.EnumValue{
			Value:       DeviceStatus_RETURN,
			Description: "归还",
		},
	}
}

type BorrowStatusEnumType uint

const (
	BorrowStatus_NONE          BorrowStatusEnumType = 0 // 未知
	BorrowStatus_NO_APPLY      BorrowStatusEnumType = 1 // 未申请
	BorrowStatus_TOBE_BORROWED BorrowStatusEnumType = 2 // 待领用
	BorrowStatus_TOBE_RETURNED BorrowStatusEnumType = 3 // 待归还
)

func (s BorrowStatusEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"NONE": funplugin.EnumValue{
			Value:       BorrowStatus_NONE,
			Description: "未知",
		},
		"NO_APPLY": funplugin.EnumValue{
			Value:       BorrowStatus_NO_APPLY,
			Description: "未申请",
		},
		"TOBE_BORROWED": funplugin.EnumValue{
			Value:       BorrowStatus_TOBE_BORROWED,
			Description: "待领用",
		},
		"TOBE_RETURNED": funplugin.EnumValue{
			Value:       BorrowStatus_TOBE_RETURNED,
			Description: "待归还",
		},
	}
}

type QueryTypeEnumType uint

const (
	QueryType_ID           QueryTypeEnumType = 1 // 设备编号
	QueryType_NAME         QueryTypeEnumType = 2 // 设备名称
	QueryType_TYPE         QueryTypeEnumType = 3 // 设备型号
	QueryType_STATUS       QueryTypeEnumType = 4 // 设备状态
	QueryType_BORROWSTATUE QueryTypeEnumType = 5 // 借用状态
)

func (s QueryTypeEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"ID": funplugin.EnumValue{
			Value:       QueryType_ID,
			Description: "设备编号",
		},
		"NAME": funplugin.EnumValue{
			Value:       QueryType_NAME,
			Description: "设备名称",
		},
		"TYPE": funplugin.EnumValue{
			Value:       QueryType_TYPE,
			Description: "设备型号",
		},
		"STATUS": funplugin.EnumValue{
			Value:       QueryType_STATUS,
			Description: "设备状态",
		},
		"BORROWSTATUE": funplugin.EnumValue{
			Value:       QueryType_BORROWSTATUE,
			Description: "借用状态",
		},
	}
}

////////////////////////////////////////////////////////////

type ApplyFormStatusEnumType uint

const (
	ApplyFormStatus_APPLIED  ApplyFormStatusEnumType = 1 // 申请
	ApplyFormStatus_BORROWED ApplyFormStatusEnumType = 2 // 领用
	ApplyFormStatus_RETURNED ApplyFormStatusEnumType = 3 // 归还
	ApplyFormStatus_CANCEL   ApplyFormStatusEnumType = 4 // 取消
)

func (s ApplyFormStatusEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"APPLIED": funplugin.EnumValue{
			Value:       ApplyFormStatus_APPLIED,
			Description: "申请",
		},
		"BORROWED": funplugin.EnumValue{
			Value:       ApplyFormStatus_BORROWED,
			Description: "领用",
		},
		"RETURNED": funplugin.EnumValue{
			Value:       ApplyFormStatus_RETURNED,
			Description: "归还",
		},
		"CANCEL": funplugin.EnumValue{
			Value:       ApplyFormStatus_CANCEL,
			Description: "取消",
		},
	}
}
