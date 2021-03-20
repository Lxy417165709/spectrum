package model

import (
	"spectrum/common/pb"
	"time"
)

type Element struct {
	ID        uint      `gorm:"id"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`

	ClassName string         `gorm:"class_name"`
	Name      string         `gorm:"name"`
	Type      pb.ElementType `gorm:"type"`
}

func (Element) TableName() string {
	return "element"
}

type ElementSizeInfoRecord struct {
	ID               uint      `gorm:"id"`
	CreatedAt        time.Time `gorm:"created_at"`
	UpdatedAt        time.Time `gorm:"updated_at"`
	GoodID           int64     `gorm:"good_id"`
	ClassName        string    `gorm:"class_name"`
	Name             string    `gorm:"name"`
	Size             string    `gorm:"size"`
	Price            float64   `gorm:"price"`
	PictureStorePath string    `gorm:"picture_store_path"`
}

func (ElementSizeInfoRecord) TableName() string {
	return "element_size_info_record"
}

func (r *ElementSizeInfoRecord) ToPb() *pb.SizeInfo {
	return &pb.SizeInfo{
		Id:               int64(r.ID),
		Size:             r.Size,
		Price:            GetPbPrice(r.Price),
		PictureStorePath: r.PictureStorePath,
	}
}
