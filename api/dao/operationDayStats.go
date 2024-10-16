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

func GetMonthOperationDayStats(dto entity.GetMonthODSReqDto, startDate string, endDate string) (getDayODSRspDto []entity.GetMonthODSRspDto, err error) {
	query := db.Db.Table("operation_day_stats").
		Select("DATE_FORMAT(day, '%Y-%m') AS month, plugin_name, op_id, source, SUM(count) AS total_count").
		Where("day BETWEEN ? AND ?", startDate, endDate)

	// 添加条件筛选
	if dto.PluginName != "" {
		query = query.Where("plugin_name = ?", dto.PluginName)
	}
	if dto.OpID != 0 {
		query = query.Where("op_id = ?", dto.OpID)
	}
	if dto.Source != "" {
		query = query.Where("source = ?", dto.Source)
	}

	err = query.Group("month, plugin_name, op_id, source").
		Order("month, plugin_name, op_id, source").
		Scan(&getDayODSRspDto).Error
	return getDayODSRspDto, err
}

func GetDayOperationDayStats(dto entity.GetDayODSReqDto) (getDayODSRspDto []entity.GetDayODSRspDto, err error) {
	query := db.Db.Table("operation_day_stats").
		Select("day, plugin_name, op_id, source, SUM(count) AS count").
		Where("day BETWEEN ? AND ?", dto.StartDay, dto.EndDay)

	// 添加条件筛选
	if dto.PluginName != "" {
		query = query.Where("plugin_name = ?", dto.PluginName)
	}
	if dto.OpID != "" {
		query = query.Where("op_id = ?", dto.OpID)
	}
	if dto.Source != "" {
		query = query.Where("source = ?", dto.Source)
	}

	err = query.Group("day, plugin_name, op_id, source").
		Order("day, plugin_name, op_id, source").
		Scan(&getDayODSRspDto).Error

	return getDayODSRspDto, err
}
