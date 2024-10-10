package entity

import (
	"WeAssist/common/util"
)

// OperationDayStats 表示 operation_day_stats 表的模型
type OperationDayStats struct {
	ID         uint       `gorm:"column:id;primaryKey;autoIncrement;comment:'主键，自增'" json:"id"`
	OpID       uint       `gorm:"column:op_id;not null;comment:'关联 user_operation 表的 id'" json:"opId"`
	Source     string     `gorm:"column:source;type:varchar(64);not null;comment:'来源'" json:"source"`
	Day        string     `gorm:"column:day;type:varchar(64);not null;comment:'日期'" json:"day"`
	Count      uint       `gorm:"column:count;not null;default:0;comment:'计数'" json:"count"`
	CreateTime util.HTime `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"createTime"`
	UpdateTime util.HTime `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"updateTime"`
}

// TableName 指定表名
func (OperationDayStats) TableName() string {
	return "operation_day_stats"
}

// AddOperationDayStatsDto 用于新增 OperationDayStats 的 DTO
type AddOperationDayStatsDto struct {
	OpID   uint   `json:"operationId" validate:"required"` // 操作代码必填
	Source string `json:"source" validate:"required"`      // 操作描述必填
}

// UpdateOperationDayStatsDto 用于新增 OperationDayStats 的 DTO
type UpdateOperationDayStatsDto struct {
	OpID   uint   `json:"operationId" validate:"required"` // 操作代码必填
	Source string `json:"source" validate:"required"`      // 操作描述必填
}
