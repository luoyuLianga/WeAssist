package controller

import (
	"WeAssist/api/entity"
	"WeAssist/api/service"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var dto entity.AddUserPluginDto
	_ = c.BindJSON(&dto)
	service.UserPluginService().Add(c, dto)
}
