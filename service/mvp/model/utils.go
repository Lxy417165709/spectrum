package model

import (
	"fmt"
	"go.uber.org/zap"
	"sort"
	"spectrum/common/logger"
	"spectrum/common/pb"
	"strconv"
	"time"
)

func GetPbElementSelectSizeInfo(element *pb.Element) *pb.SizeInfo {
	return element.SizeInfos[element.SelectedIndex]
}

func GetFavorExpense(nonFavorExpense float64, pbFavors []*pb.Favor) float64 {
	favors := make([]Favor, 0)
	for _, pbFavor := range pbFavors {
		favor, _ := GetFavor(pbFavor)
		favors = append(favors, favor)
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

func GetFavor(favor *pb.Favor) (Favor, error) {
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
	return &None{}, nil // 兜底
}

func GetDbPrice(priceString string) float64 {
	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		logger.Error("Fail to finish strconv.ParseFloat", zap.String("priceString", priceString))
	}
	return price
}

func GetPbPrice(price float64) string {
	return fmt.Sprintf("%0.2f", price)
}

func ToTime(timestamp int64) time.Time {
	if timestamp != 0 {
		return time.Unix(timestamp, 0)
	}
	return NilTime
}
