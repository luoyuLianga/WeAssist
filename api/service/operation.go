package service

import (
	"WeAssist/api/dao"
	"WeAssist/api/entity"
	"WeAssist/common/result"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// IOperationService 定义接口
type IOperationService interface {
	Add(c *gin.Context, dto entity.AddOperationDto)
}

type OperationServiceImpl struct{}

func (q OperationServiceImpl) Add(c *gin.Context, dto entity.AddOperationDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	operation := dao.GetOperationByCode(dto.OperationCode)
	if operation.ID > 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "操作Code已存在")
		return
	}

	_, err = dao.AddOperation(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "添加失败")
	}

	result.Success(c, "添加成功")
}

var operationService = OperationServiceImpl{}

func OperationService() IOperationService {
	return &operationService
}
