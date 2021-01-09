package model

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	DeskID int
}

func (Order) TableName() string {
	return "order"
}
