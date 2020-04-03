package models

import "time"

type Subscription struct {
	SubscriptionId uint `gorm:"primary_key:true"`
	UpdatedAt      time.Time
	Data           []SubscriptionData `gorm:"foreignkey:SubscriptionId"`
}

type SubscriptionData struct {
	ID             uint   `gorm:"primary_key;"`
	SubscriptionId uint   `gorm:"primary_key; column:subscription_id; auto_increment:false"`
	Coin           uint   `gorm:"primary_key; column:coin; auto_increment:false"`
	Address        string `gorm:"primary_key; column:address; type:varchar(128)"`
}
