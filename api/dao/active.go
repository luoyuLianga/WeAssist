package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/util"
	"WeAssist/pkg/db"
	"time"
)

// 新增活动
func AddActivity(dto entity.ActivityDto) error {
	user := entity.Activity{
		Name:       dto.Name,
		CreateTime: util.HTime{Time: time.Now()},
	}
	err := db.Db.Create(&user).Error
	return err
}
