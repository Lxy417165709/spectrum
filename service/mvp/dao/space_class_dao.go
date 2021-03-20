package dao

import (
	"fmt"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var SpaceClassDao spaceClassDao

type spaceClassDao struct{}

func (spaceClassDao) Create(obj *model.SpaceClass) (int64, error) {
	values := []interface{}{
		obj.ID, obj.Name, obj.PictureStorePath,
	}
	sql := fmt.Sprintf(`
		insert into %s(id,name,picture_store_path) values(%s)
		on duplicate key update
			name = values(name),
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

func (spaceClassDao) Get(id int64) (*model.SpaceClass, error) {
	var result model.SpaceClass
	if err := mainDB.First(&result, "id = ?", id).Error; err != nil {
		logger.Error("Fail to finish mainDB.First", zap.Error(err))
		return nil, ers.MysqlError
	}
	return &result, nil
}
func (spaceClassDao) GetByName(name string) (*model.SpaceClass, error) {
	var result model.SpaceClass
	if err := mainDB.First(&result, "name = ?", name).Error; err != nil {
		logger.Error("Fail to finish mainDB.First", zap.Error(err))
		return nil, ers.MysqlError
	}
	return &result, nil
}

func (spaceClassDao) GetAllClasses() ([]*model.SpaceClass, error) {
	var result []*model.SpaceClass
	if err := mainDB.Find(&result).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.Error(err))
		return nil, ers.MysqlError
	}
	return result, nil
}
