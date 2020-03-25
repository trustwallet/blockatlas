package models

import (
	"time"
)

type TimeModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Subscription struct {
	TimeModel
	SubscriptionId uint               `gorm:"primary_key:true"`
	Data           []SubscriptionData `gorm:"many2many:subscription_associations"`
}

type SubscriptionData struct {
	TimeModel
	ID             uint   `gorm:"primary_key:true"`
	SubscriptionId uint   `sql:"index"`
	Coin           uint   `sql:"index"`
	Address        string `sql:"index"`
}
