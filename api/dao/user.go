package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/util"
	"WeAssist/pkg/db"
	"time"
)

// 注册
func Register(dto entity.UserRegisterDto) (uint, error) {
	user := entity.User{
		Username:   dto.Username,
		Password:   util.EncryptionMd5(dto.Password),
		CreateTime: util.HTime{Time: time.Now()},
		UpdateTime: util.HTime{Time: time.Now()},
	}
	err := db.Db.Create(&user).Error
	return user.ID, err
}

// 根据用户名查询用户
func GetUserByUserName(username string) (user entity.User) {
	db.Db.Where("username = ?", username).First(&user)
	return user
}

// 根据用户名查询用户
func GetUserByUserId(id int) (user entity.User, err error) {
	err = db.Db.Where("id = ?", id).First(&user).Error
	return user, err
}
