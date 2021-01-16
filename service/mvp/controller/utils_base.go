package controller

import (
	"go.uber.org/zap"
	"sort"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"spectrum/service/mvp/dao"
	"spectrum/service/mvp/model"
	"time"
)

// todo: 这个函数不优美
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
		logger.Error("Unfix interface", zap.Any("obj", obj), zap.String("field", field))
		panic("Unfix interface")
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
	deskExpense := getExpense(enjoyExpense, desk.Favors)
	desk.ExpenseInfo = &pb.ExpenseInfo{
		NonFavorExpense: enjoyExpense,
		Expense:         deskExpense,
	}
	return deskExpense + goodsExpense
}

func formGoodsExpense(goods ...*pb.Good) float64 {
	// todo: getSelectSizeInfo 可能返回nil
	allExpense := 0.0
	for _, good := range goods {
		mainElementExpense := getSelectSizeInfo(good.MainElement.SizeInfos).Price
		attachElementsExpense := 0.0
		for _, attachElement := range good.AttachElements {
			attachElementsExpense += getSelectSizeInfo(attachElement.SizeInfos).Price
		}
		nonFavorExpense := mainElementExpense + attachElementsExpense
		goodExpense := getExpense(nonFavorExpense, good.Favors)
		good.ExpenseInfo = &pb.ExpenseInfo{
			NonFavorExpense: nonFavorExpense,
			Expense:         goodExpense,
		}
		allExpense += goodExpense
	}
	return allExpense
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

func getExpense(nonFavorExpense float64, pbFavors []*pb.Favor) float64 {
	favors := make([]model.Favor, 0)
	for _, pbFavor := range pbFavors {
		favors = append(favors, model.GetFavor(pbFavor))
	}
	sort.Slice(favors, func(i, j int) bool {
		return favors[i].GetPriority() <= favors[j].GetPriority()
	})

	expense := nonFavorExpense

	for _, favor := range favors {
		expense = favor.GetExpense(expense)
	}
	return expense

}

