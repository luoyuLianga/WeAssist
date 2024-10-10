package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/util"
	"WeAssist/pkg/db"
	"time"
)

// AddOperationDayStats 添加操作
func AddOperationDayStats(dto entity.AddOperationDayStatsDto) (uint, error) {
	operation := entity.OperationDayStats{
		OpID:       dto.OpID,
		Source:     dto.Source,
		Day:        time.Now().Format("2006-01-02"),
		CreateTime: util.HTime{Time: time.Now()},
		UpdateTime: util.HTime{Time: time.Now()},
	}
	err := db.Db.Create(&operation).Error
	return operation.ID, err
}
