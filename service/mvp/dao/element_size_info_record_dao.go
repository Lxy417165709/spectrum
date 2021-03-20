package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var ElementSizeInfoRecordDao elementSizeInfoRecordDao

type elementSizeInfoRecordDao struct{}

func (elementSizeInfoRecordDao) Get(goodID int64, elementID int64) ([]*model.ElementSizeInfoRecord, error) {
	var result []*model.ElementSizeInfoRecord
	if err := mainDB.Where("good_id = ? and element_id = ?", goodID, elementID).Find(&result).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.first", zap.Any("goodID", goodID), zap.Any("elementID", elementID), zap.Error(err))
		return nil, ers.MysqlError
	}
	return result, nil
}

func (elementSizeInfoRecordDao) Create(obj *model.ElementSizeInfoRecord) (int64, error) {
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
