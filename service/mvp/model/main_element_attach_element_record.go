package model

import "time"

type MainElementAttachElementRecord struct {
	ID                int64     `gorm:"id"`
	CreatedAt         time.Time `gorm:"created_at"`
	UpdatedAt         time.Time `gorm:"updated_at"`
	GoodID            int64     `gorm:"good_id"`
	MainElementID     int64     `gorm:"main_element_id"`
	AttachElementID   int64     `gorm:"attach_element_id"`
	SelectSizeInfoID int64     `gorm:"select_size_info_id"`
}

func (MainElementAttachElementRecord) TableName() string {
	return "main_element_attach_element_record"
}
