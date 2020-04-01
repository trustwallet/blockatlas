package models

type Subscription struct {
	SubscriptionId uint               `gorm:"primary_key:true"`
	Data           []SubscriptionData `gorm:"foreignkey:SubscriptionId"`
}

type SubscriptionData struct {
	ID             uint   `gorm:"primary_key:true"`
	SubscriptionId uint   `sql:"index"`
	Coin           uint   `sql:"index"`
	Address        string `sql:"index"`
}
