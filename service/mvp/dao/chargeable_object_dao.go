package dao

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/model"
	"strings"
)

// todo: 这个设计的不太好
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

// 以下只用到了 chargeableObj  的 GetID, GetName
func (chargeableObjectDao) UpdateExpenseInfo(obj model.Chargeable, expenseInfo *pb.ExpenseInfo) error {
	createTableWhenNotExist(obj)

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

func (chargeableObjectDao) CreateCheckOutRecord(obj *model.CheckOutRecord) error {
	var table model.CheckOutRecord
	createTableWhenNotExist(&table)

	if err := mainDB.Create(obj).Error; err != nil {
		logger.Error("Fail to finish mainDB.Create", zap.Any("obj", obj), zap.Error(err))
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

func (chargeableObjectDao) CreateFavorRecord(cObj pb.Chargeable) error {
	var table model.FavorRecord
	createTableWhenNotExist(&table)
	for _, favor := range cObj.GetFavors() {
		obj := &model.FavorRecord{
			ChargeableObjectName: cObj.GetName(),
			ChargeableObjectID:   cObj.GetId(),
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
