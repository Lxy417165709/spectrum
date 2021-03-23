package model

import (
	"spectrum/common/pb"
	"time"
)

type Desk struct {
	ID        int64     `gorm:"id"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	StartAt   time.Time `json:"start_at"`
	EndAt     time.Time `json:"end_at"`

	SessionCount int64 `json:"session_count"`
	SpaceID      int64 `json:"space_id"`

	Expense         float64   `json:"expense"`
	CheckOutAt      time.Time `gorm:"check_out_at"`
	NonFavorExpense float64   `json:"non_favor_expense"`

	OrderID int64 `json:"order_id"`
}

func (d *Desk) TableName() string {
	return d.GetName()
}

func (*Desk) GetName() string {
	return ChargeableObjectNameOfDesk
}

func (d *Desk) GetExpenseInfo(billingType pb.BillingType, price float64, favors []*pb.Favor) *pb.ExpenseInfo {
	// 1. 已结账时
	if d.CheckOutAt != NilTime {
		return &pb.ExpenseInfo{
			NonFavorExpense: d.NonFavorExpense,
			CheckOutAt:      d.CheckOutAt.Unix(),
			Expense:         d.Expense,
		}
	}

	// 2. 未结账时
	var nonFavorExpense float64
	switch billingType {
	case pb.BillingType_Timing:
		nonFavorExpense = d.getTimingNonFavorExpense(price)
	case pb.BillingType_Session:
		nonFavorExpense = float64(d.SessionCount) * price
	}
	return &pb.ExpenseInfo{
		NonFavorExpense: nonFavorExpense,
		CheckOutAt:      NilTime.Unix(),
		Expense:         GetFavorExpense(nonFavorExpense, favors),
	}
}

func (d *Desk) getTimingNonFavorExpense(pricePerHour float64) float64 {
	var endTimestamp int64
	if d.IsOpening() {
		endTimestamp = time.Now().Unix()
	} else {
		endTimestamp = d.EndAt.Unix()
	}
	hours := time.Unix(endTimestamp, 0).Sub(time.Unix(d.StartAt.Unix(), 0)).Hours()
	return hours * pricePerHour
}

func (d *Desk) IsOpening() bool {
	return d.EndAt == NilTime
}

func (d *Desk) GetID() int64 {
	return d.ID
}
