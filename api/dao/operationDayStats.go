package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/util"
	"WeAssist/pkg/db"
	"gorm.io/gorm"
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

// GetOperationDayStats 根据OpID、Source和Day查询
func GetOperationDayStats(opID uint, source string, day string) (operationDayStats entity.OperationDayStats) {
	db.Db.Where("op_id = ?", opID).Or("source = ?", source).Or("day = ?", day).First(&operationDayStats)
	return operationDayStats
}

func UpdateOperationDayStats(opID uint, source string, day string) (operationDayStats entity.OperationDayStats) {
	db.Db.Model(&operationDayStats).Where("op_id = ? AND source = ? AND day = ?", opID, source, day).
		UpdateColumn("op_id", gorm.Expr("op_id + ?", 1))
	return operationDayStats
}
