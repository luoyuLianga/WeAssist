package controller

import (
	"WeAssist/api/entity"
	"WeAssist/api/service"
	"github.com/gin-gonic/gin"
)

func AddOperationDayStats(c *gin.Context) {
	var dto entity.AddOperationDayStatsDto
	_ = c.BindJSON(&dto)
	service.OperationDayStatsService().Add(c, dto)
}

func UpdateOperationDayStats(c *gin.Context) {
	var dto entity.UpdateOperationDayStatsDto
	_ = c.BindJSON(&dto)
	service.OperationDayStatsService().Update(c, dto)
}
