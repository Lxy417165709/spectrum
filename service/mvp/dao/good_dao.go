package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var GoodDao goodDao

type goodDao struct{}

func (goodDao) Create(obj *model.Good) (int64, error) {
	values := []interface{}{
		obj.ID, obj.Expense, obj.CheckOutAt, obj.NonFavorExpense, obj.MainElementID, obj.OrderID,
	}
	sql := fmt.Sprintf(`
		insert into %s(id,expense,check_out_at,non_favor_expense,main_element_id,order_id) values(%s)
		on duplicate key update
			expense = values(expense),
			check_out_at = values(check_out_at),
			non_favor_expense = values(non_favor_expense),
			main_element_id = values(main_element_id),
			order_id = values(order_id);
	`, fmt.Sprintf("`%s`", obj.TableName()), GetPlaceholderClause(len(values)))
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

func (goodDao) GetByOrderID(orderID int64) ([]*model.Good, error) {
	var table model.Good
	createTableWhenNotExist(&table)

	var result []*model.Good
	if err := mainDB.Find(&result, "order_id = ?", orderID).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.Int64("orderID", orderID), zap.Error(err))
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
