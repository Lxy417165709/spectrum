package model

import (
	"github.com/jinzhu/gorm"
	"spectrum/common/pb"
)

type Space struct {
	gorm.Model
	Name          string           `json:"name"`
	Price         float64          `json:"price"`
	PriceRuleType pb.PriceRuleType `json:"price_rule_type"` // 定时、记时
}
