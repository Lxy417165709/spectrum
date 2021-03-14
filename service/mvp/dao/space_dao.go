package dao

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var SpaceDao spaceDao

type spaceDao struct{}

func (spaceDao) Create(obj *model.Space) error {
	var table model.Space
	createTableWhenNotExist(&table)

	if err := mainDB.Create(obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}

func (spaceDao) GetAll() ([]*model.Space, error) {
	var table model.Space
	createTableWhenNotExist(&table)

	var result []*model.Space
	if err := mainDB.Find(&result).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find",
			zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (spaceDao) GetByClassName(className string) ([]*model.Space, error) {
	var table model.Space
	createTableWhenNotExist(&table)

	var result []*model.Space
	if err := mainDB.Where("class_name = ?", className).Find(&result).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find",
			zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (spaceDao) Get(name string, className string) (*model.Space, error) {
	var table model.Space
	createTableWhenNotExist(&table)

	var result model.Space
	if err := mainDB.First(&result, "name = ? and class_name = ?", name, className).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.First",
			zap.String("name", name),
			zap.String("className", className),
			zap.Error(err))
		return nil, err
	}
	return &result, nil
}
