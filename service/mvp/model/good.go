package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
)

// todo: Favor 还未添加
type Good struct {
	gorm.Model
	Name   string `json:"name"`
	DeskID int64  `json:"desk_id"`

	Expense     float64 `json:"expense"`
	HadCheckOut bool    `json:"had_check_out"`

	// 未使用字段
	FavorType      pb.FavorType `json:"favor_type"`
	FavorParameter string       `json:"favor_parameter"` // 这里将参数浓缩为字符串
}

func (*Good) TableName() string {
	return "good"
}
