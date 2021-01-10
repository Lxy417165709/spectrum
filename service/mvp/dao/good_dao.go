package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var GoodDao goodDao

type goodDao struct{}

func (goodDao) Create(obj *model.Good) error {
	var table model.Good
	createTableWhenNotExist(&table)

	if err := mainDB.Create(&obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}

func (goodDao) GetByDeskID(deskID int64) ([]*model.Good, error) {
	var table model.Good
	createTableWhenNotExist(&table)

	var result []*model.Good
	if err := mainDB.Find(&result, "desk_id = ?", deskID).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.Int64("deskID", deskID), zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (goodDao) Update(to map[string]interface{}) error {
	var table model.Good
	createTableWhenNotExist(&table)

	// todo: 要确定 where 条件，是否是 id == to[id]
	if err := mainDB.Table(table.TableName()).Update(to).Error; err != nil {
		logger.Error("Fail to finish mainDB.Update", zap.Any("to", to), zap.Error(err))
		return err
	}
	return nil
}