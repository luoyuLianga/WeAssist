package service

import (
	"WeAssist/api/dao"
	"WeAssist/api/entity"
	"WeAssist/common/result"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// IQaRecordService 定义接口
type IQaRecordService interface {
	Add(c *gin.Context, dto entity.AddQaRecordDto)
}

type QaRecordServiceImpl struct{}

func (q QaRecordServiceImpl) Add(c *gin.Context, dto entity.AddQaRecordDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	userPlugin, _ := dao.GetUserPluginById(dto.UserPluginID)
	if userPlugin.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "用户插件不存在")
		return
	}

	_, err = dao.AddQaRecord(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "添加失败")
	}

	result.Success(c, "添加成功")
}

var qaRecordService = QaRecordServiceImpl{}

func QaRecordService() IQaRecordService {
	return &qaRecordService
}
