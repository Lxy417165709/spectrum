package controller

import (
	"go.uber.org/zap"
	"sort"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
)

// --------------------------------------------- ID ---------------------------------------------
func getClassGoods(classID int64) ([]*pb.Good, error) {
	var goods []*pb.Good
	elements, errResult := dao.ElementDao.GetByClassID(classID)
	if errResult != nil {
		return nil, errResult
	}
	for _, element := range elements {
		pbGood := getPbGood(0, element.ID)
		goods = append(goods, pbGood)
	}
	return goods, nil
}

func getPbOrder(orderID int64) *pb.Order {
	order, errResult := dao.OrderDao.Get(orderID)
	if errResult != nil {
		return nil
	}
	if order == nil {
		logger.Warn("Order not exist", zap.Any("orderID", orderID))
		return nil
	}

	desk, errResult := dao.DeskDao.GetByOrderID(orderID)
	if errResult != nil {
		return nil
	}
	if desk == nil {
		logger.Warn("Desk not exist", zap.Any("orderID", orderID))
		return nil
	}

	pbDesk := getPbDesk(desk)
	pbGoods := getOrderPbGoods(orderID)
	favors := getFavors(order)
	return &pb.Order{
		Id:     orderID,
		Desk:   pbDesk,
		Goods:  pbGoods,
		Favors: favors,
		ExpenseInfo: &pb.ExpenseInfo{
			NonFavorExpense: 0,
			CheckOutAt:      model.NilTime.Unix(),
			Expense:         0,
		},
		//ExpenseInfo: order.GetExpenseInfo(pbDesk, pbGoods, favors),	// todo:
	}
}

func getOrderPbGoods(orderID int64) []*pb.Good {
	dbGoods, errResult := dao.GoodDao.GetByOrderID(orderID)
	if errResult != nil {
		return nil
	}
	var goods []*pb.Good
	for _, dbGood := range dbGoods {
		goods = append(goods, getPbGood(dbGood.ID, dbGood.MainElementID))
	}
	return goods
}

// 返回的 good:
// 已结账时: 返回结账的金额信息
// 未结账时: 返回最新的金额信息
func getPbGood(goodID int64, mainElementID int64) *pb.Good {
	return &pb.Good{
		Id:             goodID,
		MainElement:    getPbElement(goodID, mainElementID),
		AttachElements: getPbAttachElements(goodID, mainElementID),
		//Favors:         getFavors(good),// todo: 之后再说
		//ExpenseInfo:    good.GetExpenseInfo(mainElement, attachElements, favors),	// todo: 之后再说
		Favors: []*pb.Favor{},
		ExpenseInfo: &pb.ExpenseInfo{
			NonFavorExpense: 0,
			CheckOutAt:      model.NilTime.Unix(),
			Expense:         0,
		},
	}
}

func getPbAttachElements(goodID, mainElementID int64) []*pb.Element {
	var attachElements []*pb.Element
	attachRecords, errResult := dao.MainElementAttachElementRecordDao.Get(goodID, mainElementID)
	if errResult != nil {
		logger.Error("Fail to finish MainElementAttachElementRecordDao.GetOne", zap.Error(errResult))
		return nil
	}
	logger.Info("Success to get attachRecords", zap.Any("attachRecords", attachRecords))

	for _, attachRecord := range attachRecords {
		attachElements = append(attachElements, getPbElement(goodID, attachRecord.AttachElementID))
	}
	return attachElements
}

func getPbElement(goodID, elementID int64) *pb.Element {
	// 1. 形成 pbSizeInfos、并排序
	dbSizeInfos, errResult := dao.ElementSizeInfoDao.Get(elementID)
	if errResult != nil {
		// todo: log
		return nil
	}
	var pbSizeInfos []*pb.SizeInfo
	for _, dbSizeInfo := range dbSizeInfos {
		pbSizeInfos = append(pbSizeInfos, dbSizeInfo.ToPb())
	}
	sort.Slice(pbSizeInfos, func(i, j int) bool {
		return pbSizeInfos[i].Id < pbSizeInfos[j].Id
	})

	// 2. 获取默认选择记录(管理页、用户差异处)
	selectSizeRecord, errResult := dao.ElementSelectSizeRecordDao.GetOne(goodID, elementID)
	if errResult != nil {
		// todo: log
		return nil
	}
	if selectSizeRecord == nil {
		logger.Warn("Size record is blank", zap.Any("goodID", goodID), zap.Any("elementID", elementID))
		return nil
	}

	// 4. 获取默认选择索引
	selectedSizeInfoIndex := int32(getSelectedIndex(pbSizeInfos, selectSizeRecord.SelectSizeInfoID))

	// 5. 查询数据库
	element, errResult := dao.ElementDao.Get(elementID)
	if errResult != nil {
		// todo: log
		return nil
	}

	// 6. 返回
	return &pb.Element{
		Id:            element.ID,
		Name:          element.Name,
		Type:          element.Type,
		SelectedIndex: selectedSizeInfoIndex,
		SizeInfos:     pbSizeInfos,
	}
}


// 返回的 desk:
// 已结账时: 返回结账的金额信息
// 未结账时: 返回最新的金额信息
func getPbDesk(desk *model.Desk) *pb.Desk {
	space, errResult := dao.SpaceDao.Get(desk.SpaceID)
	if errResult != nil {
		// todo: log
		return nil
	}
	favor := getFavors(desk)

	spaceClass, errResult := dao.SpaceClassDao.Get(space.ClassID)
	if errResult != nil {
		return nil
	}
	return &pb.Desk{
		Id:          desk.ID,
		Space:       space.ToPb(spaceClass.Name),
		StartAt:     desk.StartAt.Unix(),
		EndAt:       desk.EndAt.Unix(),
		Favors:      favor,
		ExpenseInfo: desk.GetExpenseInfo(space.BillingType, space.Price, favor),
		OrderID:     desk.OrderID,
	}
}

func getFavors(chargeableObj model.Chargeable) []*pb.Favor {
	records, err := dao.FavorRecordDao.GetFavorRecords(chargeableObj)
	if err != nil {
		return nil
	}
	result := make([]*pb.Favor, 0)
	for _, record := range records {
		result = append(result, record.ToPb())
	}
	return result
}