package service

import (
	"WeAssist/api/dao"
	"WeAssist/api/entity"
	"WeAssist/common/result"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

// IOperationService 定义接口
type IOperationService interface {
	Add(c *gin.Context, dto entity.AddOperationDto)
	Get(c *gin.Context)
	Update(c *gin.Context, dto entity.UpdateOperationDto)
	Delete(c *gin.Context)
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
		return
	}

	result.Success(c, "添加成功")
}

func (q OperationServiceImpl) Get(c *gin.Context) {
	operations, err := dao.GetOperation()
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "GetOperation() Failed")
		return
	}
	result.Success(c, operations)
}

func (q OperationServiceImpl) Update(c *gin.Context, dto entity.UpdateOperationDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	operation, _ := dao.GetOperationById(dto.ID)
	if operation.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "OperationId不存在")
		return
	}

	_, err = dao.UpdateOperation(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "UpdateOperation Failed")
		return
	}

	result.Success(c, "UpdateOperation Success")
}

func (q OperationServiceImpl) Delete(c *gin.Context) {
	// 1. 从路径参数中获取 id，并检查是否存在
	idStr, ok := c.Params.Get("id")
	if !ok {
		result.Failed(c, int(result.ApiCode.FAILED), "DeleteOperation Id Invalid")
		return
	}

	// 2. 将 id 从字符串转换为 uint
	id, err := strconv.ParseUint(idStr, 10, 64) // 64位无符号整数
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "Invalid Id Format")
		return
	}

	err = dao.DeleteOperation(uint(id))
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "DeleteOperation Failed: "+err.Error())
		return
	}

	result.Success(c, fmt.Sprintf("DeleteOperation Success for ID %d", id))
}

var operationService = OperationServiceImpl{}

func OperationService() IOperationService {
	return &operationService
}
