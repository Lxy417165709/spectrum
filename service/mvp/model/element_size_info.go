package model

import (
	"spectrum/common/pb"
	"spectrum/service/mvp/utils"
	"time"
)

type ElementSizeInfo struct {
	ID               int64     `gorm:"id"`
	CreatedAt        time.Time `gorm:"created_at"`
	UpdatedAt        time.Time `gorm:"updated_at"`
	ElementID        int64     `gorm:"element_id"`
	Size             string    `gorm:"size"`
	Price            float64   `gorm:"price"`
	PictureStorePath string    `gorm:"picture_store_path"`
}

func (ElementSizeInfo) TableName() string {
	return "element_size_info"
}

func (r *ElementSizeInfo) ToPb() *pb.SizeInfo {
	return &pb.SizeInfo{
		Id:               r.ID,
		Size:             r.Size,
		Price:            utils.GetPbPrice(r.Price),
		PictureStorePath: r.PictureStorePath,
	}
}
