package service

import (
	"WeAssist/api/dao"
	"WeAssist/api/entity"
	"WeAssist/common/result"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 定义接口
type IActivityService interface {
	AddActivity(c *gin.Context, dto entity.ActivityDto)
}

type ActivityServiceImpl struct{}

// 投票
func (p *ActivityServiceImpl) AddActivity(c *gin.Context, dto entity.ActivityDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	err = dao.AddActivity(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "活动创建失败")
	}

	result.Success(c, "创建成功")
}

var activityService = ActivityServiceImpl{}

func ActivityService() IActivityService {
	return &activityService
}
