package dao

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var DeskDao deskDao

type deskDao struct{}

func (deskDao) Create(obj *model.Desk) error {
	if err := mainDB.Create(&obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return ers.MysqlError
	}
	return nil
}

func (deskDao) Update(to map[string]interface{}) error {
	var table model.Desk
	createTableWhenNotExist(&table)
	// todo: 要确定 where 条件，是否是 id == to[id]
	if err := mainDB.Table(table.TableName()).Update(to).Error; err != nil {
		logger.Error("Fail to finish mainDB.Update", zap.Any("to", to), zap.Error(err))
		return ers.MysqlError
	}
	return nil
}

func (deskDao) GetNonCheckOutDesk(spaceID int64) (*model.Desk, error) {
	var result model.Desk
	if err := mainDB.First(&result, "space_id = ? and check_out_at = 0", spaceID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.First",
			zap.Any("spaceID", spaceID),
			zap.Error(err))
		return nil, ers.MysqlError
	}
	return &result, nil
}

func (deskDao) Get(id int64) (*model.Desk, error) {
	var table model.Desk
	createTableWhenNotExist(&table)

	var result model.Desk
	if err := mainDB.First(&result, "id = ?", id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.First", zap.Int64("id", id), zap.Error(err))
		return nil, ers.MysqlError
	}
	return &result, nil
}

func (deskDao) GetByOrderID(orderID int64) (*model.Desk, error) {
	var table model.Desk
	createTableWhenNotExist(&table)

	var result model.Desk
	if err := mainDB.First(&result, "order_id = ?", orderID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.First", zap.Int64("orderID", orderID), zap.Error(err))
		return nil, ers.MysqlError
	}
	return &result, nil
}

//func (deskDao) GetName() string {
//	return model.ChargeableObjectNameOfDesk
//}
