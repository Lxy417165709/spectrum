package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/service/mvp/model"
)

var ElementSizeRecordDao elementSizeRecordDao

type elementSizeRecordDao struct{}

func (elementSizeRecordDao) Create(obj *model.ElementSizeRecord) (int64, error) {
	values := []interface{}{
		obj.GoodID, obj.ElementClassName, obj.ElementName, obj.SelectSize,
	}
	sql := fmt.Sprintf(`
		insert into %s(good_id,element_class_name,element_name,select_size) values(%s)
		on duplicate key update
			good_id = values(good_id),
			element_name = values(element_name),
			element_class_name = values(element_class_name),
			select_size = values(select_size);
	`, obj.TableName(), GetPlaceholderClause(len(values)))
	result, err := mainDB.CommonDB().Exec(sql, values...)
	if err != nil {
		logger.Error("Fail to finish create", zap.Any("obj", obj), zap.Error(err))
		return 0, ers.New("数据库执行出错，添加主元素(%v)默认选项(%v)失败。", obj.ElementName, obj.SelectSize)
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Fail to get id", zap.Any("obj", obj), zap.Error(err))
		return 0, ers.MysqlError
	}
	return id, nil
}

func (elementSizeRecordDao) GetByGoodIdAndElementName(goodID int64, elementName, className string) (*model.ElementSizeRecord, error) {
	var table model.ElementSizeRecord
	createTableWhenNotExist(&table)

	var result model.ElementSizeRecord
	var whereClause string
	var parameters []interface{}
	if goodID == 0 {
		whereClause = "good_id = ? and element_name = ? and element_class_name = ?"
		parameters = []interface{}{goodID, elementName, className}
	} else {
		whereClause = "good_id = ? and element_class_name = ?"
		parameters = []interface{}{goodID, className}
	}

	if err := mainDB.Where(whereClause, parameters...).First(&result).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		logger.Error("Fail to finish mainDB.first", zap.String("elementName", elementName), zap.Error(err))
		return nil, ers.MysqlError
	}
	return &result, nil
}

func (elementSizeRecordDao) BatchDelete(goodIDs []int64) error {
	var table model.ElementSizeRecord
	createTableWhenNotExist(&table)

	// todo: 这里要测试
	if err := mainDB.Where("good_id in (?)", goodIDs).Delete(&table).Error; err != nil {
		logger.Error("Fail to finish mainDB.Delete", zap.Any("goodIDs", goodIDs), zap.Error(err))
		return err
	}
	return nil
}
