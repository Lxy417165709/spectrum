package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
)

// todo: Favor 还未添加
type Good struct {
	gorm.Model
	Name   string `json:"name"`
	DeskID int64  `json:"desk_id"`

	Expense           float64 `json:"expense"`
	CheckOutTimestamp int64   `json:"check_out_timestamp"`
	NonFavorExpense   float64 `json:"non_favor_expense"`
}

func (*Good) TableName() string {
	return "good"
}

func (g *Good) GetExpenseInfo(nonFavorExpense float64, favors []*pb.Favor) *pb.ExpenseInfo {
	if g.CheckOutTimestamp != 0 {
		return &pb.ExpenseInfo{
			NonFavorExpense:   g.NonFavorExpense,
			CheckOutTimestamp: g.CheckOutTimestamp,
			Expense:           g.Expense,
		}
	}
	return &pb.ExpenseInfo{
		NonFavorExpense:   nonFavorExpense,
		CheckOutTimestamp: g.CheckOutTimestamp,
		Expense:           GetFavorExpense(nonFavorExpense, favors),
	}
}
