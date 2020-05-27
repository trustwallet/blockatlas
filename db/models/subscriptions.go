package models

import (
	"time"
)

type Subscription struct {
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index"`
	Coin      uint       `gorm:"primary_key; column:coin; auto_increment:false";sql:"index"`
	Address   string     `gorm:"primary_key; column:address; type:varchar(128)";sql:"index"`
}
