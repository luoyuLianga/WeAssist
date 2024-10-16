package service

import (
	"WeAssist/api/dao"
	"WeAssist/api/entity"
	"WeAssist/common/result"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// IQaDayStatsService 定义接口
type IQaDayStatsService interface {
	QaDayStats() (dao.QaData, error)
	GetDay(c *gin.Context)
}

type QaDayStatsServiceImpl struct{}

func (q QaDayStatsServiceImpl) GetDay(c *gin.Context) {
	var dto entity.GetDayQDSReqDto
	if err := c.ShouldBindQuery(&dto); err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "GetDayQaDayStats() ShouldBindQuery Failed")
		return
	}

	getDayQDSDto, err := dao.GetDayQaDayStats(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "GetDayQaDayStats() Failed")
		return
	}
	result.Success(c, getDayQDSDto)
}

func (q QaDayStatsServiceImpl) QaDayStats() (qaDataList dao.QaData, err error) {
	// 1. 查询T+1天的 新增用户记录，获取pluginName、count
	yesterdayStart := time.Now().AddDate(0, 0, -1).Truncate(24 * time.Hour) // 昨天 00:00:00
	yesterdayEnd := yesterdayStart.Add(24 * time.Hour)                      // 昨天 23:59:59
	qaDataList, err = dao.GetQaData(yesterdayStart, yesterdayEnd)

	// 将userUsers转UserDataDayStats的数组
	var qaDayStatsList []entity.QADayStats
	for _, qaData := range qaDataList {
		qaDayStats := entity.QADayStats{
			PluginName: qaData.PluginName,
			Type:       qaData.Type,
			Source:     qaData.Source,
			Day:        time.Now().Format("2006-01-02"),
			Count:      qaData.Count,
			CodeNumber: qaData.CodeNumber,
		}
		qaDayStatsList = append(qaDayStatsList, qaDayStats)
	}

	// 检查是否有需要插入的数据
	if len(qaDayStatsList) == 0 {
		return nil, fmt.Errorf("no qaDayStats to insert")
	}
	// 批量插入
	err = dao.AddOrUpdateBatchQaDayStats(qaDayStatsList)
	return qaDataList, err
}

var qaDayStatsService = QaDayStatsServiceImpl{}

func QaDayStatsService() IQaDayStatsService {
	return &qaDayStatsService
}
