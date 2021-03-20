package model

import (
	"time"
)

type CheckOutRecord struct {
	ID              int64     `gorm:"id"`
	CreatedAt       time.Time `gorm:"created_at"`
	UpdatedAt       time.Time `gorm:"updated_at"`
	ChargeableObjectName string `json:"chargeable_object_name"`
	ChargeableObjectID   int64  `json:"chargeable_object_id"`
	CheckOutAt    int64  `json:"check_out_at"`
}

func (c *CheckOutRecord) TableName() string {
	return "check_out_record"
}
