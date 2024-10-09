package entity

import "WeAssist/common/util"

// Operation 表示 operation 表的模型
type Operation struct {
	ID            uint       `gorm:"column:id;primaryKey;autoIncrement;comment:'主键，自增'" json:"id"`
	OperationCode string     `gorm:"column:operation_code;type:varchar(128);not null;comment:'操作代码'" json:"operationCode"`
	OperationDesc string     `gorm:"column:operation_desc;type:varchar(128);not null;comment:'操作描述'" json:"operationDesc"`
	CreateTime    util.HTime `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"createTime"`
	UpdateTime    util.HTime `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"updateTime"`
}

// TableName 表名 用于创建表
func (Operation) TableName() string {
	return "operation"
}

// AddOperationDto 用于新增 Operation 的 DTO
type AddOperationDto struct {
	OperationCode string `json:"operationCode" validate:"required"` // 操作代码必填
	OperationDesc string `json:"operationDesc" validate:"required"` // 操作描述必填
}
