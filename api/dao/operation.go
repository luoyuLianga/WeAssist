package dao

import (
	"WeAssist/api/entity"
	"WeAssist/pkg/db"
	"fmt"
)

// AddOperation 添加操作
func AddOperation(dto entity.AddOperationDto) (uint, error) {
	operation := entity.Operation{
		OperationCode: dto.OperationCode,
		OperationDesc: dto.OperationDesc,
	}
	err := db.Db.Create(&operation).Error
	return operation.ID, err
}

// GetOperationById 根据OperationId查询用户
func GetOperationById(id uint) (operation entity.Operation, err error) {
	err = db.Db.Where("id = ?", id).First(&operation).Error
	return operation, err
}

// GetOperationByCode 根据OperationCode查询
func GetOperationByCode(operationCode string) (operation entity.Operation) {
	db.Db.Where("operation_code = ?", operationCode).First(&operation)
	return operation
}

// GetOperation 查询
func GetOperation() (operations []entity.Operation, err error) {
	err = db.Db.Find(&operations).Error
	return operations, err
}

// UpdateOperation 查询
func UpdateOperation(dto entity.UpdateOperationDto) (operations []entity.Operation, err error) {
	err = db.Db.Model(&operations).Where("id = ?", dto.ID).
		Updates(map[string]interface{}{
			"operation_code": dto.OperationCode,
			"operation_desc": dto.OperationDesc,
		}).Error
	return operations, err
}

// DeleteOperation 删除
func DeleteOperation(id uint) (err error) {
	// 删除指定ID的记录
	result := db.Db.Delete(&entity.Operation{}, id)

	// 检查是否发生错误
	if result.Error != nil {
		return result.Error
	}

	// 如果没有任何记录被删除，则返回提示信息
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到ID为 %d 的记录", id)
	}

	return nil // 删除成功
}
