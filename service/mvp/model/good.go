package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Good struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
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
