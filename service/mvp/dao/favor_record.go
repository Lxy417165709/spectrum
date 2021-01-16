package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)


var FavorRecordDao favorRecordDao

type favorRecordDao struct{}

func (favorRecordDao) Create(obj *model.FavorRecord) error {
	var table model.FavorRecord
	createTableWhenNotExist(&table)

	if err := mainDB.Create(obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}
