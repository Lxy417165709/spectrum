package model

import (
	"sort"
	"spectrum/common/pb"
)

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
