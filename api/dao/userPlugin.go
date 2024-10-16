package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/util"
	"WeAssist/pkg/db"
	"time"
)

// AddUserPlugin 添加用户插件
func AddUserPlugin(dto entity.AddUserPluginDto) (uint, error) {
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

// GetUserPluginById 根据用户插件id查询
func GetUserPluginById(id uint) (userPlugin entity.UserPlugin, err error) {
	err = db.Db.Where("id = ?", id).First(&userPlugin).Error
	return userPlugin, err
}

// GetUserPluginByPM 根据PluginName和ModelName查询
func GetUserPluginByPM(userId int, pluginName string) (userPlugin entity.UserPlugin, err error) {
	err = db.Db.Where("user_id = ? AND plugin_name = ?", userId, pluginName).First(&userPlugin).Error
	return userPlugin, err
}
