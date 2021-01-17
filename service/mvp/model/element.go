package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
)

type Element struct {
	gorm.Model
	Name      string         `json:"name"`
	Type      pb.ElementType `json:"type"`
	ClassName string         `json:"class_name"`

	Size             string  `json:"size"`
	Price            float64 `json:"price"`
	PictureStorePath string  `json:"picture_store_path"`
}

func (Element) TableName() string {
	return "element"
}
