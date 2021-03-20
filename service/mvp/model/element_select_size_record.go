package model

import (
	"time"
)

type ElementSelectSizeRecord struct {
	ID               int64     `gorm:"id"`
	CreatedAt        time.Time `gorm:"created_at"`
	UpdatedAt        time.Time `gorm:"updated_at"`
	GoodID           int64     `gorm:"good_id"`
	ElementID        int64     `gorm:"element_Id"`
	SelectSizeInfoID int64     `gorm:"select_size_info_id"`
}

func (ElementSelectSizeRecord) TableName() string {
	return "element_select_size_record"
}
