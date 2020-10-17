package dao

import (
	"github.com/jinzhu/gorm"
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


func (goodOptionClassRecordDao) GetByGoodID(goodID int) ([]*model.GoodOptionClassRecord,error){
	createTableWhenNotExist(&model.GoodOptionClassRecord{})

	var goodOptionClassRecords []*model.GoodOptionClassRecord
	db := mainDB.Find(&goodOptionClassRecords,"good_id = ?",goodID)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.Find",
			zap.Any("goodID",goodID),
			zap.Error(err))
		return nil, err
	}
	return goodOptionClassRecords, nil
}
