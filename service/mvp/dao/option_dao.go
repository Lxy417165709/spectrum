package dao

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var OptionDao optionDao

type optionDao struct{}

func (optionDao) Get(id int) (*model.Option, error) {
	obj, err := universalGet(id, &model.Option{})
	if err != nil {
		logger.Error("Fail to finish universalGet",
			zap.Any("id", id),
			zap.Error(err))
		return nil, err
	}
	if obj == nil {
		return nil, nil
	}
	return obj.(*model.Option), nil
}

func (optionDao) Create(optionClassID int, optionName string) error {
	createTableWhenNotExist(&model.Option{})
	db := mainDB.Create(&model.Option{
		OptionClassID: optionClassID,
		Name:          optionName,
	})
	if err := db.Error; err != nil {
		logger.Error("Fail to create option",
			zap.Any("optionClassID", optionClassID),
			zap.Any("optionName", optionName),
			zap.Error(err))
		return err
	}
	return nil
}

func (optionDao) GetByOptionClassID(optionClassID int) ([]*model.Option, error) {
	createTableWhenNotExist(&model.Option{})
	var options []*model.Option
	db := mainDB.Find(&options, "option_class_id = ?", optionClassID)
	if err := db.Error; err != nil {
		logger.Error("Fail to get options",
			zap.Any("optionClassID", optionClassID),
			zap.Error(err))
		return nil, err
	}
	return options, nil
}

func (optionDao) Del(optionClassID int, optionName string) error {
	createTableWhenNotExist(&model.Option{})
	db := mainDB.Delete(&model.Option{}, "name = ? and option_class_id = ?", optionName, optionClassID)
	if err := db.Error; err != nil {
		logger.Error("Fail to del option",
			zap.Any("optionClassID", optionClassID),
			zap.Any("optionName", optionName),
			zap.Error(err))
		return err
	}
	return nil
}

func (optionDao) GetByName(name string) (*model.Option, error) {
	var obj model.Option
	createTableWhenNotExist(&obj)
	db := mainDB.First(&obj, "name = ?", name)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.First",
			zap.Any("name", name),
			zap.Error(err))
		return nil, err
	}
	return &obj, nil
}