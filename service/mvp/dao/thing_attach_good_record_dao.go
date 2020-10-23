package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var ThingAttachGoodRecordDao thingAttachGoodRecordDao

type thingAttachGoodRecordDao struct{}

func (thingAttachGoodRecordDao) Create(value *model.ThingAttachGoodRecord) error {
	if err := universalCreate(value); err != nil {
		logger.Error("Fail to finish universalCreate", zap.Any("value", value), zap.Error(err))
		return err
	}
	return nil
}

func (thingAttachGoodRecordDao)GetByThingID(thingID int) ([]*model.ThingAttachGoodRecord,error) {
	createTableWhenNotExist(&model.ThingAttachGoodRecord{})

	var thingAttachGoodRecords []*model.ThingAttachGoodRecord
	db := mainDB.Find(&thingAttachGoodRecords, "thing_id = ?", thingID)
	if err := db.Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.Int("thingID", thingID), zap.Error(err))
		return nil, err
	}
	return thingAttachGoodRecords, nil
}
