package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var DeskClassDao deskClassDao

type deskClassDao struct{}

func (deskClassDao) Create(obj *model.DeskClass) error {
	var table model.DeskClass
	createTableWhenNotExist(&table)

	if err := mainDB.Create(obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}

func (deskClassDao) GetAllClasses() ([]*model.DeskClass, error) {
	var table model.DeskClass
	createTableWhenNotExist(&table)
	var result []*model.DeskClass
	if err := mainDB.Find(&result).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.Error(err))
		return nil, err
	}
	return result, nil
}
