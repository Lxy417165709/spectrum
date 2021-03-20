package model

import (
	"spectrum/common/pb"
	"time"
)

type Element struct {
	ID        int64      `gorm:"id"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`

	ClassName string         `gorm:"class_name"`
	Name      string         `gorm:"name"`
	Type      pb.ElementType `gorm:"type"`
}

func (Element) TableName() string {
	return "element"
}

