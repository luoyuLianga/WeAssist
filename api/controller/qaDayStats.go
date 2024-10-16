package controller

import (
	"WeAssist/api/service"
	"github.com/gin-gonic/gin"
)

func GetMonthQaDayStats(c *gin.Context) {
	service.QaDayStatsService().GetMonth(c)
}

func GetDayQaDayStats(c *gin.Context) {
	service.QaDayStatsService().GetDay(c)
}
