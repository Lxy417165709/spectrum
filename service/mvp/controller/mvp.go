package controller

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
	"time"
)

type MvpServer struct {
	pb.UnimplementedMvpServer
}

func (MvpServer) AddGood(ctx context.Context, req *pb.AddGoodReq) (*pb.AddGoodRes, error) {
	logger.Info("AddGood", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddGoodRes
	if err := createElement(req.Good.MainElement, req.ClassName); err != nil {
		logger.Error("Fail to finish createElement",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	if err := writeGoodSizeToDB(req.Good); err != nil {
		logger.Error("Fail to finish writeGoodSizeToDB",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	return &res, nil
}

func (MvpServer) AddElement(ctx context.Context, req *pb.AddElementReq) (*pb.AddElementRes, error) {
	logger.Info("AddElement", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddElementRes

	if err := createElement(req.Element, req.ClassName); err != nil {
		logger.Error("Fail to finish createElement",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

func (MvpServer) GetAllGoodClasses(ctx context.Context, req *pb.GetAllGoodClassesReq) (*pb.GetAllGoodClassesRes, error) {
	logger.Info("GetAllGoodClasses", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetAllGoodClassesRes

	// 1. 获得主元素的所有类
	classes, err := dao.ElementClassDao.GetAllClasses()
	if err != nil {
		logger.Error("Fail to finish ElementClassDao.GetMainElementClass",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	// 2. 形成 GoodClasses
	for _, class := range classes {
		res.GoodClasses = append(res.GoodClasses, &pb.GoodClass{
			Name:  class.Name,
			Goods: getClassGoods(class.Name),
		})
	}

	return &res, nil
}

func (MvpServer) AddGoodClass(ctx context.Context, req *pb.AddGoodClassReq) (*pb.AddGoodClassRes, error) {
	logger.Info("AddGoodClass", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.AddGoodClassRes
	// todo: 判断类名是否为空、是否存在

	// 1. 创建商品类
	if err := dao.ElementClassDao.Create(&model.ElementClass{
		Name: req.GoodClass.Name,
	}); err != nil {
		logger.Error("Fail to finish ElementDao.Create",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

func (MvpServer) OrderGood(ctx context.Context, req *pb.OrderGoodReq) (*pb.OrderGoodRes, error) {
	logger.Info("OrderGood", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.OrderGoodRes

	for _, good := range req.Goods {
		// 生成货物编号, 将货物与桌位联结
		dbGood := &model.Good{
			Name:   good.MainElement.Name,
			DeskID: req.DeskID,
		}
		if err := dao.GoodDao.Create(dbGood); err != nil {
			// todo:log
			return nil, err
		}
		good.Id = int64(dbGood.ID)
		if err := writeGoodSizeToDB(good); err != nil {
			logger.Error("Fail to finish createGood",
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}

	}
	return &res, nil

}

func (MvpServer) OpenDesk(ctx context.Context, req *pb.OpenDeskReq) (*pb.OpenDeskRes, error) {
	logger.Info("OpenDesk", zap.Any("ctx", ctx), zap.Any("req", req))

	// todo: 通过 SpaceName Num 查询 Price PriceRuleType
	var res pb.OpenDeskRes

	desk := &model.Desk{
		SpaceName:      req.SpaceName,
		SpaceNum:       int(req.SpaceNum),
		StartTimestamp: time.Now().Unix(),
	}
	if err := dao.DeskDao.Create(desk); err != nil {
		//todo: log
		return nil, err
	}

	//order := &model.Order{
	//	DeskID: int(desk.ID),
	//}
	//if err := dao.OrderDao.Create(order); err != nil {
	//	//todo: log
	//	return nil, err
	//}

	res.DeskID = int64(desk.ID)
	return &res, nil
}

func (MvpServer) GetDesk(ctx context.Context, req *pb.GetDeskReq) (*pb.GetDeskRes, error) {
	logger.Info("GetDesk", zap.Any("ctx", ctx), zap.Any("req", req))

	var res pb.GetDeskRes

	res.Desk = getDesk(req.DeskID)
	return &res, nil
}

func (MvpServer) AddSpace(ctx context.Context, req *pb.AddSpaceReq) (*pb.AddSpaceRes, error) {
	logger.Info("AddDesk", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.AddSpaceRes

	// 1. 创建
	if err := dao.SpaceDao.Create(getDbSpace(req.Space)); err != nil {
		logger.Error("Fail to finish SpaceDao.Create",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

func (MvpServer) CloseDesk(ctx context.Context, req *pb.CloseDeskReq) (*pb.CloseDeskRes, error) {
	logger.Info("CloseDesk", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.CloseDeskRes
	if err := closeDeskIfOpening(req.DeskID, req.EndTimestamp); err != nil {
		logger.Error("Fail to finish closeDesk",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

func (MvpServer) FormExpense(ctx context.Context, req *pb.FormExpenseReq) (*pb.FormExpenseRes, error) {
	logger.Info("FormExpense", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.FormExpenseRes
	desk := getDesk(req.DeskID)
	if desk == nil {
		err := fmt.Errorf("desk(id = %d) is non", req.DeskID)
		logger.Error("Desk is non", zap.Error(err))
		return nil, err
	}
	formDeskExpense(desk)
	writeToDB(desk, "expense")
	res.Desk = desk
	return &res, nil
}

func (MvpServer) CheckOut(ctx context.Context, req *pb.CheckOutReq) (*pb.CheckOutRes, error) {
	// todo: 形成订单
	logger.Info("CheckOut", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.CheckOutRes

	writeToDB(req.Desk, "had_check_out")
	return &res, nil
}

func writeToDB(desk *pb.Desk, field string) {
	if err := dao.DeskDao.Update(getToMap(desk, field)); err != nil {
		logger.Error("Fail to finish DeskDao.Update", zap.Error(err))
		return
	}

	for _, good := range desk.Goods {
		if err := dao.GoodDao.Update(getToMap(good, field)); err != nil {
			logger.Error("Fail to finish GoodDao.Update", zap.Error(err))
			return
		}
	}
}

func getToMap(obj interface{}, field string) map[string]interface{} {
	to := make(map[string]interface{})
	switch obj.(type) {
	case *pb.Desk:
		desk := obj.(*pb.Desk)
		to["id"] = desk.Id
		if field == "expense" {
			to["expense"] = desk.ExpenseInfo.Expense
		}
		if field == "had_check_out" {
			to["had_check_out"] = desk.ExpenseInfo.HadCheckOut
		}
	case *pb.Good:
		good := obj.(*pb.Good)
		to["id"] = good.Id
		if field == "expense" {
			to["expense"] = good.ExpenseInfo.Expense
		}
		if field == "had_check_out" {
			to["had_check_out"] = good.ExpenseInfo.HadCheckOut
		}
	default:
		logger.Error("UnFix interface", zap.Any("obj", obj), zap.String("field", field))
		panic("UnFix interface")
	}
	return to
}

func formDeskExpense(desk *pb.Desk) float64 {
	// todo: priceRuleType 还未应用
	if desk.EndTimestamp == 0 {
		closeTimestamp := time.Now().Unix()
		if err := closeDeskIfOpening(desk.Id, closeTimestamp); err != nil {
			logger.Error("Fail to finish closeDesk",
				zap.Error(err))
			return 0
		}
		desk.EndTimestamp = closeTimestamp
	}
	enjoyHours := time.Unix(desk.EndTimestamp, 0).Sub(time.Unix(desk.StartTimestamp, 0)).Hours()
	enjoyExpense := desk.Space.Price * enjoyHours
	goodsExpense := formGoodsExpense(desk.Goods...)
	desk.ExpenseInfo = &pb.ExpenseInfo{
		Expense: enjoyExpense,
	}
	return enjoyExpense + goodsExpense
}

func formGoodsExpense(goods ...*pb.Good) float64 {
	// todo: getSelectSizeInfo 可能返回nil
	// todo: 折扣计算还未进行
	allExpense := 0.0
	for _, good := range goods {
		mainElementExpense := getSelectSizeInfo(good.MainElement.SizeInfos).Price
		attachElementsExpense := 0.0
		for _, attachElement := range good.AttachElements {
			attachElementsExpense += getSelectSizeInfo(attachElement.SizeInfos).Price
		}
		goodExpense := mainElementExpense + attachElementsExpense
		allExpense += goodExpense

		good.ExpenseInfo = &pb.ExpenseInfo{
			Expense: goodExpense,
		}
	}
	return allExpense
}

func getClassGoods(className string) []*pb.Good {
	var goods []*pb.Good
	for _, mainElementName := range getElementNames(className) {
		goods = append(goods, getGood(0, mainElementName))
	}
	return goods
}
func getDesk(deskID int64) *pb.Desk {
	desk, err := dao.DeskDao.Get(deskID)
	if err != nil {
		// todo: log
		return nil
	}

	space, err := dao.SpaceDao.Get(desk.SpaceName, int64(desk.SpaceNum))
	if err != nil {
		// todo: log
		return nil
	}

	return &pb.Desk{
		Goods:          getDeskGoods(deskID),
		StartTimestamp: desk.StartTimestamp,
		EndTimestamp:   desk.EndTimestamp,
		Space:          space.ToPb(),
	}
}

func getDeskGoods(deskID int64) []*pb.Good {
	dbGoods, err := dao.GoodDao.GetByDeskID(deskID)
	if err != nil {
		// todo: log
		return nil
	}
	var goods []*pb.Good
	for _, dbGood := range dbGoods {
		goods = append(goods, getGood(int64(dbGood.ID), dbGood.Name))
	}
	return goods
}
func getGood(goodID int64, mainElementName string) *pb.Good {
	return &pb.Good{
		Id:             int64(goodID),
		MainElement:    getMainElement(goodID, mainElementName),
		AttachElements: getAttachElements(goodID, mainElementName),
	}
}
func getMainElement(goodID int64, mainElementName string) *pb.Element {
	mainElements, err := dao.ElementDao.GetByName(mainElementName)
	if err != nil {
		// todo: log
		return nil
	}
	var sizeRecord *model.MainElementSizeRecord
	if goodID == 0 {
		sizeRecord, err = dao.MainElementSizeRecordDao.GetByMainElementName(mainElementName)
	} else {
		sizeRecord, err = dao.MainElementSizeRecordDao.GetByGoodID(goodID)
	}
	if err != nil {
		// todo: log
		return nil
	}
	if sizeRecord == nil {
		// todo: log
		return nil
	}
	return &pb.Element{
		Name:      mainElementName,
		SizeInfos: getSizeInfos(sizeRecord.SelectSize, mainElements),
	}
}
func getAttachElements(goodID int64, mainElementName string) []*pb.Element {
	var attachElements []*pb.Element
	var attachRecords []*model.MainElementAttachElementRecord
	var err error
	if goodID == 0 {
		attachRecords, err = dao.MainElementAttachElementRecordDao.GetByMainElementName(mainElementName)
	} else {
		attachRecords, err = dao.MainElementAttachElementRecordDao.GetByGoodID(goodID)
	}
	if err != nil {
		logger.Error("Fail to finish MainElementAttachElementRecordDao.GetByMainElementName", zap.Error(err))
		return nil
	}
	for _, attachRecord := range attachRecords {
		elements, err := dao.ElementDao.GetByName(attachRecord.AttachElementName)
		if err != nil {
			logger.Error("Fail to finish ElementDao.GetByName", zap.Error(err))
			return nil
		}
		attachElements = append(attachElements, &pb.Element{
			Name:      attachRecord.AttachElementName,
			SizeInfos: getSizeInfos(attachRecord.SelectSize, elements),
		})
	}
	return attachElements
}
func getSizeInfos(selectSize string, sameNameElements []*model.Element) []*pb.SizeInfo {
	var sizeInfos []*pb.SizeInfo
	for _, element := range sameNameElements {
		sizeInfos = append(sizeInfos, &pb.SizeInfo{
			Size:             element.Size,
			Price:            element.Price,
			PictureStorePath: element.PictureStorePath,
			IsSelected:       selectSize == element.Size,
		})
	}
	return sizeInfos
}

func getElementNames(className string) []string {
	var isLiving = make(map[string]bool)
	elements, err := dao.ElementDao.GetByClassName(className)
	var elementNames []string
	if err != nil {
		logger.Error("Fail to finish ElementDao.GetByClassName", zap.Error(err))
		return nil
	}
	for _, element := range elements {
		if isLiving[element.Name] {
			continue
		}
		isLiving[element.Name] = true
		elementNames = append(elementNames, element.Name)
	}
	return elementNames
}

func getSelectSizeInfo(infos []*pb.SizeInfo) *pb.SizeInfo {
	for _, sizeInfo := range infos {
		if sizeInfo.IsSelected {
			return sizeInfo
		}
	}
	return nil
}

func writeGoodSizeToDB(good *pb.Good) error {
	// 1. 判断主元素是否存在
	//elements, err := dao.ElementDao.GetByName(good.MainElement.Name)
	//if err != nil {
	//	logger.Error("Fail to finish ElementDao.GetByName", zap.Error(err))
	//	return err
	//}
	//if len(elements) != 0 {
	//	err := ers.New(0, "商品(%s)已存在", good.MainElement.Name)
	//	logger.Error("Fail to finish ElementDao.GetByName",
	//		zap.Any("elements", elements),
	//		zap.Error(err))
	//	return err
	//}

	// 2. 创建主元素
	if len(good.MainElement.SizeInfos) == 0 {
		return nil
	}
	//if good.Id == 0 {
	//	if err := createElement(good.MainElement, className); err != nil {
	//		// todo: log
	//		return err
	//	}
	//}

	if err := dao.MainElementSizeRecordDao.Create(&model.MainElementSizeRecord{
		GoodID:          good.Id,
		MainElementName: good.MainElement.Name,

		SelectSize: good.MainElement.SizeInfos[0].Size,
	}); err != nil {
		logger.Error("Fail to finish MainElementSizeRecordDao.Create", zap.Error(err))
		return err
	}

	// 3. 主元素归类 (todo: 类不存在时才创建) todo: 这里默认类已存在
	//if err := dao.ElementClassDao.Create(&model.ElementClass{
	//	Name: className,
	//}); err != nil {
	//	logger.Error("Fail to finish ElementClassDao.Create", zap.Error(err))
	//	return err
	//}

	// 4. 创建主元素、附属元素的对应关系
	for _, attachElement := range good.AttachElements {
		if err := dao.MainElementAttachElementRecordDao.Create(&model.MainElementAttachElementRecord{
			GoodID:            good.Id,
			MainElementName:   good.MainElement.Name,
			AttachElementName: attachElement.Name,

			SelectSize: getSelectSizeInfo(attachElement.SizeInfos).Size,
		}); err != nil {
			logger.Error("Fail to finish MainElementAttachElementDao.Create", zap.Error(err))
			return err
		}
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

func getDbElements(pbElement *pb.Element, className string) []*model.Element {
	var result []*model.Element

	for _, sizeInfo := range pbElement.SizeInfos {
		result = append(result, &model.Element{
			Name:      pbElement.Name,
			Type:      pbElement.Type,
			ClassName: className,

			Size:             sizeInfo.Size,
			Price:            sizeInfo.Price,
			PictureStorePath: sizeInfo.PictureStorePath,
		})
	}
	return result
}
func getDbSpace(pbSpace *pb.Space) *model.Space {
	return &model.Space{
		Name:          pbSpace.Name,
		Num:           pbSpace.Num,
		Price:         pbSpace.Price,
		PriceRuleType: pbSpace.PriceRuleType,
	}
}

func closeDeskIfOpening(deskID int64, endTimestamp int64) error {
	desk, err := dao.DeskDao.Get(deskID)
	if err != nil {
		logger.Error("Fail to finish DeskDao.Get",
			zap.Error(err))
		return err
	}

	isOpening := desk.EndTimestamp == 0
	if !isOpening {
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
