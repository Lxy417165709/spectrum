package dao

import (
	"github.com/jinzhu/gorm"
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

func (goodDao) BatchDelete(ids []int64) error {
	var table model.Good
	createTableWhenNotExist(&table)

	// todo: 这里要测试
	if err := mainDB.Where("id in (?)", ids).Delete(&table).Error; err != nil {
		logger.Error("Fail to finish mainDB.Delete", zap.Any("ids", ids), zap.Error(err))
		return err
	}
	return nil
}

func (goodDao) Get(id int64) (*model.Good, error) {
	var table model.Good
	createTableWhenNotExist(&table)

	var result model.Good
	if err := mainDB.First(&result, "id = ?", id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.First", zap.Int64("id", id), zap.Error(err))
		return nil, err
	}
	return &result, nil
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
