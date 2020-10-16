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
	}
	return nil
}
