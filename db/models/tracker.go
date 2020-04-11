package models

type Tracker struct {
	Coin   string `gorm:"primary_key:true; type:varchar(64)"`
	Height int64
}
