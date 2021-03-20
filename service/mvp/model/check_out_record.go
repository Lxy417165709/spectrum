package model

import "github.com/jinzhu/gorm"

type CheckOutRecord struct {
	gorm.Model
	ChargeableObjectName string `json:"chargeable_object_name"`
	ChargeableObjectID   int64  `json:"chargeable_object_id"`
	CheckOutAt    int64  `json:"check_out_timestamp"`
}

func (c *CheckOutRecord) TableName() string {
	return "check_out_record"
}
