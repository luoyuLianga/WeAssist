package cron

import (
	"WeAssist/api/service"
	"WeAssist/pkg/log"
)

func GetUseUser() {
	useUsers, err := service.UserDataDayStatsService().Add()
	if err != nil {
		log.Log().Error("Failed to fetch user stats: %v", err)
	}
	log.Log().Info("useUsers:%#v", useUsers)
}
