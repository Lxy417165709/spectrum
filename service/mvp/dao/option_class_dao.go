package dao

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var OptionClassDao optionClassDao

type optionClassDao struct{}

func (optionClassDao) Get(className string) (*model.OptionClass, error) {
	createTableWhenNotExist(&model.OptionClass{})
	var optionClass model.OptionClass
	db := mainDB.First(&optionClass, "name = ?", className)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to get option class",
			zap.Any("className", className),
			zap.Error(err))
		return nil, err
	}
	return &optionClass, nil
}

func (optionClassDao) GetByID(id int) (*model.OptionClass, error) {
	createTableWhenNotExist(&model.OptionClass{})
	var optionClass model.OptionClass
	db := mainDB.First(&optionClass, "id = ?", id)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to get option class",
			zap.Any("id", id),
			zap.Error(err))
		return nil, err
	}
	return &optionClass, nil
}

func (optionClassDao) Create(className string) error {
	createTableWhenNotExist(&model.OptionClass{})
	db := mainDB.Create(&model.OptionClass{
		Name: className,
	})
	if err := db.Error; err != nil {
		logger.Error("Fail to create option class",
			zap.Any("className", className),
			zap.Error(err))
		return err
	}
	return nil
}

func (optionClassDao) GetAll() ([]*model.OptionClass, error) {
	createTableWhenNotExist(&model.OptionClass{})

	var optionClasses []*model.OptionClass
	db := mainDB.Find(&optionClasses)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to get all option classes",
			zap.Error(err))
		return nil, err
	}
	return optionClasses, nil
}

func (optionClassDao) DeleteByNames(classNames []string) error {
	createTableWhenNotExist(&model.OptionClass{})

	db := mainDB.Delete(&model.OptionClass{}, "name in (?)", classNames)
	if err := db.Error; err != nil {
		logger.Error("Fail to delete option class",
			zap.Any("classNames", classNames),
			zap.Error(err))
		return err
	}
	return nil
}
