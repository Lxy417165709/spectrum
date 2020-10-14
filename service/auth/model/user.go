package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserType = int

type User struct {
	gorm.Model
	Email             string    `json:"email,omitempty"`
	HashSaltyPassword string    `json:"hash_salty_password,omitempty"`
	Type              UserType  `json:"type,omitempty"`
	LastLoginTime     time.Time `json:"last_login_time,omitempty"`
	Salt              string    `json:"salt,omitempty"`
	AvatarPath        string    `json:"avatar_path,omitempty"`
	Username          string    `json:"username,omitempty"`
	Sex               int       `json:"sex,omitempty"`
	ContactPhone      string    `json:"contact_phone,omitempty"`
	ContactEmail      string    `json:"contact_email,omitempty"`
	Birthday          time.Time `json:"birthday,omitempty"`
}

func (User) TableName() string {
	return "user"
}
