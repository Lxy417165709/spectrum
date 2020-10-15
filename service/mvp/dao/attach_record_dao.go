package dao

import (
	"github.com/astaxie/beego/logs"
	"spectrum/service/mvp/model"
)

var AttachRecordDao attachRecordDao

type attachRecordDao struct{}

func (attachRecordDao) Create(orderID, thingID, goodID, attachGoodID int) error {
	createTableWhenNotExist(&model.AttachRecord{})

	db := mainDB.Create(&model.AttachRecord{
		OrderID:      orderID,
		GoodID:       goodID,
		ThingID:      thingID,
		AttachGoodID: attachGoodID,
	})
	if err := db.Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
