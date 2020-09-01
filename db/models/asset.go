package models

import "github.com/jinzhu/gorm"

type Asset struct {
	gorm.Model
	AssetID string `gorm:"type:varchar(128); unique_index"`
}
