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

func (mainElementSizeRecordDao) GetByGoodIdAndMainElementName(goodID int64, mainElementName string) (*model.MainElementSizeRecord, error) {
	var table model.MainElementSizeRecord
	createTableWhenNotExist(&table)

	var result model.MainElementSizeRecord
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
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.Find", zap.String("mainElementName", mainElementName), zap.Error(err))
		return nil, err
	}
	return &result, nil
}

func (mainElementSizeRecordDao) BatchDelete(goodIDs []int64) error {
	var table model.MainElementSizeRecord
	createTableWhenNotExist(&table)

	// todo: 这里要测试
	if err := mainDB.Where("good_id in (?)", goodIDs).Delete(&table).Error; err != nil {
		logger.Error("Fail to finish mainDB.Delete", zap.Any("goodIDs", goodIDs), zap.Error(err))
		return err
	}
	return nil
}
