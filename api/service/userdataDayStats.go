package service

import (
	"WeAssist/api/dao"
	"WeAssist/api/entity"
	"fmt"
	"time"
)

// IUserDataDayStatsService 定义接口
type IUserDataDayStatsService interface {
	UseUserDayStats() (dao.UseUsers, error)
}

type UserDataDayStatsServiceImpl struct{}

func (uds UserDataDayStatsServiceImpl) UseUserDayStats() (useUsers dao.UseUsers, err error) {
	// 1. 查询T+1天的 新增用户记录，获取pluginName、count
	yesterdayStart := time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour) // 昨天 00:00:00
	yesterdayEnd := yesterdayStart.Add(24 * time.Hour)                      // 昨天 23:59:59
	useUsers, err = dao.GetUseUser(yesterdayStart, yesterdayEnd)

	// 将userUsers转AddUserDataDayStatsDto的数组
	var userDataDayStatsList []entity.UserDataDayStats
	for _, user := range useUsers {
		userDataDayStats := entity.UserDataDayStats{
			Type:       "use_user",
			PluginName: user.PluginName,
			Count:      user.UserCount,
			Day:        time.Now().Format("2006-01-02"),
		}
		userDataDayStatsList = append(userDataDayStatsList, userDataDayStats)
	}

	// 检查是否有需要插入的数据
	if len(userDataDayStatsList) == 0 {
		return nil, fmt.Errorf("no user stats to insert")
	}

	// 批量插入
	err = dao.AddOrUpdateBatchUserDataDayStats(userDataDayStatsList)
	return useUsers, err
}

var userDataDayStatsService = UserDataDayStatsServiceImpl{}

func UserDataDayStatsService() IUserDataDayStatsService {
	return &userDataDayStatsService
}
