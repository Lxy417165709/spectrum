package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	FlagOfNotAttachGood = 0
	FlagOfAttachGood    = 1
)


type GoodType struct{
	gorm.Model
	Name string `json:"name"`
}

func (GoodType ) TableName() string {
	return "good_type"
}


type Good struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Type int64 `json:"type"`
}

func (Good) TableName() string {
	return "good"
}

type SellRecord struct {
	gorm.Model
	GoodID    int     `json:"good_id"`
	SellPrice float64 `json:"sell_price"`
}

func (SellRecord) TableName() string {
	return "sell_record"
}

type BilliardDesk struct {
	gorm.Model
	Name string `json:"name"`
}

func (BilliardDesk) TableName() string {
	return "billiard_desk"
}

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
