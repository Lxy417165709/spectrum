package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
)

type FavorRecord struct {
	gorm.Model
	FavorableStructName string       `json:"favorable_struct_name"`
	FavorableStructID   int64        `json:"favorable_struct_id"`
	FavorType           pb.FavorType `json:"favor_type"`
	FavorParameters     string       `json:"favor_parameters"`
}

func (f *FavorRecord) TableName() string {
	return "favor_record"
}
