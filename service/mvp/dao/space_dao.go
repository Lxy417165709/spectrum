package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var SpaceDao spaceDao

type spaceDao struct{}

func (spaceDao) Create(obj *model.Space) error {
	var table model.Space
	createTableWhenNotExist(&table)

	if err := mainDB.Create(obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}
