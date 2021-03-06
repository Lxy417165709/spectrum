package dao

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var DeskDao deskDao

type deskDao struct{}

func (deskDao) Create(obj *model.Desk) error {
	var table model.Desk
	createTableWhenNotExist(&table)

	if err := mainDB.Create(&obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}

func (deskDao) Update(to map[string]interface{}) error {
	var table model.Desk
	createTableWhenNotExist(&table)
	// todo: 要确定 where 条件，是否是 id == to[id]
	if err := mainDB.Table(table.TableName()).Update(to).Error; err != nil {
		logger.Error("Fail to finish mainDB.Update", zap.Any("to", to), zap.Error(err))
		return err
	}
	return nil
}

func (deskDao) GetNonCheckOutDesk(spaceName string, spaceNum int64) (*model.Desk, error) {
	var table model.Desk
	createTableWhenNotExist(&table)

	var result model.Desk
	if err := mainDB.First(&result, "space_name = ? and space_num = ? and check_out_timestamp = 0", spaceName, spaceNum).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.First",
			zap.String("spaceName", spaceName),
			zap.Int64("spaceNum", spaceNum),
			zap.Error(err))
		return nil, err
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
		return nil, err
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
		return nil, err
	}
	return &result, nil
}

//func (deskDao) GetName() string {
//	return model.ChargeableObjectNameOfDesk
//}
