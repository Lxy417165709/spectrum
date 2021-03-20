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

	SessionCount int64  `json:"session_count"`
	SpaceName    string `json:"space_name"`
	SpaceClassName   string `json:"space_class_name"`

	Expense           float64 `json:"expense"`
	CheckOutAt time.Time   `gorm:"check_out_at"`
	NonFavorExpense   float64 `json:"non_favor_expense"`

	OrderID int64 `json:"order_id"`
}

func (d *Desk) TableName() string {
	return d.GetName()
}

func (*Desk) GetName() string {
	return ChargeableObjectNameOfDesk
}

func (d *Desk) GetExpenseInfo(billingType pb.BillingType, price float64, favors []*pb.Favor) *pb.ExpenseInfo {
	if d.CheckOutAt.Unix() != 0 {
		return &pb.ExpenseInfo{
			NonFavorExpense:   d.NonFavorExpense,
			CheckOutAt: d.CheckOutAt.Unix(),
			Expense:           d.Expense,
		}
	}

	var nonFavorExpense float64
	switch billingType {
	case pb.BillingType_Timing:
		nonFavorExpense = d.getTimingNonFavorExpense(price)
	case pb.BillingType_Session:
		nonFavorExpense = float64(d.SessionCount) * price
	}

	return &pb.ExpenseInfo{
		NonFavorExpense:   nonFavorExpense,
		CheckOutAt: 0,
		Expense:           GetFavorExpense(nonFavorExpense, favors),
	}
}

func (d *Desk) getTimingNonFavorExpense(pricePerHour float64) float64 {
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

type DeskClass struct {
	gorm.Model
	Name             string `json:"name"`
	PictureStorePath string `json:"picture_store_path"`
}

func (d *DeskClass) TableName() string {
	return "desk_class"
}

func (d *DeskClass) ToPb() *pb.DeskClass {
	return &pb.DeskClass{
		Name:             d.Name,
		PictureStorePath: d.PictureStorePath,
	}
}
