package models

type Subscription struct {
	GUID string             `gorm:"primary_key:true"`
	Data []SubscriptionData `gorm:"many2many:subscription_associations"`
}

type SubscriptionData struct {
	ID             uint   `gorm:"primary_key:true"`
	SubscriptionId string `sql:"index"`
	Coin           uint   `sql:"index"`
	Address        string `sql:"index"`
}
