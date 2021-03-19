package model

import (
	"time"
)

type MainElementAttachElementRecord struct {
	ID                uint      `gorm:"id"`
	CreatedAt         time.Time `gorm:"created_at"`
	UpdatedAt         time.Time `gorm:"updated_at"`
	GoodID            int64     `gorm:"good_id"`
	MainElementName   string    `gorm:"main_element_name"`
	AttachElementName string    `gorm:"attach_element_name"`
	SelectSize        string    `gorm:"select_size"`
}

func (MainElementAttachElementRecord) TableName() string {
	return "main_element_attach_element_record"
}

type MainElementSizeRecord struct {
	ID              uint      `gorm:"id"`
	CreatedAt       time.Time `gorm:"created_at"`
	UpdatedAt       time.Time `gorm:"updated_at"`
	GoodID          int64     `gorm:"good_id"`
	MainElementName string    `gorm:"main_element_name"`
	SelectSize      string    `gorm:"select_size"`
}

func (MainElementSizeRecord) TableName() string {
	return "main_element_size_record"
}

