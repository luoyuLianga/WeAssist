package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/util"
	"WeAssist/pkg/db"
	"time"
)

// 添加用户插件
func Add(dto entity.AddUserPluginDto) (uint, error) {
	userPlugin := entity.UserPlugin{
		UserId:     dto.UserId,
		PluginName: dto.PluginName,
		ModelName:  dto.ModelName,
		Action:     dto.Action,
		CreateTime: util.HTime{Time: time.Now()},
		UpdateTime: util.HTime{Time: time.Now()},
	}
	err := db.Db.Create(&userPlugin).Error
	return userPlugin.ID, err
}
