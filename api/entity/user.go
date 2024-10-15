package entity

import "WeAssist/common/util"

// User 用户模型对象
type User struct {
	ID       uint   `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                //ID
	Username string `gorm:"column:username;varchar(64);comment:'用户账号';NOT NULL" json:"username"` // 用户账号
	Password string `gorm:"column:password;varchar(64);comment:'密码';NOT NULL" json:"-"`          // 密码

	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"` // 创建时间
	UpdateTime util.HTime `gorm:"column:update_time;comment:'创建时间';NOT NULL" json:"updateTime"` // 更新时间
}

// TableName 表名 用于创建表
func (User) TableName() string {
	return "user"
}

// UserRegisterDto 注册参数
type UserRegisterDto struct {
	Username        string `json:"username" validate:"required"`        // 用户名
	Password        string `json:"password" validate:"required"`        // 密码
	ConfirmPassword string `json:"confirmPassword" validate:"required"` // 密码
}

// UserLoginDto 登录参数
type UserLoginDto struct {
	Username string `json:"username" validate:"required"` // 用户名
	Password string `json:"password" validate:"required"` // 密码
}

// UpdateUserDto 登录参数
type UpdateUserDto struct {
	ID       int    `json:"id" validate:"required"`       // 用户id
	Username string `json:"username" validate:"required"` // 用户名
	Password string `json:"password" validate:"required"` // 密码
}
