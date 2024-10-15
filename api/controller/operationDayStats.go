package controller

import (
	"WeAssist/api/entity"
	"WeAssist/api/service"
	"github.com/gin-gonic/gin"
)

func UpdateOperationDayStats(c *gin.Context) {
	var dto entity.OperationDayStatsDto
	_ = c.BindJSON(&dto)
	service.OperationDayStatsService().Update(c, dto)
}

func GetMonthOperationDayStats(c *gin.Context) {
	service.OperationDayStatsService().GetMonth(c)
}
