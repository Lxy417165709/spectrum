package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
)

type Good struct {
	gorm.Model
	Name   string `json:"name"`
	DeskID int64  `json:"desk_id"`

	Expense           float64 `json:"expense"`
	CheckOutTimestamp int64   `json:"check_out_timestamp"`
	NonFavorExpense   float64 `json:"non_favor_expense"`
}

func (g *Good) SetCheckOutTimestamp(timestamp int64) {
	g.CheckOutTimestamp = timestamp
}

func (g *Good) GetCheckOutTimestamp() int64 {
	return g.CheckOutTimestamp
}

func (g *Good) GetID() int64 {
	return int64(g.ID)
}

func (g *Good) SetID(id int64) {
	g.ID = uint(id)
}

func (*Good) TableName() string {
	return "good"
}

func (*Good) GetName() string {
	return ChargeableObjectNameOfGood
}

func (g *Good) GetExpenseInfo(mainElement *pb.Element, attachElement []*pb.Element, favors []*pb.Favor) *pb.ExpenseInfo {
	if g.CheckOutTimestamp != 0 {
		return &pb.ExpenseInfo{
			NonFavorExpense:   g.NonFavorExpense,
			CheckOutTimestamp: g.CheckOutTimestamp,
			Expense:           g.Expense,
		}
	}
	nonFavorExpense := g.getNonFavorExpense(append(attachElement, mainElement))
	return &pb.ExpenseInfo{
		NonFavorExpense:   nonFavorExpense,
		CheckOutTimestamp: 0,
		Expense:           GetFavorExpense(nonFavorExpense, favors),
	}
}

func (g *Good) getNonFavorExpense(elements []*pb.Element) float64 {
	expense := 0.0
	for _, element := range elements {
		expense += GetSelectSizeInfo(element.SizeInfos).Price
	}
	return expense
}
