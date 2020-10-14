package dao

import (
	"github.com/astaxie/beego/logs"
	"test/service/mvp/model"
)

var OrderRecordDao orderRecordDao

type orderRecordDao struct{}

func (orderRecordDao) Create(orderID int, goodID int, isAttachGood int) error {
	createTableWhenNotExist(&model.OrderRecord{})

	db := mainDB.Create(&model.OrderRecord{
		OrderID:      orderID,
		GoodID:       goodID,
		IsAttachGood: isAttachGood,
	})
	if err := db.Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (orderRecordDao) GetByOrderID(orderID int) ([]*model.OrderRecord, error) {
	createTableWhenNotExist(&model.OrderRecord{})
	var orderRecords []*model.OrderRecord
	db := mainDB.Find(&orderRecords, "order_id = ?", orderID)
	if err := db.Error; err != nil {
		logs.Error(err)
		return orderRecords, err
	}
	return orderRecords, nil
}
