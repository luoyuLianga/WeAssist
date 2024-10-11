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

// GetUserByUserName 根据用户名查询用户
func GetUserByUserName(username string) (user entity.User) {
	db.Db.Where("username = ?", username).First(&user)
	return user
}

// GetUserByUserId 根据用户id查询用户
func GetUserByUserId(id int) (user entity.User, err error) {
	err = db.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

// GetCountByT1 查询T+1的新用户数
func GetCountByT1(yesterdayStart util.HTime, yesterdayEnd util.HTime) (count int64, err error) {
	err = db.Db.Model(&entity.User{}).
		Where("create_time >= ? AND create_time < ?", yesterdayStart, yesterdayEnd).
		Count(&count).Error
	return count, err
}
