package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email,omitempty" gorm:"Column:email"`
	Password string `json:"password,omitempty" gorm:"Column:password"`
}

func (User) TableName() string {
	return "user"
}
