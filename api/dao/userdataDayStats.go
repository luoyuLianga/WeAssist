package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/config"
	"WeAssist/pkg/db"
	"gorm.io/gorm/clause"
	"time"
)

type UseUsers []struct {
	PluginName string
	UserCount  uint
}

// GetUseUser 根据OpID、Source和Day查询
func GetUseUser(yesterdayStart time.Time, yesterdayEnd time.Time) (useUsers UseUsers, err error) {
	// 执行查询
	err = db.Db.Table("user_plugin").
		Select("user_plugin.plugin_name, COUNT(DISTINCT user.id) as user_count").
		Joins("JOIN user ON user.id = user_plugin.user_id").
		Where("user.create_time >= ? AND user.create_time < ?", yesterdayStart, yesterdayEnd).
		Group("user_plugin.plugin_name").
		Scan(&useUsers).Error
	return useUsers, err
}

// AddOrUpdateBatchUserDataDayStats 添加操作
func AddOrUpdateBatchUserDataDayStats(dto []entity.AddUserDataDayStatsDto) (err error) {
	// 分批插入
	batchSize := config.Config.Db.BatchSize
	for i := 0; i < batchSize; i += batchSize {
		end := i + batchSize
		if end > len(dto) {
			end = len(dto)
		}

		data := dto[i:end]
		// 使用 ON DUPLICATE KEY UPDATE 的方式处理冲突
		err := db.Db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "type"}, {Name: "plugin_name"}, {Name: "day"}}, // 唯一键冲突字段
			DoUpdates: clause.AssignmentColumns([]string{"count"}),                           // 更新 count 和 update_time
		}).Create(&data).Error

		if err != nil {
			return err
		}
	}
	return nil
}
