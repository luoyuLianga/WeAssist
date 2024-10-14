package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/config"
	"WeAssist/pkg/db"
	"gorm.io/gorm/clause"
	"time"
)

type QaData []struct {
	PluginName string
	Type       string
	Source     string
	Count      int
	CodeNumber int
}

// GetQaData 根据T+1时间查询
func GetQaData(yesterdayStart time.Time, yesterdayEnd time.Time) (qaData QaData, err error) {
	// 执行查询
	err = db.Db.Table("qa_record").Select(`
		up.plugin_name,
		qa_record.type,
		qa_record.source,
		COUNT(*) AS count,
		SUM(qa_record.code_number) AS code_number`).
		Joins("JOIN user_plugin up ON qa_record.user_plugin_id = up.id").
		Where("qa_record.create_time >= ? AND qa_record.create_time < ?", yesterdayStart, yesterdayEnd).
		Group("up.plugin_name, qa_record.type, qa_record.source").
		Scan(&qaData).Error
	return qaData, err
}

// AddOrUpdateBatchQaDayStats 添加操作
func AddOrUpdateBatchQaDayStats(dto []entity.QADayStats) (err error) {
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
			Columns:   []clause.Column{{Name: "plugin_name"}, {Name: "type"}, {Name: "source"}, {Name: "day"}}, // 唯一键冲突字段
			DoUpdates: clause.AssignmentColumns([]string{"count"}),                                             // 更新 count 和 update_time
		}).Create(&data).Error

		if err != nil {
			return err
		}
	}
	return nil
}
