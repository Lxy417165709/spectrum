package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var MainElementSizeRecordDao mainElementSizeRecordDao

type mainElementSizeRecordDao struct{}

func (mainElementSizeRecordDao) Create(obj *model.MainElementSizeRecord) (int64, error) {
	values := []interface{}{
		obj.GoodID, obj.MainElementName, obj.SelectSize,
	}
	sql := fmt.Sprintf(`
		insert into %s(good_id,main_element_name,select_size) values(%s)
		on duplicate key update
			good_id = values(good_id),
			main_element_name = values(main_element_name),
			select_size = values(select_size);
	`, obj.TableName(), GetPlaceholderClause(len(values)))
	result, err := mainDB.CommonDB().Exec(sql, values...)
	if err != nil {
		logger.Error("Fail to finish create", zap.Any("obj", obj), zap.Error(err))
		return 0, ers.New("数据库执行出错，添加主元素(%v)默认选项(%v)失败。", obj.MainElementName, obj.SelectSize)
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Fail to get id", zap.Any("obj", obj), zap.Error(err))
		return 0, ers.MysqlError
	}
	return id, nil
}

func (mainElementSizeRecordDao) GetByGoodIdAndMainElementName(goodID int64, mainElementName string) (*model.MainElementSizeRecord, error) {
	var table model.MainElementSizeRecord
	createTableWhenNotExist(&table)

	var result model.MainElementSizeRecord
	var whereClause string
	var parameters []interface{}
	if goodID == 0 {
		whereClause = "good_id = ? and main_element_name = ?"
		parameters = []interface{}{goodID, mainElementName}
	} else {
		whereClause = "good_id = ?"
		parameters = []interface{}{goodID}
	}

	if err := mainDB.Where(whereClause, parameters...).First(&result).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.Find", zap.String("mainElementName", mainElementName), zap.Error(err))
		return nil, err
	}
	return &result, nil
}

func (mainElementSizeRecordDao) BatchDelete(goodIDs []int64) error {
	var table model.MainElementSizeRecord
	createTableWhenNotExist(&table)

	// todo: 这里要测试
	if err := mainDB.Where("good_id in (?)", goodIDs).Delete(&table).Error; err != nil {
		logger.Error("Fail to finish mainDB.Delete", zap.Any("goodIDs", goodIDs), zap.Error(err))
		return err
	}
	return nil
}
