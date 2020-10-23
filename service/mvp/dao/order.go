package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var OrderDao orderDao

type orderDao struct{}

func (orderDao) Create(value *model.Order) error {
	if err := universalCreate(value); err != nil {
		logger.Error("Fail to finish universalCreate", zap.Any("value", value), zap.Error(err))
		return err
	}
	return nil
}
func (orderDao) Get(id int) (*model.Order, error) {
	obj, err := universalGet(id, &model.Order{})
	if err != nil {
		logger.Error("Fail to finish universalGet",
			zap.Any("id", id),
			zap.Error(err))
		return nil, err
	}
	if obj == nil{
		return nil,nil
	}
	return obj.(*model.Order), nil
}

