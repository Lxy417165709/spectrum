package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
)

const (
	FlagOfNotAttachGood = 0
	FlagOfAttachGood    = 1
	FlagOfNotCheckout   = 0
	FlagOfHasCheckedOut = 1
)

type Element struct {
	gorm.Model
	Name      string         `json:"name"`
	Type      pb.ElementType `json:"type"`
	ClassName string         `json:"class_name"`

	Size      string         `json:"size"`
	Price            float64 `json:"price"`
	PictureStorePath string  `json:"picture_store_path"`
}

func (Element) TableName() string {
	return "element"
}

type MainElementAttachElementRecord struct {
	gorm.Model
	MainElementName   string `json:"main_element_name"`
	GoodID            int    `json:"good_id"`

	AttachElementName string `json:"attach_element_name"`
	SelectSize        string `json:"select_size"`
}

func (MainElementAttachElementRecord) TableName() string {
	return "main_element_attach_element_record"
}

type MainElementSizeRecord struct {
	gorm.Model
	MainElementName string `json:"main_element_name"`
	GoodID          int    `json:"good_id"`

	SelectSize      string `json:"select_size"`
}


type ElementClass struct {
	gorm.Model
	Name string `json:"name"`
}

func (ElementClass) TableName() string {
	return "element_class"
}
