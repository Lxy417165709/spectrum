package model

import (
	"spectrum/common/pb"
	"time"
)

type GoodClass struct {
	ID               uint      `gorm:"id"`
	CreatedAt        time.Time `gorm:"created_at"`
	UpdatedAt        time.Time `gorm:"updated_at"`
	Name             string    `gorm:"name"`
	PictureStorePath string    `gorm:"picture_store_path"`
}

func (GoodClass) TableName() string {
	return "good_class"
}

func (g *GoodClass) ToPb() *pb.GoodClass {
	return &pb.GoodClass{
		Id:               int64(g.ID),
		Name:             g.Name,
		PictureStorePath: g.PictureStorePath,
	}
}