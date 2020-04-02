package models

type Tracker struct {
	Coin   uint `gorm:"primary_key:true"`
	Height int64
}
