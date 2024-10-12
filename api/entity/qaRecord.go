package entity

import (
	"WeAssist/common/util"
)

// QaRecord 表示 qa_record 表的模型
type QaRecord struct {
	ID            uint       `gorm:"column:id;primaryKey;autoIncrement;comment:'主键，自增'" json:"id"`
	UserPluginID  uint       `gorm:"column:user_plugin_id;not null;comment:'用户插件ID'" json:"userPluginId"`
	Type          string     `gorm:"column:type;type:enum('qa', 'accept', 'reject');not null;comment:'记录类型'" json:"type"`
	Source        string     `gorm:"column:source;type:enum('chat', 'edit');not null;comment:'记录来源'" json:"source"`
	UserQuestion  *string    `gorm:"column:user_question;default:NULL;comment:'用户提问'" json:"userQuestion"`   // 允许为 NULL
	ModelResponse *string    `gorm:"column:model_response;default:NULL;comment:'模型回复'" json:"modelResponse"` // 允许为 NULL
	CodeNumber    int        `gorm:"column:code_number;not null;default:0;comment:'代码编号'" json:"codeNumber"`
	CreateTime    util.HTime `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"createTime"`
	UpdateTime    util.HTime `gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"updateTime"`
}

// TableName 表名 用于创建表
func (QaRecord) TableName() string {
	return "qa_record"
}

// AddQaRecordDto 新增QaRecordDto
type AddQaRecordDto struct {
	UserPluginID  uint   `json:"userPluginId"  validate:"required"`
	Type          string `json:"type" validate:"required,oneof=qa accept reject"`
	Source        string `json:"source" validate:"required,oneof=edit chat"`
	UserQuestion  string `json:"userQuestion"`
	ModelResponse string `json:"modelResponse"`
	CodeNumber    string `json:"codeNumber"`
}
