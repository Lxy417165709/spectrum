package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
)

// todo: Favor 还未添加
type Desk struct {
	gorm.Model
	StartTimestamp int64 `json:"start_timestamp"`
	EndTimestamp   int64 `json:"end_timestamp"`

	SpaceName string `json:"space_name"`
	SpaceNum  int    `json:"space_num"`

	Expense     float64 `json:"expense"`
	HadCheckOut bool    `json:"had_check_out"`

	// 未使用字段
	FavorType      pb.FavorType `json:"favor_type"`
	FavorParameter string       `json:"favor_parameter"` // 这里将参数浓缩为字符串
}

func (*Desk) TableName() string {
	return "desk"
}
