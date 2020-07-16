package models

import "time"

type Tracker struct {
	UpdatedAt time.Time
	Coin      string `gorm:"primary_key:true; type:varchar(64)"`
	Height    int64
}
