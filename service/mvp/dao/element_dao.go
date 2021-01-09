package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var ElementDao elementDao

type elementDao struct{}

func (elementDao) Create(obj *model.Element) error {
	var table model.Element
	createTableWhenNotExist(&table)

	if err := mainDB.Create(obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}

//func (elementDao) Get(id int) (*model.Element, error) {
//	createTableWhenNotExist(&table)
//
//	var element model.Element
//	db := mainDB.First(&element, "id = ?", id)
//	if err := db.Error; err != nil {
//		if gorm.IsRecordNotFoundError(err) {
//			return nil, nil
//		}
//		logger.Error("Fail to finish mainDB.First",zap.Int("id",id),zap.Error(err))
//		return nil, err
//	}
//	return &element, nil
//}

func (elementDao) GetByName(name string) ([]*model.Element, error) {
	var table model.Element
	createTableWhenNotExist(&table)

	var result []*model.Element
	if err := mainDB.Find(&result, "name = ?", name).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.String("name", name), zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (elementDao) GetByClassName(className string) ([]*model.Element, error) {
	var table model.Element
	createTableWhenNotExist(&table)

	var result []*model.Element
	if err := mainDB.Find(&result, "class_name = ?", className).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.String("className", className), zap.Error(err))
		return nil, err
	}
	return result, nil
}

//func (elementDao) GetAll() ([]*model.Element, error) {
//	createTableWhenNotExist(&model.Element{})
//
//	var elements []*model.element
//	db := mainDB.Find(&elements)
//	if err := db.Error; err != nil {
//		if gorm.IsRecordNotFoundError(err) {
//			return nil, nil
//		}
//		logs.Error(err)
//		return nil, err
//	}
//	return elements, nil
//}

//func (elementDao) GetByelementClassID(elementClassID int) ([]*model.element, error) {
//	createTableWhenNotExist(&model.element{})
//
//	var elements []*model.element
//	db := mainDB.Find(&elements, "class_id = ?", elementClassID)
//	if err := db.Error; err != nil {
//		if gorm.IsRecordNotFoundError(err) {
//			return nil, nil
//		}
//		logs.Error(err)
//		return nil, err
//	}
//	return elements, nil
//}

//func (elementDao) UpdateelementClassID(name string, elementClassID int) error {
//	createTableWhenNotExist(&model.element{})
//
//	db := mainDB.Table((&model.element{}).TableName()).Where("name = ?", name).Update(&model.element{
//		ClassID: elementClassID,
//	})
//	if err := db.Error; err != nil {
//		logger.Error("Fail to update element",
//			zap.String("elementName", name),
//			zap.Int("elementClassID", elementClassID),
//			zap.Error(err))
//		return err
//	}
//	return nil
//}
