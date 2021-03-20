package model

import (
	"time"
)

type MainElementAttachElementRecord struct {
	ID                     uint      `gorm:"id"`
	CreatedAt              time.Time `gorm:"created_at"`
	UpdatedAt              time.Time `gorm:"updated_at"`
	GoodID                 int64     `gorm:"good_id"`
	AttachElementClassName string    `gorm:"attach_element_class_name"`
	MainElementClassName string    `gorm:"main_element_class_name"`
	MainElementName        string    `gorm:"main_element_name"`
	AttachElementName      string    `gorm:"attach_element_name"`
	SelectSize             string    `gorm:"select_size"`
}

func (MainElementAttachElementRecord) TableName() string {
	return "main_element_attach_element_record"
}

type ElementSizeRecord struct {
	ID          uint      `gorm:"id"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
	GoodID      int64     `gorm:"good_id"`
	ElementName string    `gorm:"element_name"`
	SelectSize  string    `gorm:"select_size"`
}

func (ElementSizeRecord) TableName() string {
	return "element_size_record"
}
