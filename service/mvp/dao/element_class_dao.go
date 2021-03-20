package dao

import (
	"fmt"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/model"
)

var ElementClassDao elementClassDao

type elementClassDao struct{}

func (elementClassDao) Create(obj *model.ElementClass) (int64, error) {
	values := []interface{}{
		obj.ID, obj.Name, obj.PictureStorePath, obj.ClassType,
	}
	sql := fmt.Sprintf(`
		insert into %s(id, name, picture_store_path,class_type) values(%s)
		on duplicate key update
			name = values(name),
			picture_store_path = values(picture_store_path);
	`, fmt.Sprintf("`%s`", obj.TableName()), GetPlaceholderClause(len(values)))

	result, err := mainDB.CommonDB().Exec(sql, values...)
	if err != nil {
		logger.Error("Fail to finish create", zap.Any("obj", obj), zap.Error(err))
		return 0, ers.New("操作失败，可能存在同名的商品类(%s)。", obj.Name)
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Fail to get id", zap.Any("obj", obj), zap.Error(err))
		return 0, ers.MysqlError
	}
	return id, nil
}

func (elementClassDao) GetClasses(classType pb.ElementType) ([]*model.ElementClass, error) {
	var result []*model.ElementClass
	if err := mainDB.Find(&result, "class_type = ?", classType).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (elementClassDao) GetByName(name string) (*model.ElementClass, error) {
	var result model.ElementClass
	if err := mainDB.First(&result, "name = ?", name).Error; err != nil {
		logger.Error("Fail to finish mainDB.First", zap.Error(err))
		return nil, ers.MysqlError
	}
	return &result, nil
}

func (elementClassDao) Get(id int64) (*model.ElementClass, error) {
	var result model.ElementClass
	if err := mainDB.First(&result, "id = ?", id).Error; err != nil {
		logger.Error("Fail to finish mainDB.First", zap.Error(err))
		return nil, ers.MysqlError
	}
	return &result, nil
}
