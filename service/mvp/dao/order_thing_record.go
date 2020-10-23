package dao

import (
	"spectrum/service/mvp/model"
)

var OrderThingRecordDao orderThingRecordDao

type orderThingRecordDao struct{}

func (orderThingRecordDao) Create(value *model.OrderThingRecord) error {
	return universalCreate(value)
}
