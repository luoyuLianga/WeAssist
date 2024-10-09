package controller

import (
	"WeAssist/api/entity"
	"WeAssist/api/service"
	"github.com/gin-gonic/gin"
)

func AddOperation(c *gin.Context) {
	var dto entity.AddOperationDto
	_ = c.BindJSON(&dto)
	service.OperationService().Add(c, dto)
}
