package cron

import (
	"WeAssist/api/service"
	"WeAssist/pkg/log"
)

// UseUserDayStats 用户数据统计 使用用户数据统计
func UseUserDayStats() {
	useUsers, err := service.UserDataDayStatsService().UseUserDayStats()
	if err != nil {
		log.Log().Errorf("Failed to fetch use user stats: %v", err)
	}
	log.Log().Infof("useUsers:%v", useUsers)
}

// ActiveUserDayStats 用户数据统计 活跃用户数据统计
func ActiveUserDayStats() {
	activeUsers, err := service.UserDataDayStatsService().ActiveUserDayStats()
	if err != nil {
		log.Log().Errorf("Failed to fetch user stats: %v", err)
	}
	log.Log().Infof("activeUsers:%v", activeUsers)
}
