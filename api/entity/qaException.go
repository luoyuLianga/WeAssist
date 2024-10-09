package entity

import (
	"WeAssist/common/util"
)

// QaException represents the 'qa_exception' table structure
type QaException struct {
	ID            uint       `gorm:"column:id;comment:'主键';primaryKey;autoIncrement;NOT NULL" json:"id"`
	UserPluginID  uint       `gorm:"column:user_plugin_id;comment:'用户插件ID';NOT NULL" json:"user_plugin_id"`
	Source        string     `gorm:"column:source;type:varchar(100);comment:'来源';NOT NULL" json:"source"`
	UserQuestion  string     `gorm:"column:user_question;type:text;comment:'用户问题';NOT NULL" json:"user_question"`
	ExceptionInfo string     `gorm:"column:exception_info;type:text;comment:'异常信息';NOT NULL" json:"exception_info"`
	CreateTime    util.HTime `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"createTime"`
	UpdateTime    util.HTime `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"updateTime"`
}

// TableName 表名 用于创建表
func (QaException) TableName() string {
	return "qa_exception"
}

// AddQaExceptionDto 新增QaRecordDto
type AddQaExceptionDto struct {
	UserPluginID  uint   `json:"userPluginId"  validate:"required"`
	Source        string `json:"source" validate:"required,oneof=edit chat"`
	UserQuestion  string `json:"userQuestion" validate:"required"`
	ExceptionInfo string `json:"exceptionInfo" validate:"required"`
}
