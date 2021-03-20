package model

import (
	"spectrum/common/pb"
	"time"
)

type ElementSizeInfoRecord struct {
	ID               int64     `gorm:"id"`
	CreatedAt        time.Time `gorm:"created_at"`
	UpdatedAt        time.Time `gorm:"updated_at"`
	ElementID        int64     `gorm:"element_id"`
	Size             string    `gorm:"size"`
	Price            float64   `gorm:"price"`
	PictureStorePath string    `gorm:"picture_store_path"`
}

func (ElementSizeInfoRecord) TableName() string {
	return "element_size_info_record"
}

func (r *ElementSizeInfoRecord) ToPb() *pb.SizeInfo {
	return &pb.SizeInfo{
		Id:               r.ID,
		Size:             r.Size,
		Price:            GetPbPrice(r.Price),
		PictureStorePath: r.PictureStorePath,
	}
}
