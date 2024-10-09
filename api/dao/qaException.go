package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/util"
	"WeAssist/pkg/db"
	"time"
)

// AddQaException 添加用户插件
func AddQaException(dto entity.AddQaExceptionDto) (uint, error) {
	qaException := entity.QaException{
		UserPluginID:  dto.UserPluginID,
		Source:        dto.Source,
		UserQuestion:  dto.UserQuestion,
		ExceptionInfo: dto.ExceptionInfo,
		CreateTime:    util.HTime{Time: time.Now()},
		UpdateTime:    util.HTime{Time: time.Now()},
	}
	err := db.Db.Create(&qaException).Error
	return qaException.ID, err
}
