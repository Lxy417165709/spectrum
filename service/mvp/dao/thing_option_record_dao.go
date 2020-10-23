package dao

import (
	"spectrum/service/mvp/model"
)

var ThingOptionRecordDao thingOptionRecordDao

type thingOptionRecordDao struct{}

func (thingOptionRecordDao) Create(value *model.ThingOptionRecord) error {
	return universalCreate(value)
}
