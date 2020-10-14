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
