package entity

import "WeAssist/common/util"

// 活动投票模型对象
type Vote struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`         //ID
	UserId     int        `gorm:"column:user_id;comment:'用户账号';NOT NULL" json:"userId"`         // 用户账号
	PlayerId   int        `gorm:"column:player_id;comment:'密码';NOT NULL" json:"playId"`         // 密码
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"` // 创建时间
}

// 表名 用于创建表
func (Vote) TableName() string {
	return "vote"
}

// VoteDto 投票参数
type VoteDto struct {
	PlayerId int `json:"playerId" validate:"required"` // 用户名
	UserId   int `json:"userId" validate:"required"`   // 密码
}
