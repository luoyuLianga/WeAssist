package service

import (
	"WeAssist/api/dao"
	"WeAssist/api/entity"
	"WeAssist/common/result"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// IUserPluginService 定义接口
type IUserPluginService interface {
	Add(c *gin.Context, dto entity.AddUserPluginDto)
}

type UserPluginServiceImpl struct{}

func (u UserPluginServiceImpl) Add(c *gin.Context, dto entity.AddUserPluginDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	user, _ := dao.GetUserByUserId(dto.UserId)
	if user.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "用户不存在")
		return
	}

	userPlugin, _ := dao.GetUserPluginByPM(dto.PluginName, dto.ModelName)
	if userPlugin.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "PluginName and ModelName 已存在")
		return
	}

	_, err = dao.AddUserPlugin(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), err.Error())
	}

	result.Success(c, "添加成功")
}

var userPluginService = UserPluginServiceImpl{}

func UserPluginService() IUserPluginService {
	return &userPluginService
}
