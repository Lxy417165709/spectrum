package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var OrderDao orderDao

type orderDao struct{}

func (orderDao) Get(id int64) (*model.Order, error) {
	var result model.Order
	if err := mainDB.First(&result, "id = ?", id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.First", zap.Int64("id", id), zap.Error(err))
		return nil, ers.MysqlError
	}
	return &result, nil
}

func (orderDao) Query(parameter *model.QueryOrderParameter) ([]*model.Order, error) {
	var result []*model.Order
	whereClause, whereValues := parameter.GetWhereClauseAndValues()
	if err := mainDB.Where(whereClause, whereValues...).Find(&result).Error; err != nil {
		logger.Error("Fail to finish mainDB.First", zap.Any("parameter", parameter), zap.Error(err))
		return nil, ers.MysqlError
	}
	return result, nil
}

func (orderDao) Create(obj *model.Order) (int64, error) {
	values := []interface{}{
		obj.ID, obj.Expense, obj.CheckOutAt, obj.NonFavorExpense,
	}
	sql := fmt.Sprintf(`
		insert into %s(id, expense, check_out_at, non_favor_expense) values(%s)
		on duplicate key update
			expense = values(expense),
			check_out_at = values(check_out_at),
			non_favor_expense = values(non_favor_expense);
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
