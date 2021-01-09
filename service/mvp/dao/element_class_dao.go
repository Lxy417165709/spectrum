package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var ElementClassDao elementClassDao

type elementClassDao struct{}

func (elementClassDao) Create(obj *model.ElementClass) error {
	var table model.ElementClass
	createTableWhenNotExist(&table)

	if err := mainDB.Create(obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}

func (elementClassDao) GetAllClasses() ([]*model.ElementClass, error) {
	var table model.ElementClass
	createTableWhenNotExist(&table)
	var result []*model.ElementClass
	if err := mainDB.Find(&result).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.Error(err))
		return nil, err
	}
	return result, nil
}
