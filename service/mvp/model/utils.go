package model

import (
	"sort"
	"spectrum/common/pb"
)


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
