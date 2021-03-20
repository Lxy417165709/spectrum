package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserType = int

type User struct {
	ID              int64     `gorm:"id"`
	CreatedAt       time.Time `gorm:"created_at"`
	UpdatedAt       time.Time `gorm:"updated_at"`
	Email             string    `json:"email"`
	HashSaltyPassword string    `json:"hash_salty_password"`
	Type              UserType  `json:"type"`
	LastLoginTime     time.Time `json:"last_login_time"`
	Salt              string    `json:"salt"`
	AvatarPath        string    `json:"avatar_path"`
	Username          string    `json:"username"`
	Sex               int       `json:"sex"`
	ContactPhone      string    `json:"contact_phone"`
	ContactEmail      string    `json:"contact_email"`
	Birthday          time.Time `json:"birthday"`
}

func (User) TableName() string {
	return "user"
}
