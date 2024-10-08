package controller

import (
	"WeAssist/api/entity"
	"WeAssist/api/service"

	"github.com/gin-gonic/gin"
)

// 创建活动
// @Summary 创建活动接口
// @Tags      Activity
// @Produce json
// @Description 创建活动接口
// @Param data body entity.ActivityDto true "data"
// @Success 200 {object} result.Result
// @router /api/activity/add [post]
func AddActivity(c *gin.Context) {
	var dto entity.ActivityDto
	_ = c.BindJSON(&dto)
	service.ActivityService().AddActivity(c, dto)
}
