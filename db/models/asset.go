package models

type Asset struct {
	ID    uint   `gorm:"primary_key; type:int4;"`
	Asset string `gorm:"type:varchar(128); uniqueIndex"`
}
