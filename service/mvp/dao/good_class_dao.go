package dao

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var GoodClassDao goodClassDao

type goodClassDao struct{}

func (goodClassDao) Get(goodClassName string) (*model.GoodClass, error) {
	createTableWhenNotExist(&model.GoodClass{})
	var goodClass model.GoodClass
	db := mainDB.First(&goodClass, "name = ?", goodClassName)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to get good class",
			zap.Any("goodClassName", goodClassName),
			zap.Error(err))
		return nil, err
	}
	return &goodClass, nil
}

func (goodClassDao) GetAll() ([]*model.GoodClass, error) {
	createTableWhenNotExist(&model.GoodClass{})

	var goodClasses []*model.GoodClass
	db := mainDB.Find(&goodClasses)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to get all good classes",
			zap.Error(err))
		return nil, err
	}
	return goodClasses, nil
}

func (goodClassDao) Create(className string) error {
	createTableWhenNotExist(&model.GoodClass{})
	db := mainDB.Create(&model.GoodClass{
		Name: className,
	})
	if err := db.Error; err != nil {
		logger.Error("Fail to create good class",
			zap.Any("className", className),
			zap.Error(err))
		return err
	}
	return nil
}
