package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/util"
	"WeAssist/pkg/db"
	"gorm.io/gorm"
	"time"
)

// AddOperationDayStats 添加操作
func AddOperationDayStats(dto entity.OperationDayStatsDto) (uint, error) {
	operation := entity.OperationDayStats{
		PluginName: dto.PluginName,
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
func GetOperationDayStats(dto entity.OperationDayStatsDto, day string) (operationDayStats entity.OperationDayStats) {
	db.Db.Where("plugin_name = ? AND op_id = ? AND source = ? AND day = ?", dto.PluginName, dto.OpID, dto.Source, day).
		First(&operationDayStats)
	return operationDayStats
}

func UpdateOperationDayStats(dto entity.OperationDayStatsDto, day string) (operationDayStats entity.OperationDayStats, err error) {
	err = db.Db.Model(&operationDayStats).Where("plugin_name = ? AND op_id = ? AND source = ? AND day = ?",
		dto.PluginName, dto.OpID, dto.Source, day).
		UpdateColumn("count", gorm.Expr("count + ?", 1)).Error
	return operationDayStats, err
}

func GetMonthOperationDayStats(startDate string, endDate string) (getMonthODSDto []entity.GetMonthODSDto, err error) {
	err = db.Db.Table("operation_day_stats").
		Select("DATE_FORMAT(day, '%Y-%m') AS month, plugin_name, op_id, source, SUM(count) AS total_count").
		Where("day BETWEEN ? AND ?", startDate, endDate).
		Group("month, plugin_name, op_id, source").
		Order("month, plugin_name, op_id, source").
		Scan(&getMonthODSDto).Error
	return getMonthODSDto, err
}
