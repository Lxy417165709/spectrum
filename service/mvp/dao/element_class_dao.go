package dao

import (
	"fmt"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var ElementClassDao elementClassDao

type elementClassDao struct{}

// 	ID               uint      `gorm:"id"`
//	CreatedAt        time.Time `gorm:"created_at"`
//	UpdatedAt        time.Time `gorm:"updated_at"`
//	Name             string    `gorm:"name"`
//	PictureStorePath string    `gorm:"picture_store_path"`
func (elementClassDao) Create(obj *model.ElementClass) error {
	values := []interface{}{
		obj.Name, obj.PictureStorePath,
	}
	sql := fmt.Sprintf(`
		insert into %s(name, picture_store_path) values(%s)
		on duplicate key update
			name = values(name),
			picture_store_path = values(picture_store_path);
	`, obj.TableName(), GetPlaceholderClause(len(values)))
	if err := mainDB.Raw(sql, values...).Error; err != nil {
		logger.Error("Fail to finish create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}

func (elementClassDao) GetAllClasses() ([]*model.ElementClass, error) {
	var table model.ElementClass
	createTableWhenNotExist(&table)
	var result []*model.ElementClass
	if err := mainDB.Find(&result).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.Error(err))
		return nil, err
	}
	return result, nil
}
