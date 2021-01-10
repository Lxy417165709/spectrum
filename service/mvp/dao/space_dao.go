package dao

import (
	"github.com/jinzhu/gorm"
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

func (spaceDao) Get(name string, num int64) (*model.Space, error) {
	var table model.Space
	createTableWhenNotExist(&table)

	var result model.Space
	if err := mainDB.First(&result, "name = ? and num = ?", name, num).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.First",
			zap.String("name", name),
			zap.Int64("num", num),
			zap.Error(err))
		return nil, err
	}
	return &result, nil
}
