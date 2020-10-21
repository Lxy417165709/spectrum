package model

import (
	"github.com/jinzhu/gorm"
)

const (
	FlagOfNotAttachGood = 0
	FlagOfAttachGood    = 1
	FlagOfNotCheckout   = 0
	FlagOfHasCheckedOut = 1
)

type GoodType struct {
	gorm.Model
	Name string `json:"name"`
}

func (GoodType) TableName() string {
	return "good_type"
}

type Good struct {
	gorm.Model
	Name             string  `json:"name"`
	Price            float64 `json:"price"`
	Type             int     `json:"type"`
	PictureStorePath string  `json:"picture_store_path"`
	ClassID          int     `json:"class_id"`
}

func (Good) TableName() string {
	return "good"
}

type BilliardDesk struct {
	gorm.Model
	Name string `json:"name"`
}

func (BilliardDesk) TableName() string {
	return "billiard_desk"
}

type GoodOptionClassRecord struct {
	gorm.Model
	GoodID        int `json:"good_id"`
	OptionClassID int `json:"option_class_id"`
}

func (GoodOptionClassRecord) TableName() string {
	return "good_option_class_record"
}

type GoodClass struct {
	gorm.Model
	Name string `json:"name"`
}

func (GoodClass) TableName() string {
	return "good_class"
}
