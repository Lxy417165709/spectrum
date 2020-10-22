package dao

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/logger"
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

func (goodDao) Create(name string, price float64,pictureStorePath string,classID int) error {
	createTableWhenNotExist(&model.Good{})

	db := mainDB.Create(&model.Good{
		Name:             name,
		Price:            price,
		PictureStorePath: pictureStorePath,
		ClassID:  classID,
	})
	if err := db.Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (goodDao) GetAll() ([]*model.Good, error) {
	createTableWhenNotExist(&model.Good{})

	var goods []*model.Good
	db := mainDB.Find(&goods)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logs.Error(err)
		return nil, err
	}
	return goods, nil
}

func (goodDao) GetByGoodClassID(goodClassID int) ([]*model.Good, error) {
	createTableWhenNotExist(&model.Good{})

	var goods []*model.Good
	db := mainDB.Find(&goods, "class_id = ?", goodClassID)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logs.Error(err)
		return nil, err
	}
	return goods, nil
}

func (goodDao) UpdateGoodClassID(name string, goodClassID int) error {
	createTableWhenNotExist(&model.Good{})

	db := mainDB.Table((&model.Good{}).TableName()).Where("name = ?", name).Update(&model.Good{
		ClassID: goodClassID,
	})
	if err := db.Error; err != nil {
		logger.Error("Fail to update good",
			zap.String("goodName", name),
			zap.Int("goodClassID", goodClassID),
			zap.Error(err))
		return err
	}
	return nil
}
