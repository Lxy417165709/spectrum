package model

import "github.com/jinzhu/gorm"

type MainElementAttachElementRecord struct {
	gorm.Model
	MainElementName string `json:"main_element_name"`
	GoodID          int64  `json:"good_id"`

	AttachElementName string `json:"attach_element_name"`
	SelectSize        string `json:"select_size"`
}

func (MainElementAttachElementRecord) TableName() string {
	return "main_element_attach_element_record"
}

type MainElementSizeRecord struct {
	gorm.Model
	GoodID          int64  `json:"good_id"`
	MainElementName string `json:"main_element_name"`

	SelectSize string `json:"select_size"`
}

type ElementClass struct {
	gorm.Model
	Name string `json:"name"`
	PictureStorePath string `json:"picture_store_path"`
}

func (ElementClass) TableName() string {
	return "element_class"
}
