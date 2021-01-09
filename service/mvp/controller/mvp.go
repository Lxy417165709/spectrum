package controller

import (
	"context"
	"github.com/astaxie/beego/logs"
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
	logs.Info("AddGood", ctx, req)

	var res pb.AddGoodRes

	if err := createGood(req.Good, req.ClassName); err != nil {
		logger.Error("Fail to finish createGood",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}

	return &res, nil
}

func (MvpServer) GetAllGoodClasses(ctx context.Context, req *pb.GetAllGoodClassesReq) (*pb.GetAllGoodClassesRes, error) {
	logs.Info("GetAllGoodClasses", ctx, req)
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
	logs.Info("AddGoodClass", ctx, req)

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
	logs.Info("OrderGood", ctx, req)

	var res pb.OrderGoodRes

	for _, good := range req.Goods {

		// 生成货物编号
		dbGood := &model.Good{
			//Name: good.MainElement.Name,
		}
		good.Id = int64(dbGood.ID)
		if err := createGood(good, req.ClassName); err != nil {
			logger.Error("Fail to finish createGood",
				zap.Any("req", req),
				zap.Error(err))
			return nil, err
		}

		// todo: 将货物与Desk联结起来
		if err := dao.DeskGoodRecordDao.Create(&model.DeskGoodRecord{
			GoodID:          int(dbGood.ID),
			DeskID:          int(req.DeskID),
			MainElementName: good.MainElement.Name,
		}); err != nil {
			// todo:log
			return nil, err
		}

	}
	return &res, nil

}

func (MvpServer) OpenDesk(ctx context.Context, req *pb.OpenDeskReq) (*pb.OpenDeskRes, error) {
	logs.Info("OpenDesk", ctx, req)
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
	logs.Info("GetDesk", ctx, req)

	var res pb.GetDeskRes

	res.Desk = getDesk(int(req.DeskID))
	return &res, nil
}

func (MvpServer) AddDesk(ctx context.Context, req *pb.AddDeskReq) (*pb.AddDeskRes, error) {
	logger.Info("AddDesk", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.AddDeskRes

	// 1. 创建
	if err := dao.SpaceDao.Create(&model.Space{
		Name:          req.Desk.Space.Name,
		Price:         float64(req.Desk.Space.Price),
		PriceRuleType: req.Desk.Space.PriceRuleType,
	}); err != nil {
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
	if err := dao.DeskDao.Update(map[string]interface{}{
		"id":            req.DeskID,
		"end_timestamp": time.Now().Unix(),
	}); err != nil {
		logger.Error("Fail to finish DeskDao.Update",
			zap.Any("req", req),
			zap.Error(err))
		return nil, err
	}
	return &res, nil
}

func (MvpServer) CheckOut(ctx context.Context, req *pb.CheckOutReq) (*pb.CheckOutRes, error) {
	logger.Info("CheckOut", zap.Any("ctx", ctx), zap.Any("req", req))
	var res pb.CheckOutRes
	desk := getDesk(int(req.DeskID))
	// todo: desk 可能等于 nil
	deskExpense := getDeskExpense(desk)
	res.Expanse = deskExpense
	return &res, nil
}

func getDeskExpense(desk *pb.Desk) float64 {
	enjoyHours := time.Unix(desk.EndTimestamp, 0).Sub(time.Unix(desk.StartTimestamp, 0)).Hours()
	enjoyExpense := desk.Space.Price * enjoyHours
	goodsExpense := getGoodsExpense(desk.Goods...)
	return enjoyExpense + goodsExpense
}
func getGoodsExpense(goods ...*pb.Good) float64 {
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
func getDesk(deskID int) *pb.Desk {
	var pbDesk pb.Desk
	pbDesk.Goods = getDeskGoods(deskID)
	desk, err := dao.DeskDao.Get(deskID)
	if err != nil {
		// todo: log
		return nil
	}
	pbDesk.StartTimestamp = desk.StartTimestamp
	pbDesk.EndTimestamp = desk.EndTimestamp
	pbDesk.Space = &pb.Space{
		Name:          desk.SpaceName,
		Num:           int64(desk.SpaceNum),
		PriceRuleType: desk.PriceRuleType,
		Price:         desk.Price,
	}
	return &pbDesk

}
func getDeskGoods(deskID int) []*pb.Good {
	records, err := dao.DeskGoodRecordDao.GetByDeskID(deskID)
	if err != nil {
		// todo: log
		return nil
	}
	var goods []*pb.Good
	for _, record := range records {
		goods = append(goods, getGood(record.GoodID, record.MainElementName))
	}
	return goods
}
func getGood(goodID int, mainElementName string) *pb.Good {
	return &pb.Good{
		Id:             int64(goodID),
		MainElement:    getMainElement(goodID, mainElementName),
		AttachElements: getAttachElements(goodID, mainElementName),
	}
}
func getMainElement(goodID int, mainElementName string) *pb.Element {
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
func getAttachElements(goodID int, mainElementName string) []*pb.Element {
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

func createGood(good *pb.Good, className string) error {
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

	for _, sizeInfo := range good.MainElement.SizeInfos {
		if err := dao.ElementDao.Create(&model.Element{
			Name:             good.MainElement.Name,
			Size:             sizeInfo.Size,
			Price:            float64(sizeInfo.Price),
			PictureStorePath: sizeInfo.PictureStorePath,
			Type:             pb.ElementType_Main,
			ClassName:        className,
		}); err != nil {
			logger.Error("Fail to finish ElementDao.Create", zap.Error(err))
			return err
		}
	}
	if err := dao.MainElementSizeRecordDao.Create(&model.MainElementSizeRecord{
		GoodID:          int(good.Id),
		MainElementName: good.MainElement.Name,
		SelectSize:      good.MainElement.SizeInfos[0].Size,
	}); err != nil {
		logger.Error("Fail to finish MainElementSizeRecordDao.Create", zap.Error(err))
		return err
	}

	// 3. 主元素归类 (todo: 类不存在时才创建)
	if err := dao.ElementClassDao.Create(&model.ElementClass{
		Name: className,
	}); err != nil {
		logger.Error("Fail to finish ElementClassDao.Create", zap.Error(err))
		return err
	}

	// 4. 创建主元素、附属元素的对应关系
	for _, attachElement := range good.AttachElements {
		if err := dao.MainElementAttachElementRecordDao.Create(&model.MainElementAttachElementRecord{
			MainElementName:   good.MainElement.Name,
			AttachElementName: attachElement.Name,
			SelectSize:        getSelectSizeInfo(attachElement.SizeInfos).Size,
			GoodID:            int(good.Id),
		}); err != nil {
			logger.Error("Fail to finish MainElementAttachElementDao.Create", zap.Error(err))
			return err
		}
	}
	return nil
}
