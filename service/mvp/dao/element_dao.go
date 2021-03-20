package dao

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/model"
)

var ElementDao elementDao

type elementDao struct{}

func (elementDao) Create(obj *model.Element) (int64, error) {
	values := []interface{}{
		obj.ID, obj.ClassID, obj.Name, obj.Type,
	}
	sql := fmt.Sprintf(`
		insert into %s(id,class_id,name, type) values(%s)
		on duplicate key update
			class_id = values(class_id),
			name = values(name),
			type = values(type);
	`, obj.TableName(), GetPlaceholderClause(len(values)))
	result, err := mainDB.CommonDB().Exec(sql, values...)
	if err != nil {
		logger.Error("Fail to finish create", zap.Any("obj", obj), zap.Error(err))
		return 0, ers.New("操作失败，可能存在同名的元素(%s)。", obj.Name)
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Fail to get id", zap.Any("obj", obj), zap.Error(err))
		return 0, ers.MysqlError
	}
	return id, nil
}

//func (elementDao) Del(name string, size string) error {
//	obj := &model.Element{}
//	values := []interface{}{
//		name, size,
//	}
//	sql := fmt.Sprintf(`delete from %s where name = ? and size = ?;`, obj.TableName())
//	if err := mainDB.Exec(sql, values...).Error; err != nil {
//		logger.Error("Fail to finish delete", zap.Any("name", name), zap.Any("size", size), zap.Error(err))
//		return ers.MysqlError
//	}
//	return nil
//}

func (elementDao) Get(id int64) (*model.Element, error) {
	var result model.Element
	if err := mainDB.First(&result, "id = ?", id).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.Any("id", id), zap.Error(err))
		return nil, err
	}
	return &result, nil
}

func (elementDao) GetOne(name string, classID int64) (*model.Element, error) {
	var result model.Element
	if err := mainDB.First(&result, "name = ? and class_id = ?", name, classID).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.String("name", name), zap.Any("classID", classID), zap.Error(err))
		return nil, ers.MysqlError
	}
	return &result, nil
}

func (elementDao) GetByClassID(classID int64) ([]*model.Element, error) {
	var result []*model.Element
	if err := mainDB.Find(&result, "class_id = ?", classID).Error; err != nil {
		logger.Error("Fail to finish mainDB.Find", zap.Any("classID", classID), zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (elementDao) GetAllAttachElements(classID int64) ([]*model.Element, error) {
	var whereClause string
	var whereValues []interface{}

	whereClause = " type != ? "
	whereValues = append(whereValues, pb.ElementType_Main)

	if classID != 0 {
		whereClause += " and class_id = ? "
		whereValues = append(whereValues, classID)
	}

	var elements []*model.Element
	db := mainDB.Where(whereClause, whereValues...).Find(&elements)
	if err := db.Error; err != nil {
		logs.Error("Fail to finish mainDB.Find", zap.Error(err))
		return nil, err
	}
	return elements, nil
}
