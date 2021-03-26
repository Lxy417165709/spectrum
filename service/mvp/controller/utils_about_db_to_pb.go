package controller

import (
	"go.uber.org/zap"
	"sort"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/utils"
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

	// 优化: 这里只需要获得 deskID，spaceID 就可以了
	desk, errResult := dao.DeskDao.GetByOrderID(orderID)
	if errResult != nil {
		return nil
	}
	if desk == nil {
		logger.Warn("Desk not exist", zap.Any("orderID", orderID))
		return nil
	}

	pbDesk := getPbDesk(desk.ID, desk.SpaceID)
	pbGoods := getOrderPbGoods(orderID)
	favors, errResult := getFavors(order.GetChargeableObjectName(), order.GetID())
	if errResult != nil {
		return nil
	}
	return &pb.Order{
		Id:          orderID,
		Desk:        pbDesk,
		Goods:       pbGoods,
		Favors:      favors,
		ExpenseInfo: order.GetExpenseInfo(pbDesk, pbGoods, favors), // todo:
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

func getPbDesk(deskID int64, spaceID int64) *pb.Desk {
	space, errResult := dao.SpaceDao.Get(spaceID)
	if errResult != nil {
		return nil
	}

	spaceClass, errResult := dao.SpaceClassDao.Get(space.ClassID)
	if errResult != nil {
		return nil
	}
	if deskID == 0 {
		return &pb.Desk{
			Id:      0,
			OrderID: 0,
			Space:   space.ToPb(getDbSpaceClassByID(space.ClassID).Name),
			StartAt: utils.NilTime.Unix(),
			EndAt:   utils.NilTime.Unix(),
			Favors:  []*pb.Favor{},
			ExpenseInfo: &pb.ExpenseInfo{
				NonFavorExpense: 0,
				CheckOutAt:      utils.NilTime.Unix(),
				Expense:         0,
			},
		}
	}

	desk, errResult := dao.DeskDao.Get(deskID, spaceID)
	if errResult != nil {
		return nil
	}

	favors, errResult := getFavors(desk.GetChargeableObjectName(), desk.GetID())
	if errResult != nil {
		return nil
	}
	return &pb.Desk{
		Id:          desk.ID,
		Space:       space.ToPb(spaceClass.Name),
		StartAt:     desk.StartAt.Unix(),
		EndAt:       desk.EndAt.Unix(),
		Favors:      favors,
		ExpenseInfo: desk.GetExpenseInfo(space.BillingType, space.Price, favors),
		OrderID:     desk.OrderID,
	}
}

func getPbGood(goodID int64, mainElementID int64) *pb.Good {
	mainElement := getPbElement(goodID, mainElementID)
	attachElements := getPbAttachElements(goodID, mainElementID)
	if goodID == 0 {
		return &pb.Good{
			Id:             0,
			MainElement:    mainElement,
			AttachElements: attachElements,
			Favors:         []*pb.Favor{},
			ExpenseInfo: &pb.ExpenseInfo{
				NonFavorExpense: 0,
				CheckOutAt:      utils.NilTime.Unix(),
				Expense:         0,
			},
		}
	}

	good, errResult := dao.GoodDao.Get(goodID, mainElementID)
	if errResult != nil {
		return nil
	}
	favors, errResult := getFavors(good.GetChargeableObjectName(), good.GetID())
	if errResult != nil {
		return nil
	}
	return &pb.Good{
		Id:             goodID,
		MainElement:    mainElement,
		AttachElements: attachElements,
		Favors:         favors,
		ExpenseInfo:    good.GetExpenseInfo(mainElement, attachElements, favors), // todo: 之后再说
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

func getFavors(chargeableObjName string, chargeableObjID int64) ([]*pb.Favor, error) {
	// 1. 查询
	records, errResult := dao.FavorRecordDao.GetFavorRecords(chargeableObjName, chargeableObjID)
	if errResult != nil {
		return nil, errResult
	}

	// 2. 转换
	result := make([]*pb.Favor, 0)
	for _, record := range records {
		result = append(result, record.ToPb())
	}

	// 3. 返回
	return result, nil
}
