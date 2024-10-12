package dao

import (
	"WeAssist/pkg/db"
	"time"
)

type QaData []struct {
	PluginName string
	Type       string
	Source     string
	Count      uint
	CodeNumber uint
}

// GetQaData 根据T+1时间查询
func GetQaData(yesterdayStart time.Time, yesterdayEnd time.Time) (qaData QaData, err error) {
	// 执行查询
	err = db.Db.Table("qa_record").Select(`
		up.plugin_name,
		qr.type,
		qr.source,
		COUNT(*) AS count,
		SUM(qr.code_number) AS code_number`).
		Joins("JOIN user_plugin up ON qr.user_plugin_id = up.id").
		Where("qr.create_time >= ? AND qr.create_time < ?", yesterdayStart, yesterdayEnd).
		Group("up.plugin_name, qr.type, qr.source").
		Scan(&qaData).Error
	return qaData, err
}
