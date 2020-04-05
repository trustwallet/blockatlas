package models

type Tracker struct {
	Coin   uint `gorm:"primary_key:true; auto_increment:false"`
	Height int64
}
