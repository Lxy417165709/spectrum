package model

import (
	"spectrum/common/pb"
	"time"
)

type ElementClass struct {
	ID               uint           `gorm:"id"`
	CreatedAt        time.Time      `gorm:"created_at"`
	UpdatedAt        time.Time      `gorm:"updated_at"`
	Name             string         `gorm:"name"`
	ClassType        pb.ElementType `gorm:"class_type"`
	PictureStorePath string         `gorm:"picture_store_path"`
}

func (ElementClass) TableName() string {
	return "element_class"
}

func (g *ElementClass) ToPb() *pb.GoodClass {
	return &pb.GoodClass{
		Id:               int64(g.ID),
		Name:             g.Name,
		PictureStorePath: g.PictureStorePath,
	}
}
