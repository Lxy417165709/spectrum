package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/model"
	"strings"
)

var ChargeableObjectDao chargeableObjectDao

type chargeableObjectDao struct{}

func (chargeableObjectDao) Create(obj model.Chargeable) error {
	createTableWhenNotExist(obj)

	if err := mainDB.Create(&obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
		return err
	}
	return nil
}

func (chargeableObjectDao) UpdateExpenseInfo(obj model.Chargeable, expenseInfo *pb.ExpenseInfo) error {
	createTableWhenNotExist(obj)

	// todo: 这三个字段名是约定，但这容易出错
	to := map[string]interface{}{
		"id":                  obj.GetID(),
		"check_out_timestamp": expenseInfo.CheckOutTimestamp,
		"expense":             expenseInfo.Expense,
		"non_favor_expense":   expenseInfo.NonFavorExpense,
	}
	// todo: 要确定 where 条件，是否是 id == to[id]
	if err := mainDB.Table(obj.GetName()).Update(to).Error; err != nil {
		logger.Error("Fail to finish mainDB.Update", zap.Any("to", to), zap.Error(err))
		return err
	}
	return nil
}

func (chargeableObjectDao) GetFavorRecords(obj model.Chargeable) ([]*model.FavorRecord, error) {
	var table model.FavorRecord
	createTableWhenNotExist(&table)

	var result []*model.FavorRecord
	if err := mainDB.Find(
		&result,
		"favorable_struct_name = ? and favorable_struct_id = ？",
		obj.GetName(), obj.GetID(),
	).Error; err != nil {
		logger.Error(
			"Fail to finish mainDB.Find",
			zap.Int64("favorableStructID", obj.GetID()),
			zap.String("favorableStructName", obj.GetName()),
			zap.Error(err))
		return nil, err
	}
	return result, nil
}

func (chargeableObjectDao) CreateFavorRecord(chargeableObjName string, chargeableObjID int64, favors []*pb.Favor) error {
	var table model.FavorRecord
	createTableWhenNotExist(&table)

	for _, favor := range favors {
		obj := &model.FavorRecord{
			ChargeableObjectName: chargeableObjName,
			ChargeableObjectID:   chargeableObjID,
			FavorType:            favor.FavorType,
			FavorParameters:      strings.Join(favor.Parameters, "|"),
		}
		if err := mainDB.Create(obj).Error; err != nil {
			logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
			return err
		}
	}
	return nil
}

func (chargeableObjectDao) DeleteFavorRecord(chargeableObjName string, chargeableObjID int64, favor *pb.Favor) error {
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

func (chargeableObjectDao) BatchDeleteFavorRecord(chargeableObjName string, chargeableObjIDs []int64) error {
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
