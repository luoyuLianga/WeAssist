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

func GetOperation(c *gin.Context) {
	service.OperationService().Get(c)
}

func UpdateOperation(c *gin.Context) {
	var dto entity.UpdateOperationDto
	_ = c.BindJSON(&dto)
	service.OperationService().Update(c, dto)
}

func DeleteOperation(c *gin.Context) {
	service.OperationService().Delete(c)
}
