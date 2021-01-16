package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
	"strings"
)

type FavorRecord struct {
	gorm.Model
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
