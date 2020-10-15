package dao

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"spectrum/service/mvp/model"
)

var GoodTypeDao goodTypeDao

type goodTypeDao struct{}

func (goodTypeDao) GetByName(goodTypeName string) (*model.GoodType, error) {
	createTableWhenNotExist(&model.GoodType{})

	var goodType model.GoodType
	db := mainDB.First(&goodType, "name = ?", goodTypeName)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logs.Error(err)
		return nil, err
	}
	return &goodType, nil
}

func (goodTypeDao) Create(name string) error {
	createTableWhenNotExist(&model.GoodType{})

	db := mainDB.Create(&model.GoodType{
		Name: name,
	})
	if err := db.Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
