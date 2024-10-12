package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/config"
	"WeAssist/pkg/db"
	"gorm.io/gorm/clause"
	"time"
)

type UserData []struct {
	PluginName string
	UserCount  uint
}

// GetUseUser 根据T+1时间查询
func GetUseUser(yesterdayStart time.Time, yesterdayEnd time.Time) (userData UserData, err error) {
	// 执行查询
	err = db.Db.Table("user_plugin").
		Select("user_plugin.plugin_name, COUNT(DISTINCT user.id) as user_count").
		Joins("JOIN user ON user.id = user_plugin.user_id").
		Where("user.create_time >= ? AND user.create_time < ?", yesterdayStart, yesterdayEnd).
		Group("user_plugin.plugin_name").
		Scan(&userData).Error
	return userData, err
}

// GetActiveUser 根据T+1时间查询
func GetActiveUser(yesterdayStart time.Time, yesterdayEnd time.Time) (userData UserData, err error) {
	// 执行查询
	err = db.Db.Table("qa_record AS qr").
		Select("up.plugin_name, COUNT(DISTINCT up.user_id) AS user_count").
		Joins("JOIN user_plugin AS up ON qr.user_plugin_id = up.id").
		Where("qr.create_time >= ? AND qr.create_time < ?", yesterdayStart, yesterdayEnd).
		Group("up.plugin_name").
		Scan(&userData).Error

	return userData, err
}

// AddOrUpdateBatchUserDataDayStats 添加操作
func AddOrUpdateBatchUserDataDayStats(dto []entity.UserDataDayStats) (err error) {
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
