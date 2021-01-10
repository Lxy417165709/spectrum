package model

import (
	"github.com/jinzhu/gorm"
)

// todo: discount 还未添加
type Desk struct {
	gorm.Model
	StartTimestamp int64 `json:"start_timestamp"`
	EndTimestamp   int64 `json:"end_timestamp"`

	SpaceName string `json:"space_name"`
	SpaceNum  int    `json:"space_num"`

	Expense     float64 `json:"expense"`
	HadCheckOut bool    `json:"had_check_out"`
}

func (*Desk) TableName() string {
	return "desk"
}
