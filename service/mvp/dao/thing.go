package dao

import (
	"spectrum/service/mvp/model"
)

var ThingDao thingDao

type thingDao struct{}

func (thingDao) Create(value *model.Thing) error {
	return universalCreate(value)
}
