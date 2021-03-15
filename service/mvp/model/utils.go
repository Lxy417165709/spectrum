package model

import (
	"fmt"
	"go.uber.org/zap"
	"sort"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"strconv"
)

func GetSelectSizeInfo(element *pb.Element) *pb.SizeInfo {
	return element.SizeInfos[element.SelectedIndex]
}

func GetFavorExpense(nonFavorExpense float64, pbFavors []*pb.Favor) float64 {
	favors := make([]Favor, 0)
	for _, pbFavor := range pbFavors {
		favors = append(favors, GetFavor(pbFavor))
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

func GetFavor(favor *pb.Favor) Favor {
	// todo: 要确认优惠参数是合法的
	if favor == nil {
		return &None{}
	}
	switch favor.FavorType {
	case pb.FavorType_NONE:
		return (&None{}).ParseParameters(favor.Parameters)
	case pb.FavorType_REBATE:
		return (&Rebate{}).ParseParameters(favor.Parameters)
	case pb.FavorType_FULL_REDUCTION:
		return (&FullReduction{}).ParseParameters(favor.Parameters)
	case pb.FavorType_FREE:
		return (&Free{}).ParseParameters(favor.Parameters)
	}
	return &None{}
}

func GetSizeInfos(sameNameElements []*Element) []*pb.SizeInfo {
	var sizeInfos []*pb.SizeInfo

	for _, element := range sameNameElements {
		sizeInfos = append(sizeInfos, &pb.SizeInfo{
			Size:             element.Size,
			Price:            GetPbPrice(element.Price),
			PictureStorePath: element.PictureStorePath,
		})
	}
	return sizeInfos
}

func GetDbPrice(priceString string) float64 {
	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		logger.Error("Fail to finish strconv.ParseFloat", zap.String("priceString", priceString))
		// 这里不返回
	}
	return price
}

func GetPbPrice(price float64) string {
	return fmt.Sprintf("%0.2f", price)
}
