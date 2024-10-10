package service

import (
	"WeAssist/api/dao"
	"WeAssist/api/entity"
	"WeAssist/common/result"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"time"
)

// IOperationDayStatsService 定义接口
type IOperationDayStatsService interface {
	Add(c *gin.Context, dto entity.AddOperationDayStatsDto)
	Update(c *gin.Context, dto entity.UpdateOperationDayStatsDto)
}

type OperationDayStatsServiceImpl struct{}

func (ods OperationDayStatsServiceImpl) Add(c *gin.Context, dto entity.AddOperationDayStatsDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	operation, _ := dao.GetOperationById(dto.OpID)
	if operation.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "OperationId不存在")
		return
	}

	_, err = dao.AddOperationDayStats(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "添加失败")
	}
	result.Success(c, "添加成功")
}

func (ods OperationDayStatsServiceImpl) Update(c *gin.Context, dto entity.UpdateOperationDayStatsDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	operation, _ := dao.GetOperationById(dto.OpID)
	if operation.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "OperationId不存在")
		return
	}

	operationDayStats := dao.GetOperationDayStats(dto.OpID, dto.Source, time.Now().Format("2006-01-02"))
	if operationDayStats.ID == 0 {
		_, err = dao.AddOperationDayStats(entity.AddOperationDayStatsDto(dto))
		if err != nil {
			result.Failed(c, int(result.ApiCode.FAILED), "表单不存在，添加失败")
			return
		}
	}
	operationDayStats = dao.UpdateOperationDayStats()
	if operationDayStats.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "dao.UpdateOperationDayStats() 失败")
		return
	}
	result.Success(c, "添加成功")
}

var operationDayStatsService = OperationDayStatsServiceImpl{}

func OperationDayStatsService() IOperationDayStatsService {
	return &operationDayStatsService
}
