package entity

import (
	"WeAssist/common/util"
)

// OperationDayStats 表示 operation_day_stats 表的模型
type OperationDayStats struct {
	ID         uint       `gorm:"column:id;primaryKey;autoIncrement;comment:'主键，自增'" json:"id"`
	PluginName string     `gorm:"column:plugin_name;type:varchar(64);not null;comment:'插件名'" json:"plugin_name"`
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

// OperationDayStatsDto 用于新增 OperationDayStats 的 DTO
type OperationDayStatsDto struct {
	PluginName string `json:"pluginName" validate:"required"`
	OpID       uint   `json:"operationId" validate:"required"`            // 操作代码必填
	Source     string `json:"source" validate:"required,oneof=edit chat"` // 操作描述必填
}

// GetMonthODSDto 用于保存查询结果的结构体
type GetMonthODSDto struct {
	Month      string `json:"month"`
	PluginName string `json:"plugin_name"`
	OpID       uint   `json:"opId"`
	Source     string `json:"source"`
	TotalCount uint   `json:"total_count"`
}

type GetDayODSReqDto struct {
	StartDay string `json:"startDay"`
	EndDay   string `json:"EndDay"`
}

type GetDayODSRspDto struct {
	Day        string `json:"Day"`
	PluginName string `json:"plugin_name"`
	Source     string `json:"source"`
	TotalCount uint   `json:"total_count"`
}
