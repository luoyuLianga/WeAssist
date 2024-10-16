package service

import (
	"WeAssist/api/dao"
	"WeAssist/api/entity"
	"WeAssist/common/result"
	"WeAssist/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"time"
)

// IOperationDayStatsService 定义接口
type IOperationDayStatsService interface {
	Update(c *gin.Context, dto entity.OperationDayStatsDto)
	GetMonth(c *gin.Context)
	GetDay(c *gin.Context)
}

type OperationDayStatsServiceImpl struct{}

func (ods OperationDayStatsServiceImpl) GetDay(c *gin.Context) {
	startDay := c.Query("startDay")
	log.Log().Infof("startDay:%s", startDay)

	var dto entity.GetDayODSReqDto
	if err := c.ShouldBindQuery(&dto); err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "GetDayOperationDayStats() Failed")
		return
	}

	log.Log().Infof("dto:%v", dto)
	getDayODSDto, err := dao.GetDayOperationDayStats(dto.StartDay, dto.EndDay)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "GetDayOperationDayStats() Failed")
		return
	}
	result.Success(c, getDayODSDto)
}

func (ods OperationDayStatsServiceImpl) GetMonth(c *gin.Context) {
	// 获取当前时间
	now := time.Now()
	// 计算前11个月的第一天
	startDate := time.Date(now.Year(), now.Month()-11, 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
	// 当前月的最后一天
	endDate := time.Date(now.Year(), now.Month()+1, 0, 23, 59, 59, 0, now.Location()).Format("2006-01-02")

	log.Log().Infof("startDate:%s endDate:%s", startDate, endDate)
	getMonthODSDto, err := dao.GetMonthOperationDayStats(startDate, endDate)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "GetMonthOperationDayStats() Failed")
		return
	}
	result.Success(c, getMonthODSDto)
}

func (ods OperationDayStatsServiceImpl) Update(c *gin.Context, dto entity.OperationDayStatsDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	// 校验OpID字段
	operation, _ := dao.GetOperationById(dto.OpID)
	if operation.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "OperationId不存在")
		return
	}

	// 根据OpID、Source、Day查询记录是否存在，如果不存在需要进行新增
	operationDayStats := dao.GetOperationDayStats(dto, time.Now().Format("2006-01-02"))
	if operationDayStats.ID == 0 {
		_, err = dao.AddOperationDayStats(dto)
		if err != nil {
			result.Failed(c, int(result.ApiCode.FAILED), "表单不存在，添加失败")
			return
		}
	}

	// 更新OpID、Source、Day记录下的count，增加1
	_, err = dao.UpdateOperationDayStats(dto, time.Now().Format("2006-01-02"))
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "UpdateOperationDayStats 失败")
	}
	result.Success(c, "添加成功")
}

var operationDayStatsService = OperationDayStatsServiceImpl{}

func OperationDayStatsService() IOperationDayStatsService {
	return &operationDayStatsService
}
