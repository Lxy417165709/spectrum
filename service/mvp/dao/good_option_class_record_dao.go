package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var GoodOptionClassRecordDao goodOptionClassRecordDao

type goodOptionClassRecordDao struct{}

func (goodOptionClassRecordDao) Create(goodID, optionClassID int) error {
	createTableWhenNotExist(&model.GoodOptionClassRecord{})

	db := mainDB.Create(&model.GoodOptionClassRecord{
		GoodID:        goodID,
		OptionClassID: optionClassID,
	})
	if err := db.Error; err != nil {
		logger.Error("Fail to create",
			zap.Any("goodID", goodID),
			zap.Any("optionClassID", optionClassID))
		return err
	}
	return nil
}
