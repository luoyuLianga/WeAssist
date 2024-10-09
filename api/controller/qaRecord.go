package controller

import (
	"WeAssist/api/entity"
	"WeAssist/api/service"
	"github.com/gin-gonic/gin"
)

func AddQaRecord(c *gin.Context) {
	var dto entity.AddQaRecordDto
	_ = c.BindJSON(&dto)
	service.QaRecordService().Add(c, dto)
}
