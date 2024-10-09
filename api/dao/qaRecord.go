package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/util"
	"WeAssist/pkg/db"
	"time"
)

// AddQaRecord 添加用户插件
func AddQaRecord(dto entity.AddQaRecordDto) (uint, error) {
	qaRecord := entity.QaRecord{
		UserPluginID:  dto.UserPluginID,
		Source:        dto.Source,
		UserQuestion:  dto.UserQuestion,
		ModelResponse: dto.ModelResponse,
		CreateTime:    util.HTime{Time: time.Now()},
		UpdateTime:    util.HTime{Time: time.Now()},
	}
	err := db.Db.Create(&qaRecord).Error
	return qaRecord.ID, err
}
