package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var ThingOptionRecordDao thingOptionRecordDao

type thingOptionRecordDao struct{}

func (thingOptionRecordDao) Create(value *model.ThingOptionRecord) error {
	if err := universalCreate(value); err != nil {
		logger.Error("Fail to finish universalCreate", zap.Any("value", value), zap.Error(err))
		return err
	}
	return nil
}

func (thingOptionRecordDao) GetByThingID(thingID int) ([]*model.ThingOptionRecord, error) {
	createTableWhenNotExist(&model.ThingOptionRecord{})

	var thingOptionRecords []*model.ThingOptionRecord
	db := mainDB.Find(&thingOptionRecords, "thing_id = ?", thingID)
	if err := db.Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.Int("thingID", thingID), zap.Error(err))
		return nil, err
	}
	return thingOptionRecords, nil
}
