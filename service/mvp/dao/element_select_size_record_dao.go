package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var ElementSelectSizeRecordDao elementSelectSizeRecordDao

type elementSelectSizeRecordDao struct{}

func (elementSelectSizeRecordDao) Create(obj *model.ElementSelectSizeRecord) (int64, error) {
	values := []interface{}{
		obj.GoodID, obj.ElementID, obj.SelectSizeInfoID,
	}
	sql := fmt.Sprintf(`
		insert into %s(good_id,element_id,select_size_info_id) values(%s)
		on duplicate key update
			good_id = values(good_id),
			element_id = values(element_id),
			select_size_info_id = values(select_size_info_id);
	`, fmt.Sprintf("`%s`", obj.TableName()), getPlaceholderClause(len(values)))
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

func (elementSelectSizeRecordDao) GetOne(goodID, elementID int64) (*model.ElementSelectSizeRecord, error) {
	var result model.ElementSelectSizeRecord
	if err := mainDB.Where("good_id = ? and element_id = ?", goodID, elementID).First(&result).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.first", zap.Any("goodID", goodID), zap.Any("elementID", elementID), zap.Error(err))
		return nil, ers.MysqlError
	}
	return &result, nil
}

func (elementSelectSizeRecordDao) BatchDelete(goodIDs []int64) error {
	var table model.ElementSelectSizeRecord
	createTableWhenNotExist(&table)

	// todo: 这里要测试
	if err := mainDB.Where("good_id in (?)", goodIDs).Delete(&table).Error; err != nil {
		logger.Error("Fail to finish mainDB.Delete", zap.Any("goodIDs", goodIDs), zap.Error(err))
		return err
	}
	return nil
}
