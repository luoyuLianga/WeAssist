// 用户服务层
package service

import (
	"WeAssist/api/dao"
	"WeAssist/api/entity"
	"WeAssist/common/result"
	"WeAssist/common/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// IUserService 定义接口
type IUserService interface {
	Register(c *gin.Context, dto entity.UserRegisterDto)
	Login(c *gin.Context, dto entity.UserLoginDto)
	Get(c *gin.Context)
	Update(c *gin.Context, dto entity.UpdateUserDto)
}

type UserServiceImpl struct{}

func (u UserServiceImpl) Update(c *gin.Context, dto entity.UpdateUserDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}
	user, err := dao.GetUserByUserId(dto.ID)
	if user.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "UserID 不存在")
		return
	}
	user, err = dao.UpdateUser(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "UpdateUser Failed")
		return
	}
	result.Success(c, "更新用户成功")
}

func (u UserServiceImpl) Get(c *gin.Context) {
	users, err := dao.GetUser()
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "GetUser() Failed")
		return
	}
	result.Success(c, users)
}

// Register 注册
func (u UserServiceImpl) Register(c *gin.Context, dto entity.UserRegisterDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}
	if dto.Password != dto.ConfirmPassword {
		result.Failed(c, int(result.ApiCode.FAILED), "两次密码不一致")
		return
	}
	// 判断用户是否存在
	userByName := dao.GetUserByUserName(dto.Username)
	if userByName.ID > 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "用户已存在")
		return
	}
	_, err = dao.Register(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "注册失败")
		return
	}
	result.Success(c, "注册成功")
}

// Login 登录
func (u UserServiceImpl) Login(c *gin.Context, dto entity.UserLoginDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}
	user := dao.GetUserByUserName(dto.Username)
	if user.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "用户名不正确")
		return
	}
	// 判断是否存在
	if user.Password != util.EncryptionMd5(dto.Password) {
		result.Failed(c, int(result.ApiCode.FAILED), "密码不正确")
		return
	}
	result.Success(c, user)
}

var userService = UserServiceImpl{}

func UserService() IUserService {
	return &userService
}
