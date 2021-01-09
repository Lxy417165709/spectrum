package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
)

type Desk struct {
	gorm.Model
	SpaceName      string           `json:"space_name"`
	SpaceNum       int              `json:"space_num"`
	Price          float64          `json:"price"`
	PriceRuleType  pb.PriceRuleType `json:"price_rule_type"`
	StartTimestamp int64            `json:"start_timestamp"`
	EndTimestamp   int64            `json:"end_timestamp"`
}

func (*Desk) TableName() string {
	return "desk"
}

type DeskGoodRecord struct {
	gorm.Model
	DeskID          int    `json:"desk_id"`
	GoodID          int    `json:"good_id"`
	MainElementName string `json:"main_element_name"`
}

func (*DeskGoodRecord) TableName() string {
	return "desk_good_record"
}
