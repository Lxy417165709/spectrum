package dao

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var GoodAttachClassRecordDao goodAttachClassRecordDao

type goodAttachClassRecordDao struct{}

func (goodAttachClassRecordDao) Create(goodID, classID int) error {
	createTableWhenNotExist(&model.GoodAttachClassRecord{})

	db := mainDB.Create(&model.GoodAttachClassRecord{
		GoodID:            goodID,
		AttachGoodClassID: classID,
	})
	if err := db.Error; err != nil {
		logger.Error("Fail to create",
			zap.Any("goodID", goodID),
			zap.Any("attachGoodClassID", classID))
		return err
	}
	return nil
}

func (goodAttachClassRecordDao) GetByGoodID(goodID int) ([]*model.GoodAttachClassRecord, error) {
	createTableWhenNotExist(&model.GoodAttachClassRecord{})

	var goodAttachClassRecords []*model.GoodAttachClassRecord
	db := mainDB.Find(&goodAttachClassRecords, "good_id = ?", goodID)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.Find",
			zap.Any("goodID", goodID),
			zap.Error(err))
		return nil, err
	}
	return goodAttachClassRecords, nil
}
