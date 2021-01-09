package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type PlayRecord struct {
	gorm.Model
	BilliardDeskID int       `json:"billiard_desk_id"`
	BeginPlayTime  time.Time `json:"begin_play_time"`
	StopPlayTime   time.Time `json:"stop_play_time"`
}

func (PlayRecord) TableName() string {
	return "play_record"
}

type OrderRecord struct {
	gorm.Model
	OrderID       int `json:"order_id"`
	GoodID        int `json:"good_id"`
	HasCheckedOut int `json:"has_checked_out"` // 0: 未结账 1: 已结账
	IsAttachGood  int `json:"is_attach_good"`  // 0: 否    1: 是
}

func (OrderRecord) TableName() string {
	return "order_record"
}

type AttachRecord struct {
	gorm.Model
	OrderID      int `json:"order_id"`
	ThingID      int `json:"thing_id"`
	GoodID       int `json:"good_id"`
	AttachGoodID int `json:"attach_good_id"`
}

func (AttachRecord) TableName() string {
	return "attach_record"
}

type SellRecord struct {
	gorm.Model
	GoodID    int     `json:"good_id"`
	SellPrice float64 `json:"sell_price"`
}

func (SellRecord) TableName() string {
	return "sell_record"
}

type Thing struct {
	gorm.Model // ID需额外生成
	Price      float64
	GoodID     int
}

func (Thing ) TableName() string {
	return "thing"
}




type OrderThingRecord struct {
	gorm.Model
	OrderID int
	ThingID int
}
func (OrderThingRecord) TableName() string {
	return "order_thing_record"
}

type ThingOptionRecord struct {
	gorm.Model
	ThingID  int
	OptionID int
}

func (ThingOptionRecord) TableName() string {
	return "thing_option_record"
}

type ThingAttachGoodRecord struct {
	gorm.Model
	ThingID      int
	AttachGoodID int
}

func (ThingAttachGoodRecord) TableName() string {
	return "thing_attach_good_record"
}

//type GoodAttachClassRecord struct {
//	gorm.Model
//	GoodID        int `json:"good_id"`
//	AttachGoodClassID int `json:"attach_good_class_id"`
//}
//
//func (GoodAttachClassRecord) TableName() string {
//	return "good_attach_class_record"
//}