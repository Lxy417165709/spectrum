package dao

import (
	"spectrum/service/mvp/model"
)

var OrderDao orderDao

type orderDao struct{}

func (orderDao) Create(value *model.Order) error {
	return universalCreate(value)
}
