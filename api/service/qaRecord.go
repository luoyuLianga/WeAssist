package service

import (
	"WeAssist/api/dao"
	"WeAssist/api/entity"
	"WeAssist/common/constant"
	"WeAssist/common/result"
	"WeAssist/pkg/log"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// IQaRecordService 定义接口
type IQaRecordService interface {
	Add(c *gin.Context, dto entity.AddQaRecordDto)
}

type QaRecordServiceImpl struct{}

func (q QaRecordServiceImpl) Add(c *gin.Context, dto entity.AddQaRecordDto) {
	log.Log().Infof("Add dto:%v", dto)

	// 1. 基本参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	// 2. UserPluginID有效性校验
	userPlugin, _ := dao.GetUserPluginById(dto.UserPluginID)
	if userPlugin.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "用户插件不存在")
		return
	}

	// 3. QA、Accept场景参数校验
	err = ProcessQaRecord(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), err.Error())
		return
	}

	_, err = dao.AddQaRecord(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "添加失败")
	}

	result.Success(c, "添加成功")
}

// ProcessQaRecord 处理 QA 记录的业务逻辑
func ProcessQaRecord(dto entity.AddQaRecordDto) error {
	switch dto.Type {
	case constant.QA:
		// 处理 QA 类型的逻辑
		return CheckQaParam(dto)
	case constant.Accept:
		// 处理 Accept 类型的逻辑
		return CheckAcceptParam(dto)
	case constant.Reject:
		// 处理 Reject 类型的逻辑  TODO Reject下暂不进行校验
		return nil
	default:
		return errors.New("invalid record type")
	}
}

func CheckQaParam(dto entity.AddQaRecordDto) error {
	if dto.UserQuestion == "" || dto.ModelResponse == "" {
		return errors.New("user_question and model_response cannot be empty for QA type")
	}
	return nil
}

func CheckAcceptParam(dto entity.AddQaRecordDto) error {
	if dto.CodeNumber == 0 {
		return errors.New("code_number cannot be empty for Accept type")
	}
	return nil
}

var qaRecordService = QaRecordServiceImpl{}

func QaRecordService() IQaRecordService {
	return &qaRecordService
}
