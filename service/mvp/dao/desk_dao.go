package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
	"spectrum/service/mvp/utils"
	"time"
)

var DeskDao deskDao

type deskDao struct{}

func (deskDao) Create(obj *model.Desk) (int64, error) {
	values := []interface{}{
		obj.ID, obj.Expense, obj.CheckOutAt, obj.NonFavorExpense, obj.StartAt, obj.EndAt, obj.SpaceID, obj.OrderID,
	}
	sql := fmt.Sprintf(`
		insert into %s(id,expense,check_out_at,non_favor_expense,start_at,end_at,space_id,order_id) values(%s)
		on duplicate key update
			expense = values(expense),
			check_out_at = values(check_out_at),
			non_favor_expense = values(non_favor_expense),
			start_at = values(start_at),
			end_at = values(end_at),
			space_id = values(space_id),
			order_id = values(order_id);
	`, fmt.Sprintf("`%s`", obj.TableName()), getPlaceholderClause(len(values)))
	result, err := mainDB.CommonDB().Exec(sql, values...)
	if err != nil {
		logger.Error("Fail to finish create", zap.Any("obj", obj), zap.Error(err))
		return 0, ers.MysqlError
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Fail to get id", zap.Any("obj", obj), zap.Error(err))
		return 0, ers.MysqlError
	}
	return id, nil
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
	if err := mainDB.First(&result, "space_id = ? and check_out_at = ?", spaceID, utils.NilTime).Error; err != nil {
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

func (deskDao) Get(id, spaceID int64) (*model.Desk, error) {
	var result model.Desk
	if err := mainDB.First(&result, "id = ? and space_id = ?", id, spaceID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.First", zap.Int64("id", id),
			zap.Int64("spaceID", spaceID), zap.Error(err))
		return nil, ers.MysqlError
	}
	return &result, nil
}

func (deskDao) GetByID(id int64) (*model.Desk, error) {
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

func (deskDao) CheckOut(id int64, nonFavorExpense float64, checkOutAt time.Time, expense float64,
	endAt time.Time) error {
	sql := fmt.Sprintf("update %s set non_favor_expense = ?,check_out_at = ?,expense = ?,end_at = ? where id = ?",
		(&model.Desk{}).TableName())
	if _, err := mainDB.CommonDB().Exec(sql, nonFavorExpense, checkOutAt, expense, endAt, id); err != nil {
		logger.Error("Fail to finish check out", zap.Any("id", id), zap.Error(err))
		return ers.MysqlError
	}
	return nil
}
