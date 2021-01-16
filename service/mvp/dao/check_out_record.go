package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var CheckOutRecordDao checkOutRecordDao

type checkOutRecordDao struct{}

func (checkOutRecordDao) Create(obj *model.CheckOutRecord) error {
	var table model.CheckOutRecord
	createTableWhenNotExist(&table)

	if err := mainDB.Create(obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}
