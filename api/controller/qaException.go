package controller

import (
	"WeAssist/api/entity"
	"WeAssist/api/service"
	"github.com/gin-gonic/gin"
)

func AddQaException(c *gin.Context) {
	var dto entity.AddQaExceptionDto
	_ = c.BindJSON(&dto)
	service.QaExceptionService().Add(c, dto)
}
