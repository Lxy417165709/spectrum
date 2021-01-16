package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
	"time"
)

type Desk struct {
	gorm.Model
	StartTimestamp int64 `json:"start_timestamp"`
	EndTimestamp   int64 `json:"end_timestamp"`

	SpaceName string `json:"space_name"`
	SpaceNum  int    `json:"space_num"`

	Expense           float64 `json:"expense"`
	CheckOutTimestamp int64   `json:"check_out_timestamp"`
	NonFavorExpense   float64 `json:"non_favor_expense"`
}

func (*Desk) TableName() string {
	return "desk"
}

func (d *Desk) GetExpenseInfo(nonFavorExpense float64, favors []*pb.Favor) *pb.ExpenseInfo {
	if d.CheckOutTimestamp != 0 {
		return &pb.ExpenseInfo{
			NonFavorExpense:   d.NonFavorExpense,
			CheckOutTimestamp: d.CheckOutTimestamp,
			Expense:           d.Expense,
		}
	}
	return &pb.ExpenseInfo{
		NonFavorExpense:   nonFavorExpense,
		CheckOutTimestamp: d.CheckOutTimestamp,
		Expense:           GetFavorExpense(nonFavorExpense, favors),
	}
}

func (d *Desk) GetExpense(pricePerHour float64) float64 {
	var endTimestamp int64
	if d.IsOpening() {
		endTimestamp = time.Now().Unix()
	} else {
		endTimestamp = d.EndTimestamp
	}
	hours := time.Unix(endTimestamp, 0).Sub(time.Unix(d.StartTimestamp, 0)).Hours()
	return hours * pricePerHour

}
func (d *Desk) IsOpening() bool {
	return d.EndTimestamp == 0
}
