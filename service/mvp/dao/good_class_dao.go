package dao

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var GoodClassDao goodClassDao

type goodClassDao struct{}

func (goodClassDao) Get(id int) (*model.GoodClass, error) {
	obj, err := universalGet(id, &model.GoodClass{})
	if err != nil {
		logger.Error("Fail to finish universalGet",
			zap.Any("id", id),
			zap.Error(err))
		return nil, err
	}
	return obj.(*model.GoodClass), nil
}

func (goodClassDao) GetByName(goodClassName string) (*model.GoodClass, error) {
	createTableWhenNotExist(&model.GoodClass{})
	var goodClass model.GoodClass
	db := mainDB.First(&goodClass, "name = ?", goodClassName)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to get good class",
			zap.Any("goodClassName", goodClassName),
			zap.Error(err))
		return nil, err
	}
	return &goodClass, nil
}

func (goodClassDao) GetAll() ([]*model.GoodClass, error) {
	createTableWhenNotExist(&model.GoodClass{})

	var goodClasses []*model.GoodClass
	db := mainDB.Find(&goodClasses)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to get all good classes",
			zap.Error(err))
		return nil, err
	}
	return goodClasses, nil
}

func (goodClassDao) Create(className string, classType int) error {
	createTableWhenNotExist(&model.GoodClass{})
	db := mainDB.Create(&model.GoodClass{
		Name:      className,
		ClassType: classType,
	})
	if err := db.Error; err != nil {
		logger.Error("Fail to create good class",
			zap.Any("className", className),
			zap.Any("classType", classType),
			zap.Error(err))
		return err
	}
	return nil
}

func (goodClassDao) DeleteByNames(classNames []string) error {
	createTableWhenNotExist(&model.GoodClass{})

	db := mainDB.Delete(&model.GoodClass{}, "name in (?)", classNames)
	if err := db.Error; err != nil {
		logger.Error("Fail to delete good class",
			zap.Any("classNames", classNames),
			zap.Error(err))
		return err
	}
	return nil
}

func (goodClassDao) GetByIDs(ids []int) ([]*model.GoodClass, error) {
	createTableWhenNotExist(&model.GoodClass{})
	var classes []*model.GoodClass
	db := mainDB.Find(&classes, "id in (?)", ids)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to get good class",
			zap.Any("ids", ids),
			zap.Error(err))
		return nil, err
	}
	return classes, nil
}
