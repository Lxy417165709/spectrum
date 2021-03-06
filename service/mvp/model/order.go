package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
)

type Order struct {
	gorm.Model

	Expense           float64 `json:"expense"`
	CheckOutTimestamp int64   `json:"check_out_timestamp"`
	NonFavorExpense   float64 `json:"non_favor_expense"`
}

func (Order) GetName() string {
	return ChargeableObjectNameOfOrder
}

func (o *Order) GetID() int64 {
	return int64(o.ID)
}
func (o *Order) TableName() string {
	return o.GetName()
}

func (o *Order) GetExpenseInfo(desk *pb.Desk, goods []*pb.Good, favors []*pb.Favor) *pb.ExpenseInfo {
	if o.CheckOutTimestamp != 0 {
		return &pb.ExpenseInfo{
			NonFavorExpense:   o.NonFavorExpense,
			CheckOutTimestamp: o.CheckOutTimestamp,
			Expense:           o.Expense,
		}
	}

	// todo: 如果已结账，则不应该算入
	deskExpense := desk.GetExpenseInfo().Expense
	goodsExpense := 0.0
	for _, good := range goods {
		goodsExpense += good.GetExpenseInfo().Expense
	}
	return &pb.ExpenseInfo{
		NonFavorExpense:   deskExpense + goodsExpense,
		CheckOutTimestamp: 0,
		Expense:           GetFavorExpense(deskExpense+goodsExpense, favors),
	}
}
