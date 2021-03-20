package model

import (
	"spectrum/common/pb"
	"time"
)

type SpaceClass struct {
	ID               int64     `gorm:"id"`
	CreatedAt        time.Time `gorm:"created_at"`
	UpdatedAt        time.Time `gorm:"updated_at"`
	Name             string    `json:"name"`
	PictureStorePath string    `json:"picture_store_path"`
}

func (d *SpaceClass) TableName() string {
	return "space_class"
}

func (d *SpaceClass) ToPb() *pb.DeskClass {
	return &pb.DeskClass{
		Id:               d.ID,
		Name:             d.Name,
		PictureStorePath: d.PictureStorePath,
	}
}
