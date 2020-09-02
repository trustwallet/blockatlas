package models

type Asset struct {
	ID      uint   `gorm:"primary_key"`
	AssetID string `gorm:"type:varchar(128); unique_index"`
}
