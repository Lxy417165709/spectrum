package model

import (
	"spectrum/common/pb"
	"strings"
	"time"
)

type FavorRecord struct {
	ID                   uint         `gorm:"id"`
	CreatedAt            time.Time    `gorm:"created_at"`
	UpdatedAt            time.Time    `gorm:"updated_at"`
	ChargeableObjectName string       `json:"chargeable_object_name"`
	ChargeableObjectID   int64        `json:"chargeable_object_id"`
	FavorType            pb.FavorType `json:"favor_type"`
	FavorParameters      string       `json:"favor_parameters"`
}

func (f *FavorRecord) TableName() string {
	return "favor_record"
}

func (f *FavorRecord) ToPb() *pb.Favor {
	return &pb.Favor{
		FavorType:  f.FavorType,
		Parameters: strings.Split(f.FavorParameters, "|"),
	}
}
