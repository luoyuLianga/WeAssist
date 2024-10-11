package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/util"
	"WeAssist/pkg/db"
	"time"
)

type UseUsers []struct {
	PluginName string
	UserCount  int
}

// GetUseUser 根据OpID、Source和Day查询
func GetUseUser(yesterdayStart time.Time, yesterdayEnd time.Time) (useUsers UseUsers, err error) {
	// 执行查询
	err = db.Db.Table("user_plugin").
		Select("user_plugin.plugin_name, COUNT(DISTINCT user.id) as user_count").
		Joins("JOIN user ON user.id = user_plugin.user_id").
		Where("user.create_time >= ? AND user.create_time < ?", yesterdayStart, yesterdayEnd).
		Group("user_plugin.plugin_name").
		Scan(&useUsers).Error
	return useUsers, err
}

// AddUserDataDayStats 添加操作
func AddUserDataDayStats(dto entity.AddUserDataDayStatsDto) (uint, error) {
	userDataDayStats := entity.UserDataDayStats{
		Type:       dto.Type,
		PluginName: dto.PluginName,
		Day:        time.Now().Format("2006-01-02"),
		Count:      dto.Count,
		CreateTime: util.HTime{Time: time.Now()},
		UpdateTime: util.HTime{Time: time.Now()},
	}
	err := db.Db.Create(&userDataDayStats).Error
	return userDataDayStats.ID, err
}
