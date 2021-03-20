package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
	"time"
)

type Good struct {
	gorm.Model
	Name    string `json:"name"`
	OrderID int64  `json:"order_id"`

	Expense         float64   `json:"expense"`
	CheckOutAt      time.Time `gorm:"check_out_at"`
	NonFavorExpense float64   `json:"non_favor_expense"`
}

func (g *Good) GetID() int64 {
	return int64(g.ID)
}

func (g *Good) TableName() string {
	return g.GetName()
}

func (*Good) GetName() string {
	return ChargeableObjectNameOfGood
}

func (g *Good) GetExpenseInfo(mainElement *pb.Element, attachElement []*pb.Element, favors []*pb.Favor) *pb.ExpenseInfo {
	if g.CheckOutAt.Unix() != 0 {
		return &pb.ExpenseInfo{
			NonFavorExpense: g.NonFavorExpense,
			CheckOutAt:      g.CheckOutAt.Unix(),
			Expense:         g.Expense,
		}
	}
	nonFavorExpense := g.getNonFavorExpense(append(attachElement, mainElement))
	return &pb.ExpenseInfo{
		NonFavorExpense: nonFavorExpense,
		CheckOutAt:      0,
		Expense:         GetFavorExpense(nonFavorExpense, favors),
	}
}

func (g *Good) getNonFavorExpense(elements []*pb.Element) float64 {
	expense := 0.0
	for _, element := range elements {
		priceString := GetSelectSizeInfo(element).Price
		expense += GetDbPrice(priceString)
	}
	return expense
}
