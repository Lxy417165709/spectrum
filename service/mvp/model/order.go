package model

import (
	"spectrum/common/pb"
	"spectrum/service/mvp/utils"
	"time"
)

type Order struct {
	ID        int64     `gorm:"id"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`

	Expense         float64   `json:"expense"`
	CheckOutAt      time.Time `gorm:"check_out_at"`
	NonFavorExpense float64   `json:"non_favor_expense"`
}

func (o *Order) TableName() string {
	return o.GetChargeableObjectName()
}

func (o *Order) GetChargeableObjectName() string {
	return utils.ChargeableObjectNameOfOrder
}

func (o *Order) GetID() int64 {
	return o.ID
}

func (o *Order) GetExpenseInfo(desk *pb.Desk, goods []*pb.Good, favors []*pb.Favor) *pb.ExpenseInfo {
	// 1. 已结账时
	if o.CheckOutAt != utils.NilTime {
		return &pb.ExpenseInfo{
			NonFavorExpense: o.NonFavorExpense,
			CheckOutAt:      o.CheckOutAt.Unix(),
			Expense:         o.Expense,
		}
	}

	// 2. 未结账时
	deskExpense := 0.0
	if desk.GetExpenseInfo() != nil && desk.GetExpenseInfo().CheckOutAt == utils.NilTime.Unix(){
		deskExpense += desk.GetExpenseInfo().Expense
	}
	goodsExpense := 0.0
	for _, good := range goods {
		if desk.GetExpenseInfo() != nil  && good.GetExpenseInfo().CheckOutAt == utils.NilTime.Unix(){
			goodsExpense += good.GetExpenseInfo().Expense
		}
	}
	return &pb.ExpenseInfo{
		NonFavorExpense: deskExpense + goodsExpense,
		CheckOutAt:      utils.NilTime.Unix(),
		Expense:         GetFavorExpense(deskExpense+goodsExpense, favors),
	}
}
