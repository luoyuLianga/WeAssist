package cron

import (
	"WeAssist/api/service"
	"WeAssist/pkg/log"
)

// QaDayStats 问答数据统计
func QaDayStats() {
	qaData, err := service.QaDayStatsService().QaDayStats()
	if err != nil {
		log.Log().Errorf("Failed to fetch use user stats: %v", err)
	}
	log.Log().Infof("qaData:%v", qaData)
}
