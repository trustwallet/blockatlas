package models

import (
	"time"
)

type Subscription struct {
	CreatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Coin      uint       `gorm:"primary_key; column:coin; auto_increment:false"`
	Address   string     `gorm:"primary_key; column:address; type:varchar(128)"`
}
