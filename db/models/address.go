package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	Address string `gorm:"type:varchar(128); unique_index"`
}
