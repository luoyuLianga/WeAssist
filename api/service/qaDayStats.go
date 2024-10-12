package service

import (
	"WeAssist/api/dao"
	"time"
)

// IQaDayStatsService 定义接口
type IQaDayStatsService interface {
	QaDayStats() (dao.QaData, error)
}

type QaDayStatsServiceImpl struct{}

func (q QaDayStatsServiceImpl) QaDayStats() (qaData dao.QaData, err error) {
	// 1. 查询T+1天的 新增用户记录，获取pluginName、count
	yesterdayStart := time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour) // 昨天 00:00:00
	yesterdayEnd := yesterdayStart.Add(24 * time.Hour)                      // 昨天 23:59:59
	qaData, err = dao.GetQaData(yesterdayStart, yesterdayEnd)
	return qaData, err
}

var qaDayStatsService = QaDayStatsServiceImpl{}

func QaDayStatsService() IQaDayStatsService {
	return &qaDayStatsService
}
