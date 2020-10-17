package dao

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"spectrum/service/mvp/model"
)

var GoodDao goodDao

type goodDao struct{}

func (goodDao) Get(goodID int) (*model.Good, error) {
	createTableWhenNotExist(&model.Good{})

	var good model.Good
	db := mainDB.First(&good, "id = ?", goodID)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logs.Error(err)
		return nil, err
	}
	return &good, nil
}

func (goodDao) GetByName(goodName string) (*model.Good, error) {
	createTableWhenNotExist(&model.Good{})

	var good model.Good
	db := mainDB.First(&good, "name = ?", goodName)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logs.Error(err)
		return nil, err
	}
	return &good, nil
}

func (goodDao) Create(name string, price float64, goodType int,pictureStorePath string) error {
	createTableWhenNotExist(&model.Good{})

	db := mainDB.Create(&model.Good{
		Name:  name,
		Price: price,
		Type:  goodType,
		PictureStorePath: pictureStorePath,
	})
	if err := db.Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
