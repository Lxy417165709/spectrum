package model

import (
	"time"
)

type ElementSelectSizeRecord struct {
	ID               uint      `gorm:"id"`
	CreatedAt        time.Time `gorm:"created_at"`
	UpdatedAt        time.Time `gorm:"updated_at"`
	GoodID           int64     `gorm:"good_id"`
	ElementClassName string    `gorm:"element_class_name"`
	ElementName      string    `gorm:"element_name"`
	SelectSize       string    `gorm:"select_size"`
}

func (ElementSelectSizeRecord) TableName() string {
	return "element_select_size_record"
}