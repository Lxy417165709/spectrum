package controller

import (
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
)

func getClassGoods(className string) []*pb.Good {
	var goods []*pb.Good
	for _, mainElementName := range getElementNames(className) {
		pbGood := getPbGood(&model.Good{
			Name:              mainElementName,
			OrderID:           0,
			Expense:           0,
			CheckOutTimestamp: 0,
			NonFavorExpense:   0,
		})
		goods = append(goods, pbGood)
		logger.Info("Get pb good", zap.Any("pbGood", pbGood))
	}
	return goods
}

// 返回的 desk:
// 已结账时: 返回结账的金额信息
// 未结账时: 返回最新的金额信息
func getPbDesk(desk *model.Desk) *pb.Desk {
	space, err := dao.SpaceDao.Get(desk.SpaceName, desk.SpaceClassName)
	if err != nil {
		// todo: log
		return nil
	}
	favor := getFavors(desk)
	return &pb.Desk{
		Id:             int64(desk.ID),
		Space:          space.ToPb(),
		StartTimestamp: desk.StartTimestamp,
		EndTimestamp:   desk.EndTimestamp,
		Favors:         favor,
		ExpenseInfo:    desk.GetExpenseInfo(space.BillingType, space.Price, favor),
	}
}

func getPbOrder(orderID int64) *pb.Order {
	order, err := dao.OrderDao.Get(orderID)
	if err != nil {
		// todo: log
		return nil
	}
	desk, err := dao.DeskDao.GetByOrderID(orderID)
	if err != nil {
		// todo: log
		return nil
	}
	pbDesk := getPbDesk(desk)
	pbGoods := getOrderPbGoods(orderID)
	favors := getFavors(order)
	return &pb.Order{
		Id:          orderID,
		Desk:        pbDesk,
		Goods:       pbGoods,
		Favors:      favors,
		ExpenseInfo: order.GetExpenseInfo(pbDesk, pbGoods, favors),
	}
}

func getOrderPbGoods(orderID int64) []*pb.Good {
	dbGoods, err := dao.GoodDao.GetByOrderID(orderID)
	if err != nil {
		// todo: log
		return nil
	}
	var goods []*pb.Good
	for _, dbGood := range dbGoods {
		goods = append(goods, getPbGood(dbGood))
	}
	return goods
}

// 返回的 good:
// 已结账时: 返回结账的金额信息
// 未结账时: 返回最新的金额信息
func getPbGood(good *model.Good) *pb.Good {
	mainElement := getMainElement(int64(good.ID), good.Name)
	attachElements := getAttachElements(int64(good.ID), good.Name)
	favors := getFavors(good)
	return &pb.Good{
		Id:             int64(good.ID),
		MainElement:    mainElement,
		AttachElements: attachElements,
		Favors:         favors,
		//ExpenseInfo:    good.GetExpenseInfo(mainElement, attachElements, favors),	// todo: 之后再说
	}
}

func getMainElement(goodID int64, mainElementName string) *pb.Element {
	mainElements, err := dao.ElementDao.GetByName(mainElementName)
	if err != nil {
		// todo: log
		return nil
	}
	sizeRecord, err := dao.MainElementSizeRecordDao.GetByGoodIdAndMainElementName(goodID, mainElementName)
	if err != nil {
		// todo: log
		return nil
	}
	if sizeRecord == nil {
		logger.Warn("Size record is blank", zap.Any("goodID", goodID), zap.Any("mainElementName", mainElementName))
		return nil
	}
	sizeInfos := model.GetSizeInfos(mainElements)
	selectedSizeInfoIndex := int32(GetSelectedIndex(sizeInfos, sizeRecord.SelectSize))
	logger.Info("GetSelectedIndex",
		zap.Any("mainElementName", mainElementName),
		zap.Any("selectedSizeInfoIndex", selectedSizeInfoIndex),
		zap.Any("sizeInfos", sizeInfos),
		zap.Any("selectSize", sizeRecord.SelectSize))
	return &pb.Element{
		Name:                  mainElementName,
		Type:                  pb.ElementType_Main,
		SelectedIndex: selectedSizeInfoIndex,
		SizeInfos:             sizeInfos,
	}
}

// todo: 2021年03月16日00:41:27 这里有问题
func GetSelectedIndex(sizeInfos []*pb.SizeInfo, selectSize string) int {

	for index, sizeInfo := range sizeInfos {
		if sizeInfo.Size == selectSize {
			return index
		}
	}
	logger.Warn("Can not get selected Index", zap.Any("sizeInfos", sizeInfos), zap.Any("selectSize", selectSize))
	return 0
}

func getAttachElements(goodID int64, mainElementName string) []*pb.Element {
	var attachElements []*pb.Element
	attachRecords, err := dao.MainElementAttachElementRecordDao.GetByGoodIdAndMainElementName(goodID, mainElementName)
	if err != nil {
		logger.Error("Fail to finish MainElementAttachElementRecordDao.GetByGoodIdAndMainElementName", zap.Error(err))
		return nil
	}
	logger.Info("Success to get attachRecords", zap.Any("attachRecords", attachRecords))

	for _, attachRecord := range attachRecords {
		elements, err := dao.ElementDao.GetByName(attachRecord.AttachElementName)
		if err != nil {
			logger.Error("Fail to finish ElementDao.GetByName", zap.Error(err))
			return nil
		}
		if len(elements) == 0 {
			logger.Error("Element not exist", zap.String("name", attachRecord.AttachElementName))
			return nil
		}

		attachElements = append(attachElements, &pb.Element{
			Name:      attachRecord.AttachElementName, // 这里也等于 elements[0].Name
			Type:      elements[0].Type,               // 这里的元素包括了规格，但它们的 Type 是一样的
			SizeInfos: model.GetSizeInfos(elements),
		})
	}
	return attachElements
}

func getFavors(chargeableObj model.Chargeable) []*pb.Favor {
	records, err := dao.ChargeableObjectDao.GetFavorRecords(chargeableObj)
	if err != nil {
		// todo: log
		return nil
	}
	result := make([]*pb.Favor, 0)
	for _, record := range records {
		result = append(result, record.ToPb())
	}
	return result
}
