package controller

import (
	"WeAssist/api/entity"
	"WeAssist/api/service"

	"github.com/gin-gonic/gin"
)

// 用户注册
// @Summary 用户注册接口
// @Tags      User
// @Produce json
// @Description 用户注册接口
// @Param data body entity.UserRegisterDto true "data"
// @Success 200 {object} result.Result
// @router /api/user/register [post]
func Register(c *gin.Context) {
	var dto entity.UserRegisterDto
	_ = c.BindJSON(&dto)
	service.UserService().Register(c, dto)
}

// 用户登录
// @Summary 用户登录接口
// @Tags      User
// @Produce json
// @Description 用户登录接口
// @Param data body entity.UserLoginDto true "data"
// @Success 200 {object} result.Result
// @router /api/user/login [post]
func Login(c *gin.Context) {
	var dto entity.UserLoginDto
	_ = c.BindJSON(&dto)
	service.UserService().Login(c, dto)
}

func GetUser(c *gin.Context) {
	service.UserService().Get(c)
}

func UpdateUser(c *gin.Context) {
	var dto entity.UpdateUserDto
	_ = c.BindJSON(&dto)
	//service.UserService().Update(c)
}
