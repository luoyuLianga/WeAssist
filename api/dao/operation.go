package dao

import (
	"WeAssist/api/entity"
	"WeAssist/common/util"
	"WeAssist/pkg/db"
	"time"
)

// AddOperation 添加操作
func AddOperation(dto entity.AddOperationDto) (uint, error) {
	operation := entity.Operation{
		OperationCode: dto.OperationCode,
		OperationDesc: dto.OperationDesc,
		CreateTime:    util.HTime{Time: time.Now()},
		UpdateTime:    util.HTime{Time: time.Now()},
	}
	err := db.Db.Create(&operation).Error
	return operation.ID, err
}

// GetOperationByCode 根据OperationCode查询
func GetOperationByCode(operationCode string) (operation entity.Operation) {
	db.Db.Where("operation_code = ?", operationCode).First(&operation)
	return operation
}
