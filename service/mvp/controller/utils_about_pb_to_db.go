package controller

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
	"strings"
	"time"
)

func writeFavorToDB(favorableStruct interface{}) error {
	switch favorableStruct.(type) {
	case *pb.Good:
		good := favorableStruct.(*pb.Good)
		for _, favor := range good.Favors {
			if err := dao.FavorRecordDao.Create(&model.FavorRecord{
				ChargeableObjectName: model.ChargeableObjectNameOfGood,
				ChargeableObjectID:   good.Id,
				FavorType:            favor.FavorType,
				FavorParameters:      strings.Join(favor.Parameters, "|"),
			}); err != nil {
				// todo: log
				return err
			}
		}
	case *pb.Desk:
		desk := favorableStruct.(*pb.Desk)
		for _, favor := range desk.Favors {
			if err := dao.FavorRecordDao.Create(&model.FavorRecord{
				ChargeableObjectName: model.ChargeableObjectNameOfDesk,
				ChargeableObjectID:   desk.Id,
				FavorType:            favor.FavorType,
				FavorParameters:      strings.Join(favor.Parameters, "|"),
			}); err != nil {
				// todo: log
				return err
			}
		}
	case *pb.Order:
		order := favorableStruct.(*pb.Order)
		for _, favor := range order.Favors {
			if err := dao.FavorRecordDao.Create(&model.FavorRecord{
				ChargeableObjectName: model.ChargeableObjectNameOfOrder,
				ChargeableObjectID:   order.Id,
				FavorType:            favor.FavorType,
				FavorParameters:      strings.Join(favor.Parameters, "|"),
			}); err != nil {
				// todo: log
				return err
			}
		}
	}
	return nil
}

func checkOut(chargeableDao dao.ChargeableDao, ids []int64) error {
	for _, id := range ids {
		checkOutTimestamp := time.Now().Unix()
		if err := dao.CheckOutRecordDao.Create(&model.CheckOutRecord{
			ChargeableObjectName: chargeableDao.GetChargeableObjectName(),
			ChargeableObjectID:   id,
			CheckOutTimestamp:    checkOutTimestamp,
		}); err != nil {
			logger.Error("Fail to finish CheckOutRecordDao.Create", zap.Error(err))
			return err
		}
		if err := chargeableDao.Update(map[string]interface{}{
			"id":                  id,
			"check_out_timestamp": checkOutTimestamp,
		}); err != nil {
			logger.Error("Fail to finish chargeableDao.Update", zap.Error(err))
			return err
		}
	}
	return nil
}

func writeGoodSizeToDB(good *pb.Good) error {
	if len(good.MainElement.SizeInfos) == 0 {
		return nil
	}

	// 1. 创建主元素、主元素尺寸的对应关系
	if err := dao.MainElementSizeRecordDao.Create(&model.MainElementSizeRecord{
		GoodID:          good.Id,
		MainElementName: good.MainElement.Name,
		SelectSize:      good.MainElement.SizeInfos[0].Size,
	}); err != nil {
		logger.Error("Fail to finish MainElementSizeRecordDao.Create", zap.Error(err))
		return err
	}

	// 2. 创建主元素、附属元素、附属元素尺寸的对应关系
	for _, attachElement := range good.AttachElements {
		if err := dao.MainElementAttachElementRecordDao.Create(&model.MainElementAttachElementRecord{
			GoodID:            good.Id,
			MainElementName:   good.MainElement.Name,
			AttachElementName: attachElement.Name,
			SelectSize:        getSelectSizeInfo(attachElement.SizeInfos).Size,
		}); err != nil {
			logger.Error("Fail to finish MainElementAttachElementDao.Create", zap.Error(err))
			return err
		}
	}
	return nil
}

func closeDeskIfOpening(deskID int64, endTimestamp int64) error {
	desk, err := dao.DeskDao.Get(deskID)
	if err != nil {
		logger.Error("Fail to finish DeskDao.Get",
			zap.Error(err))
		return err
	}
	if !desk.IsOpening() {
		logger.Warn("Desk had been closed", zap.Int64("deskID", deskID))
		return nil
	}
	if err := dao.DeskDao.Update(map[string]interface{}{
		"id":            deskID,
		"end_timestamp": endTimestamp,
	}); err != nil {
		logger.Error("Fail to finish DeskDao.Update",
			zap.Error(err))
		return err
	}
	return nil
}

func createElement(pbElement *pb.Element, className string) error {
	dbElements := getDbElements(pbElement, className)
	for _, dbElement := range dbElements {
		if err := dao.ElementDao.Create(dbElement); err != nil {
			logger.Error("Fail to finish ElementDao.Create", zap.Error(err))
			return err
		}
	}
	return nil
}
