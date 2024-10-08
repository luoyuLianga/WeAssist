package entity

import "WeAssist/common/util"

// 选手模型对象
type Player struct {
	ID          uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`         //ID
	Aid         int        `gorm:"column:aid;comment:'活动id'" json:"aId"`                         // 活动id
	Ref         string     `gorm:"column:ref;varchar(50);comment:'编号';NOT NULL" json:"ref"`      // 编号
	Nickname    string     `gorm:"column:nickname;varchar(64);comment:'昵称'" json:"nickname"`     // 昵称
	Declaration string     `gorm:"column:declaration;varchar(500);declaration:'宣言'" json:"icon"` //  宣言
	Avatar      string     `gorm:"column:avatar;varchar(64);comment:'头像'" json:"avatar"`         // 头像
	Score       int        `gorm:"column:score;comment:'分数'" json:"score"`                       // 分数
	Phone       string     `gorm:"column:phone;varchar(64);comment:'电话'" json:"phone"`           // 电话
	Note        string     `gorm:"column:note;varchar(500);comment:'备注'" json:"note"`            // 备注
	CreateTime  util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"` // 创建时间
	UpdateTime  util.HTime `gorm:"column:update_time;comment:'创建时间';NOT NULL" json:"updateTime"` // 更新时间
}

// 表名 用于创建表
func (Player) TableName() string {
	return "player"
}

// 获取选手列表
type PlayerListDto struct {
	Aid uint `json:"aid" validate:"required"` // 活动id
}

// 新增选手
type AddPlayerDto struct {
	Aid         int    `json:"aId"`         // 活动id
	Ref         string `json:"ref"`         // 编号
	Nickname    string `json:"nickname"`    // 昵称
	Declaration string `json:"declaration"` //  宣言
	Avatar      string `json:"avatar"`      // 头像
	Score       int    `json:"score"`       // 分数
	Phone       string `json:"phone"`       // 电话
	Note        string `json:"note"`        // 备注
}

// 更新选手
type UpdatePlayerDto struct {
	ID    uint `json:"id"`    //ID
	Aid   int  `json:"aId"`   // 活动id
	Score int  `json:"score"` // 分数
}
