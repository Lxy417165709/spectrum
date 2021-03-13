package model

import (
	"sort"
	"spectrum/common/pb"
)

func GetSelectSizeInfo(infos []*pb.SizeInfo) *pb.SizeInfo {
	for _, sizeInfo := range infos {
		if sizeInfo.IsSelected {
			return sizeInfo
		}
	}
	//return nil
	panic("No default selected size info")
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

func GetSizeInfos(selectSize string, sameNameElements []*Element) []*pb.SizeInfo {
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
