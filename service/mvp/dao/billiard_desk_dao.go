package dao

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"test/service/mvp/model"
)

var BilliardDeskDao billiardDeskDao

type billiardDeskDao struct{}

func (billiardDeskDao) GetByName(goodName string) (*model.BilliardDesk, error) {
	createTableWhenNotExist(&model.BilliardDesk{})

	var desk model.BilliardDesk
	db := mainDB.First(&desk, "name = ?", goodName)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logs.Error(err)
		return nil, err
	}
	return &desk, nil
}

func (billiardDeskDao) Create(deskName string) error {
	createTableWhenNotExist(&model.BilliardDesk{})

	db := mainDB.Create(&model.BilliardDesk{
		Name: deskName,
	})
	if err := db.Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
