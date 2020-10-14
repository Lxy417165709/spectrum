package dao

import (
	"github.com/astaxie/beego/logs"
	"test/service/mvp/model"
)

var SellRecordDao sellRecordDao

type sellRecordDao struct{}

func (sellRecordDao) Create(goodID int, sellPrice float64) error {
	createTableWhenNotExist(&model.SellRecord{})

	db := mainDB.Create(&model.SellRecord{
		GoodID:    goodID,
		SellPrice: sellPrice,
	})
	if err := db.Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
