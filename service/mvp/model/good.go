package model

import "github.com/jinzhu/gorm"

// todo: discount 还未添加
type Good struct {
	gorm.Model
	Name   string `json:"name"`
	DeskID int    `json:"desk_id"`

	Expense     float64 `json:"expense"`
	HadCheckOut bool    `json:"had_check_out"`
}

func (*Good) TableName() string {
	return "good"
}
