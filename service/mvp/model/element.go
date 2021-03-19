package model

import (
	"spectrum/common/pb"
	"time"
)

type Element struct {
	ID               uint           `gorm:"id"`
	CreatedAt        time.Time      `gorm:"created_at"`
	UpdatedAt        time.Time      `gorm:"updated_at"`

	Name             string         `gorm:"name"`
	Type             pb.ElementType `gorm:"type"`
	ClassName        string         `gorm:"class_name"`
	Size             string         `gorm:"size"`
	Price            float64        `gorm:"price"`
	PictureStorePath string         `gorm:"picture_store_path"`
}

func (Element) TableName() string {
	return "element"
}
