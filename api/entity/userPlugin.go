package entity

import "WeAssist/common/util"

// UserPlugin 用户插件模型对象
type UserPlugin struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                    //ID
	UserId     int        `gorm:"column:user_id;comment:'用户id'" json:"userId"`                             // 用户id
	PluginName string     `gorm:"column:plugin_name;varchar(64);comment:'插件名';NOT NULL" json:"pluginName"` // 插件名
	ModelName  string     `gorm:"column:model_name;varchar(128);comment:'模型名';NOT NULL" json:"modelName"`  // 模型名
	Action     string     `gorm:"column:action;varchar(64);comment:'操作';NOT NULL" json:"action"`           // 插件名
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`            // 创建时间
	UpdateTime util.HTime `gorm:"column:update_time;comment:'创建时间';NOT NULL" json:"updateTime"`            // 更新时间
}

// TableName 表名 用于创建表
func (UserPlugin) TableName() string {
	return "user_plugin"
}

// AddUserPluginDto 新增选手
type AddUserPluginDto struct {
	UserId     int    `json:"userId"  validate:"required"`
	PluginName string `json:"pluginName" validate:"required"`
	ModelName  string `json:"modelName" validate:"required"`
	Action     string `json:"action" validate:"required,oneof=install upgrade uninstall"`
}
