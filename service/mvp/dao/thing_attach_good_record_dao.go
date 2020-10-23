package dao

import (
	"spectrum/service/mvp/model"
)

var ThingAttachGoodRecordDao thingAttachGoodRecordDao

type thingAttachGoodRecordDao struct{}

func (thingAttachGoodRecordDao) Create(value *model.ThingAttachGoodRecord) error {
	return universalCreate(value)
}
