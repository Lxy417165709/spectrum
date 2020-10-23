package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var OrderThingRecordDao orderThingRecordDao

type orderThingRecordDao struct{}

func (orderThingRecordDao) Create(value *model.OrderThingRecord) error {
	return universalCreate(value)
}

func (orderThingRecordDao) GetByOrderID(orderID int) ([]*model.OrderThingRecord, error) {
	createTableWhenNotExist(&model.OrderThingRecord{})
	var orderThingRecords []*model.OrderThingRecord
	db := mainDB.Find(&orderThingRecords, "order_id = ?", orderID)
	if err := db.Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.Int("orderID", orderID), zap.Error(err))
		return nil, err
	}
	return orderThingRecords, nil
}
