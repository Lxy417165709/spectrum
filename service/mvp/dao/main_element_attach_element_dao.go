package dao

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var MainElementAttachElementRecordDao mainElementAttachElementRecordDao

type mainElementAttachElementRecordDao struct{}

func (mainElementAttachElementRecordDao) Create(obj *model.MainElementAttachElementRecord) error {
	var table model.MainElementAttachElementRecord
	createTableWhenNotExist(&table)

	if err := mainDB.Create(obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}

func (mainElementAttachElementRecordDao) GetByGoodIdAndMainElementName(goodID int64, mainElementName string) ([]*model.MainElementAttachElementRecord, error) {
	var table model.MainElementAttachElementRecord
	createTableWhenNotExist(&table)
	var result []*model.MainElementAttachElementRecord

	var whereClause string
	var parameters []interface{}
	if goodID == 0 {
		whereClause = "good_id = ? and main_element_name = ?"
		parameters = []interface{}{goodID, mainElementName}
	} else {
		whereClause = "good_id = ?"
		parameters = []interface{}{goodID}
	}

	if err := mainDB.Where(whereClause, parameters...).First(&result).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.String("mainElementName", mainElementName), zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (mainElementAttachElementRecordDao) GetByBothName(goodID int64, mainElementName string, attachElementName string) (*model.MainElementAttachElementRecord, error) {
	var table model.MainElementAttachElementRecord
	createTableWhenNotExist(&table)
	var result model.MainElementAttachElementRecord
	if err := mainDB.First(
		&result,
		"good_id = ? and main_element_name = ? and attach_element_name = ?", goodID, mainElementName, attachElementName,
	).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			logger.Warn("Can't find out attach element",
				zap.Int64("goodID", goodID),
				zap.String("mainElementName", mainElementName),
				zap.String("attachElementName", attachElementName))
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.First",
			zap.Int64("goodID", goodID),
			zap.String("mainElementName", mainElementName),
			zap.String("attachElementName", attachElementName),
			zap.Error(err))
		return nil, err
	}
	return &result, nil
}
