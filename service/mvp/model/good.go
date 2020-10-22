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
	Name      string `json:"name"`
	ClassType int    `json:"class_type"`
}

func (GoodClass) TableName() string {
	return "good_class"
}


type GoodAttachClassRecord struct {
	gorm.Model
	GoodID        int `json:"good_id"`
	AttachGoodClassID int `json:"attach_good_class_id"`
}

func (GoodAttachClassRecord) TableName() string {
	return "good_attach_class_record"
}