package dao

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/model"
)

var ElementDao elementDao

type elementDao struct{}

func (elementDao) Create(obj *model.Element) error {
	values := []interface{}{
		obj.Name, obj.Type, obj.ClassName, obj.Size, obj.Price, obj.PictureStorePath,
	}
	sql := fmt.Sprintf(`
		insert into %s(name, type, class_name, size, price, picture_store_path) values(%s)
		on duplicate key update
			name = values(name),
			type = values(type),
			class_name = values(class_name),
			size = values(size),
			price = values(price),
			picture_store_path = values(picture_store_path);
	`, obj.TableName(), GetPlaceholderClause(len(values)))
	if err := mainDB.Raw(sql, values...).Error; err != nil {
		logger.Error("Fail to finish create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}

func (elementDao) GetByName(name string) ([]*model.Element, error) {
	var result []*model.Element
	if err := mainDB.Find(&result, "name = ?", name).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.String("name", name), zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (elementDao) GetByClassName(className string) ([]*model.Element, error) {
	var result []*model.Element
	if err := mainDB.Find(&result, "class_name = ?", className).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.String("className", className), zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (elementDao) GetAllAttachElements(className string) ([]*model.Element, error) {
	var whereClause string
	var whereValues []interface{}

	whereClause = " type != ? "
	whereValues = append(whereValues, pb.ElementType_Main)

	if className != "" {
		whereClause += " and class_name = ? "
		whereValues = append(whereValues, className)
	}

	var elements []*model.Element
	db := mainDB.Where(whereClause, whereValues...).Find(&elements)
	if err := db.Error; err != nil {
		logs.Error("Fail to finish mainDB.Find", zap.Error(err))
		return nil, err
	}
	return elements, nil
}
