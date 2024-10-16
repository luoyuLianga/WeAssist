package entity

import (
	"WeAssist/common/util"
)

// QADayStats 表示 qa_day_stats 表的模型
type QADayStats struct {
	ID         uint       `gorm:"column:id;primaryKey;autoIncrement;comment:'主键，自增'" json:"id"`
	PluginName string     `gorm:"column:plugin_name;type:varchar(64);not null;comment:'插件名称'" json:"pluginName"`
	Type       string     `gorm:"column:type;type:enum('qa', 'accept', 'reject');not null;comment:'数据类型（qa, accept, reject）'" json:"type"`
	Source     string     `gorm:"column:source;type:enum('chat', 'edit');not null;comment:'数据来源（chat, edit）'" json:"source"`
	Day        string     `gorm:"column:day;type:varchar(64);not null;comment:'日期'" json:"day"`
	Count      int        `gorm:"column:count;not null;default:0;comment:'计数'" json:"count"`
	CodeNumber int        `gorm:"column:code_number;not null;default:0;comment:'代码编号'" json:"codeNumber"`
	CreateTime util.HTime `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"createTime"`
	UpdateTime util.HTime `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"updateTime"`
}

// TableName 指定表名为 qa_day_stats
func (QADayStats) TableName() string {
	return "qa_day_stats"
}

type GetMonthQDSReqDto struct {
	PluginName string `form:"pluginName"`
	Type       string `form:"source"`
	Source     string `form:"source"`
}

type GetMonthQDSRspDto struct {
	Month           string `json:"month"`
	PluginName      string `json:"pluginName"`
	Type            string `json:"opId"`
	Source          string `json:"source"`
	TotalCount      uint   `json:"totalCount"`
	TotalCodeNumber uint   `json:"totalCodeNumber"`
}

type GetDayQDSReqDto struct {
	StartDay   string `form:"startDay" binding:"required"`
	EndDay     string `form:"endDay" binding:"required"`
	PluginName string `form:"pluginName"`
	Type       string `form:"type"`
	Source     string `form:"source"`
}

type GetDayQDSRspDto struct {
	Day        string `json:"day"`
	PluginName string `json:"pluginName"`
	Type       string `json:"type"`
	Source     string `json:"source"`
	Count      uint   `json:"count"`
	CodeNumber uint   `json:"codeNumber"`
}
