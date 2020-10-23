package dao

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var ThingDao thingDao

type thingDao struct{}

func (thingDao) Create(value *model.Thing) error {
	if err := universalCreate(value); err != nil {
		logger.Error("Fail to finish universalCreate", zap.Any("value", value), zap.Error(err))
		return err
	}
	return nil
}

func (thingDao) Get(id int) (*model.Thing, error) {
	createTableWhenNotExist(&model.Thing{})
	var thing model.Thing
	db := mainDB.First(&thing, "id = ?", id)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.first", zap.Int("id", id), zap.Error(err))
		return nil, err
	}
	return &thing, nil
}
