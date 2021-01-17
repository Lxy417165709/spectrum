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
	SpaceNum  int64  `json:"space_num"`

	Expense           float64 `json:"expense"`
	CheckOutTimestamp int64   `json:"check_out_timestamp"`
	NonFavorExpense   float64 `json:"non_favor_expense"`
}

func (d *Desk) SetCheckOutTimestamp(timestamp int64) {
	d.CheckOutTimestamp = timestamp
}

func (d *Desk) GetCheckOutTimestamp() int64 {
	return d.CheckOutTimestamp
}

func (*Desk) TableName() string {
	return "desk"
}
func (*Desk) GetName() string {
	return ChargeableObjectNameOfDesk
}
func (d *Desk) GetExpenseInfo(pricePerHour float64, favors []*pb.Favor) *pb.ExpenseInfo {
	if d.CheckOutTimestamp != 0 {
		return &pb.ExpenseInfo{
			NonFavorExpense:   d.NonFavorExpense,
			CheckOutTimestamp: d.CheckOutTimestamp,
			Expense:           d.Expense,
		}
	}
	nonFavorExpense := d.getNonFavorExpense(pricePerHour)
	return &pb.ExpenseInfo{
		NonFavorExpense:   nonFavorExpense,
		CheckOutTimestamp: 0,
		Expense:           GetFavorExpense(nonFavorExpense, favors),
	}
}

func (d *Desk) getNonFavorExpense(pricePerHour float64) float64 {
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

func (d *Desk) GetID() int64 {
	return int64(d.ID)
}
func (d *Desk) SetID(id int64) {
	d.ID = uint(id)
}