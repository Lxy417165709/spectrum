package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var OrderDao orderDao

type orderDao struct{}

func (orderDao) Create(obj *model.Order) error {
	var table model.Order
	createTableWhenNotExist(&table)

	if err := mainDB.Create(&obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}
