package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var SpaceDao spaceDao

type spaceDao struct{}

func (spaceDao) Create(obj *model.Space) (int64, error) {
	values := []interface{}{
		obj.ID, obj.ClassID, obj.Name, obj.BillingType, obj.Price, obj.PictureStorePath,
	}
	sql := fmt.Sprintf(`
		insert into %s(id,class_id,name, billing_type,price,picture_store_path) values(%s)
		on duplicate key update
			class_id = values(class_id),
			name = values(name),
			billing_type = values(billing_type),
			price = values(price),
			picture_store_path = values(picture_store_path);
	`, obj.TableName(), GetPlaceholderClause(len(values)))
	result, err := mainDB.CommonDB().Exec(sql, values...)
	if err != nil {
		logger.Error("Fail to finish create", zap.Any("obj", obj), zap.Error(err))
		return 0, ers.MysqlError
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Fail to get id", zap.Any("obj", obj), zap.Error(err))
		return 0, ers.MysqlError
	}
	return id, nil
}

func (spaceDao) GetAll() ([]*model.Space, error) {
	var result []*model.Space
	if err := mainDB.Find(&result).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find",
			zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (spaceDao) GetByClassName(classID int64) ([]*model.Space, error) {
	var result []*model.Space
	if err := mainDB.Where("class_id = ?", classID).Find(&result).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find",
			zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (spaceDao) Get(id int64) (*model.Space, error) {
	var result model.Space
	if err := mainDB.First(&result, "id = ?", id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.First",
			zap.Any("id", "id"),
			zap.Error(err))
		return nil, ers.MysqlError
	}
	return &result, nil
}
