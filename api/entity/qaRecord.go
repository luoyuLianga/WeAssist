package entity

import (
	"WeAssist/common/util"
)

// QaRecord 表示 qa_record 表的模型
type QaRecord struct {
	ID            uint       `gorm:"column:id;primaryKey;autoIncrement;comment:'主键，自增'" json:"id"`
	UserPluginID  int        `gorm:"column:user_plugin_id;comment:'用户插件ID';not null" json:"userPluginId"`
	Source        string     `gorm:"column:source;type:varchar(64);not null;comment:'记录来源'" json:"source"`
	UserQuestion  string     `gorm:"column:user_question;type:text;not null;comment:'用户提问'" json:"userQuestion"`
	ModelResponse string     `gorm:"column:model_response;type:text;not null;comment:'模型回复'" json:"modelResponse"`
	CreateTime    util.HTime `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"createTime"`
	UpdateTime    util.HTime `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"updateTime"`
}

// TableName 表名 用于创建表
func (QaRecord) TableName() string {
	return "qa_record"
}

// AddQaRecordDto 新增QaRecordDto
type AddQaRecordDto struct {
	UserPluginID  int    `json:"userPluginId"  validate:"required"`
	Source        string `json:"Source" validate:"required"`
	UserQuestion  string `json:"userQuestion" validate:"required"`
	ModelResponse string `json:"modelResponse" validate:"required"`
}
