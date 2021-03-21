package dao

import (
	"fmt"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var ElementSizeInfoDao elementSizeInfoDao

type elementSizeInfoDao struct{}

func (elementSizeInfoDao) Get(elementID int64) ([]*model.ElementSizeInfo, error) {
	var result []*model.ElementSizeInfo
	if err := mainDB.Where("element_id = ?",  elementID).Find(&result).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.Any("elementID", elementID), zap.Error(err))
		return nil, ers.MysqlError
	}
	return result, nil
}

func (elementSizeInfoDao) Create(obj *model.ElementSizeInfo) (int64, error) {
	values := []interface{}{
		obj.ID, obj.ElementID, obj.Size, obj.PictureStorePath, obj.Price,
	}

	sql := fmt.Sprintf(`
		insert into %s(id,element_id,size,picture_store_path,price) values(%s)
		on duplicate key update
			element_id = values(element_id),
			size = values(size),
			picture_store_path = values(picture_store_path),
			price = values(price);
	`, fmt.Sprintf("`%s`", obj.TableName()), GetPlaceholderClause(len(values)))
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
