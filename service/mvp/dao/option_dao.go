package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var OptionDao optionDao

type optionDao struct{}

func (optionDao) Create(optionClassID int, optionName string) error {
	createTableWhenNotExist(&model.Option{})
	db := mainDB.Create(&model.Option{
		OptionClassID: optionClassID,
		Name:          optionName,
	})
	if err := db.Error; err != nil {
		logger.Error("Fail to create option",
			zap.Any("optionClassID", optionClassID),
			zap.Any("optionName", optionName),
			zap.Error(err))
		return err
	}
	return nil
}

func (optionDao) GetByOptionClassID(optionClassID int) ([]*model.Option, error) {
	createTableWhenNotExist(&model.Option{})
	var options []*model.Option
	db := mainDB.Find(&options, "option_class_id = ?", optionClassID)
	if err := db.Error; err != nil {
		logger.Error("Fail to get options",
			zap.Any("optionClassID", optionClassID),
			zap.Error(err))
		return nil, err
	}
	return options, nil
}

func (optionDao) Del(optionClassID int, optionName string) error {
	createTableWhenNotExist(&model.Option{})
	db := mainDB.Delete(&model.Option{}, "name = ? and option_class_id = ?", optionName, optionClassID)
	if err := db.Error; err != nil {
		logger.Error("Fail to del option",
			zap.Any("optionClassID", optionClassID),
			zap.Any("optionName", optionName),
			zap.Error(err))
		return err
	}
	return nil
}
