package model

import "github.com/jinzhu/gorm"

type OptionClass struct {
	gorm.Model
	Name                     string `json:"name"`
	DefaultSelectOptionIndex int    `json:"default_select_option_index"`
}

func (OptionClass) TableName() string {
	return "option_class"
}

type Option struct {
	gorm.Model
	OptionClassID int    `json:"option_class_id"`
	Name          string `json:"name"`
}

func (Option) TableName() string {
	return "option"
}
