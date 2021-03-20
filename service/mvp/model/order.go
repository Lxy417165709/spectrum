package model

import (
	"spectrum/common/pb"
	"time"
)

type Order struct {
	ID              int64     `gorm:"id"`
	CreatedAt       time.Time `gorm:"created_at"`
	UpdatedAt       time.Time `gorm:"updated_at"`

	Expense           float64 `json:"expense"`
	CheckOutAt time.Time   `gorm:"check_out_at"`
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
	if o.CheckOutAt.Unix() != 0 {
		return &pb.ExpenseInfo{
			NonFavorExpense:   o.NonFavorExpense,
			CheckOutAt: o.CheckOutAt.Unix(),
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
		CheckOutAt: 0,
		Expense:           GetFavorExpense(deskExpense+goodsExpense, favors),
	}
}
