package model

import "github.com/jinzhu/gorm"

type OptionClass struct{
	gorm.Model
	Name string `json:"name"`
}

func (OptionClass) TableName() string{
	return "option_class"
}

type Option struct{
	gorm.Model
	OptionClassID int `json:"option_class_id"`
	Name string `json:"name"`
}

func (Option) TableName() string{
	return "option"
}