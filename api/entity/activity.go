package entity

import "WeAssist/common/util"

// 用户模型对象
type Activity struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`         // ID
	Name       string     `gorm:"column:name;varchar(64);comment:'活动名称';NOT NULL" json:"name"`  // 活动名称
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"` // 添加时间
}

// 表名 用于创建表
func (Activity) TableName() string {
	return "activity"
}

// ActivityDto
type ActivityDto struct {
	Name string `json:"name" validate:"required"` // 活动名
}
