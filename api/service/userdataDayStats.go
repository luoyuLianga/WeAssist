package service

import (
	"WeAssist/api/dao"
	"time"
)

// IUserDataDayStatsService 定义接口
type IUserDataDayStatsService interface {
	Add() (dao.UseUsers, error)
}

type UserDataDayStatsServiceImpl struct{}

func (uds UserDataDayStatsServiceImpl) Add() (useUsers dao.UseUsers, err error) {
	yesterdayStart := time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour) // 昨天 00:00:00
	yesterdayEnd := yesterdayStart.Add(24 * time.Hour)                      // 昨天 23:59:59
	useUsers, err = dao.GetUseUser(yesterdayStart, yesterdayEnd)
	return
}

var userDataDayStatsService = UserDataDayStatsServiceImpl{}

func UserDataDayStatsService() IUserDataDayStatsService {
	return &userDataDayStatsService
}
