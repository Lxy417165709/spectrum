package dao

import (
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/model"
	"strings"
)

var FavorRecordDao favorRecordDao

type favorRecordDao struct{}

//func (favorRecordDao) UpdateExpenseInfo(obj model.Chargeable, expenseInfo *pb.ExpenseInfo) error {
//	createTableWhenNotExist(obj)
//
//	// todo: 这三个字段名是约定，但这容易出错
//	to := map[string]interface{}{
//		"id":                obj.GetID(),
//		"check_out_at":      expenseInfo.CheckOutAt,
//		"expense":           expenseInfo.Expense,
//		"non_favor_expense": expenseInfo.NonFavorExpense,
//	}
//	// todo: 要确定 where 条件，是否是 id == to[id]
//	if err := mainDB.Table(obj.GetChargeableObjectName()).Update(to).Error; err != nil {
//		logger.Error("Fail to finish mainDB.Update", zap.Any("to", to), zap.Error(err))
//		return err
//	}
//	return nil
//}

// todo: 这个结构可以简化为2个参数
func (favorRecordDao) GetFavorRecords(chargeableObjectName string, chargeableObjectId int64) ([]*model.FavorRecord, error) {
	var result []*model.FavorRecord
	if err := mainDB.Find(
		&result,
		"chargeable_object_name = ? and chargeable_object_id = ?",
		chargeableObjectName, chargeableObjectId,
	).Error; err != nil {
		logger.Error(
			"Fail to finish mainDB.Find",
			zap.Int64("chargeableObjectId", chargeableObjectId),
			zap.String("chargeableObjectName", chargeableObjectName),
			zap.Error(err))
		return nil, ers.MysqlError
	}
	return result, nil
}

func (favorRecordDao) CreateFavorRecord(chargeableObjName string, chargeableObjID int64, favors []*pb.Favor) error {
	for _, favor := range favors {
		obj := &model.FavorRecord{
			ChargeableObjectName: chargeableObjName,
			ChargeableObjectID:   chargeableObjID,
			FavorType:            favor.FavorType,
			FavorParameters:      strings.Join(favor.Parameters, "|"),
		}
		if err := mainDB.Create(obj).Error; err != nil {
			logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
			return ers.MysqlError
		}
	}
	return nil
}

func (favorRecordDao) DeleteFavorRecord(chargeableObjName string, chargeableObjID int64, favor *pb.Favor) error {
	var table model.FavorRecord
	createTableWhenNotExist(&table)

	obj := &model.FavorRecord{
		ChargeableObjectName: chargeableObjName,
		ChargeableObjectID:   chargeableObjID,
		FavorType:            favor.FavorType,
		FavorParameters:      strings.Join(favor.Parameters, "|"),
	}
	if err := mainDB.Limit(1).Where(
		"chargeable_object_name = ? and chargeable_object_id = ? and favor_type = ? and favor_parameters = ?",
		chargeableObjName, chargeableObjID, favor.FavorType, strings.Join(favor.Parameters, "|"),
	).Delete(&table).Error; err != nil {
		logger.Error("Fail to finish mainDB.Delete", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}

func (favorRecordDao) BatchDeleteFavorRecord(chargeableObjName string, chargeableObjIDs []int64) error {
	var table model.FavorRecord
	createTableWhenNotExist(&table)
	if err := mainDB.Where(
		"chargeable_object_name = ? and chargeable_object_id in (?)",
		chargeableObjName, chargeableObjIDs,
	).Delete(&table).Error; err != nil {
		logger.Error("Fail to finish mainDB.Delete",
			zap.String("chargeableObjName", chargeableObjName),
			zap.Any("chargeableObjIDs", chargeableObjIDs),
			zap.Error(err))
		return err
	}
	return nil
}
