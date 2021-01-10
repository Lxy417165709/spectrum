package dao

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var MainElementSizeRecordDao mainElementSizeRecordDao

type mainElementSizeRecordDao struct{}

func (mainElementSizeRecordDao) Create(obj *model.MainElementSizeRecord) error {
	var table model.MainElementSizeRecord
	createTableWhenNotExist(&table)

	if err := mainDB.Create(obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}

func (mainElementSizeRecordDao) GetByMainElementName(mainElementName string) (*model.MainElementSizeRecord, error) {
	var table model.MainElementSizeRecord
	createTableWhenNotExist(&table)
	var result model.MainElementSizeRecord
	if err := mainDB.First(&result, "good_id = ? and main_element_name = ?", 0, mainElementName).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.Find", zap.String("mainElementName", mainElementName), zap.Error(err))
		return nil, err
	}
	return &result, nil
}

func (mainElementSizeRecordDao) GetByGoodID(goodID int64) (*model.MainElementSizeRecord, error) {
	var table model.MainElementSizeRecord
	createTableWhenNotExist(&table)
	var result model.MainElementSizeRecord
	if err := mainDB.First(&result, "good_id = ? ", goodID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.First", zap.Int64("goodID", goodID), zap.Error(err))
		return nil, err
	}
	return &result, nil
}
