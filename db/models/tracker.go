package models

type Tracker struct {
	Coin   string `gorm:"primary_key:true;"`
	Height int64
}
