package entity

import "WeAssist/common/util"

// UserDataDayStats 表示 userdata_day_stats 表的模型
type UserDataDayStats struct {
	ID         uint       `gorm:"column:id;primaryKey;autoIncrement;comment:'主键，自增'" json:"id"`
	Type       string     `gorm:"column:type;type:enum('use_user', 'active_user');not null;comment:'用户类型（use_user，active_user）'" json:"type"`
	PluginName string     `gorm:"column:plugin_name;type:varchar(64);not null;comment:'插件名称'" json:"pluginName"`
	Day        string     `gorm:"column:day;type:varchar(64);not null;comment:'日期'" json:"day"`
	Count      uint       `gorm:"column:count;not null;default:0;comment:'计数'" json:"count"`
	CreateTime util.HTime `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"createTime"`
	UpdateTime util.HTime `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"updateTime"`
}

// TableName 自定义表名
func (UserDataDayStats) TableName() string {
	return "userdata_day_stats"
}

// AddUserDataDayStatsDto 用于新增 UserDataDayStats 的 DTO
type AddUserDataDayStatsDto struct {
	Type       string `json:"type" validate:"required,oneof=use_user active_user"` // 操作代码必填
	PluginName string `json:"pluginName" validate:"required"`                      // 操作描述必填
	Day        string `json:"day" validate:"required"`                             // 操作描述必填
	Count      uint   `json:"count" validate:"required"`                           // 操作描述必填
}
