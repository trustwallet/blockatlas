package models

type Asset struct {
	ID    uint   `gorm:"primary_key"`
	Asset string `gorm:"type:varchar(128); unique_index"`
}
