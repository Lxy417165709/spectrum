package model

import "github.com/jinzhu/gorm"

type Good struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (Good) TableName() string {
	return "good"
}

type SellRecord struct {
	gorm.Model
	GoodID    int     `json:"good_id"`
	SellPrice float64 `json:"sell_price"`
}

func (SellRecord) TableName() string {
	return "sell_record"
}
