package controller

import (
	"go.uber.org/zap"
	"spectrum/common/ers"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
	"spectrum/service/mvp/utils"
)

func writePbGoodToDB(good *pb.Good, orderID int64) error {
	// 1. 创建商品记录，获取商品ID
	dbGood := &model.Good{
		ID:              good.Id,
		OrderID:         orderID,
		MainElementID:   good.MainElement.Id,
		Expense:         good.ExpenseInfo.Expense,
		CheckOutAt:      utils.ToTime(good.ExpenseInfo.CheckOutAt),
		NonFavorExpense: good.ExpenseInfo.NonFavorExpense,
	}
	goodID, errResult := dao.GoodDao.Create(dbGood)
	if errResult != nil {
		return errResult
	}

	// 2. 添加/更新商品尺寸选择、商品主元素附属元素关联关系
	if errResult := writePbGoodSelectedSizeInfoIndexRecordAndMainAttachElementRecordToDB(
		good.MainElement, good.AttachElements, goodID); errResult != nil {
		return errResult
	}

	// 3. 记录商品优惠记录
	if errResult := dao.FavorRecordDao.CreateFavorRecords(dbGood.GetChargeableObjectName(), goodID, good.Favors); errResult != nil {
		return errResult
	}

	// 4. 返回
	return nil
}

func writePbElementSelectSizeRecord(pbElement *pb.Element, goodID int64) error {
	elementSelectedSizeInfoID := getPbElementSelectSizeInfoID(pbElement)
	if _, errResult := dao.ElementSelectSizeRecordDao.Create(toDbElementSelectSizeRecord(goodID, pbElement.Id, elementSelectedSizeInfoID));
		errResult != nil {
		return errResult
	}
	return nil
}

// 将元素的元信息记录到数据库，元信息包括: 元素信息、元素拥有的尺寸信息
func writePbElementMetaObjectToDbAndUpdateID(pbElement *pb.Element, classID int64) error {
	// 1. 将元素写入数据库，如果先前该元素不存在，则更新元素ID
	elementID, errResult := dao.ElementDao.Create(toDbElement(pbElement, classID))
	if errResult != nil {
		return errResult
	}
	if pbElement.Id == 0 {
		pbElement.Id = elementID
	}

	// 2. 将元素的尺寸信息写入数据库，如果先前尺寸不存在，则更新元素的尺寸ID
	for index, pbSizeInfo := range pbElement.SizeInfos {
		sizeInfoID, errResult := dao.ElementSizeInfoDao.Create(toDbElementSizeInfo(pbSizeInfo, pbElement.Id))
		if errResult != nil {
			return ers.New("存储第 %d 个尺寸时出错。%s", index+1, errResult.Error())
		}
		if pbSizeInfo.Id == 0 {
			pbSizeInfo.Id = sizeInfoID
		}
	}

	// 3. 返回
	return nil
}

// 将商品的主元素尺寸选择信息、附属元素尺寸选择信息、主元素与附属元素的对应记录写入数据库
func writePbGoodSelectedSizeInfoIndexRecordAndMainAttachElementRecordToDB(mainElement *pb.Element, attachElements []*pb.Element, goodID int64) error {
	// 1. 创建主元素、主元素尺寸的对应关系
	if errResult := writePbElementSelectSizeRecord(mainElement, goodID); errResult != nil {
		return errResult
	}

	// 2. 创建附属元素、附属元素尺寸的对应关系
	for index, attachElement := range attachElements {
		if errResult := writePbElementSelectSizeRecord(attachElement, goodID); errResult != nil {
			return ers.New("创建第 %d 个附属元素与其尺寸的对应关系时出错。%s", index+1, errResult.Error())
		}
	}

	// 3. 创建主元素、附属元素的对应关系
	for index, attachElement := range attachElements {
		if _, errResult := dao.MainElementAttachElementRecordDao.Create(&model.MainElementAttachElementRecord{
			GoodID:          goodID,
			AttachElementID: attachElement.Id,
			MainElementID:   mainElement.Id,
		}); errResult != nil {
			return ers.New("创建第 %d 个主元素与其附属元素的对应关系时出错。%s", index+1, errResult.Error())
		}
	}

	// 4. 返回
	return nil
}

func getPbElementSelectSizeInfoID(pbElement *pb.Element) int64 {
	return model.GetPbElementSelectSizeInfo(pbElement).Id
}

func closeDeskIfOpening(deskID int64, endTimestamp int64) error {
	desk, err := dao.DeskDao.GetByID(deskID)
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

//func checkOutIfNot(chargeableObj model.Chargeable) error {
//
//	expenseInfo := getExpenseInfo(chargeableObj)
//	if expenseInfo.CheckOutAt != 0 {
//		// todo: 警告
//		// 这里表示商品已结账过了
//		return nil
//	}
//	expenseInfo.CheckOutAt = time.Now().Unix()
//	if err := dao.FavorRecordDao.UpdateExpenseInfo(chargeableObj, expenseInfo); err != nil {
//		logger.Error("Fail to finish chargeableDao.Update", zap.Error(err))
//		return err
//	}
//
//	// 添加结账记录
//	if err := dao.CheckOutRecordDao.Create(&model.CheckOutRecord{
//		ChargeableObjectName: chargeableObj.GetChargeableObjectName(),
//		ChargeableObjectID:   chargeableObj.GetID(),
//		CheckOutAt:           expenseInfo.CheckOutAt,
//	}); err != nil {
//		logger.Error("Fail to finish CheckOutRecordDao.Create", zap.Error(err))
//		return err
//	}
//
//	return nil
//}

func toDbSpace(pbSpace *pb.Space, classID int64) *model.Space {
	return &model.Space{
		ID:               pbSpace.Id,
		Name:             pbSpace.Name,
		ClassID:          classID,
		Price:            utils.GetDbPrice(pbSpace.Price),
		BillingType:      pbSpace.BillingType,
		PictureStorePath: pbSpace.PictureStorePath,
	}
}

func toDbElement(pbElement *pb.Element, classID int64) *model.Element {
	return &model.Element{
		ID:      pbElement.Id,
		Name:    pbElement.Name,
		Type:    pbElement.Type,
		ClassID: classID,
	}
}

func toDbElementSizeInfo(pbSizeInfo *pb.SizeInfo, elementID int64) *model.ElementSizeInfo {
	return &model.ElementSizeInfo{
		ID:               pbSizeInfo.Id,
		ElementID:        elementID,
		Size:             pbSizeInfo.Size,
		Price:            utils.GetDbPrice(pbSizeInfo.Price),
		PictureStorePath: pbSizeInfo.PictureStorePath,
	}
}

func toDbElementSelectSizeRecord(goodID, elementID, elementSizeInfoID int64) *model.ElementSelectSizeRecord {
	return &model.ElementSelectSizeRecord{
		GoodID:           goodID,
		ElementID:        elementID,
		SelectSizeInfoID: elementSizeInfoID,
	}
}
