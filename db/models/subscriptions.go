package models

import (
	"time"
)

type (
	Subscription struct {
		ID      uint   `gorm:"primaryKey;"`
		Address string `gorm:"uniqueIndex; type:varchar(256); not null;"`
	}

	SubscriptionsAssetAssociation struct {
		CreatedAt      time.Time    `gorm:"index;"`
		UpdatedAt      time.Time    `gorm:"index;"`
		Subscription   Subscription `gorm:"ForeignKey:SubscriptionId; not null"`
		SubscriptionId uint         `gorm:"primary_key; autoIncrement:false; index"`

		Asset   Asset `gorm:"ForeignKey:AssetId; not null"`
		AssetId uint  `gorm:"primary_key; autoIncrement:false; index"`
	}
)
