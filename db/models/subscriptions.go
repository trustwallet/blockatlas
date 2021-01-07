package models

import (
	"time"
)

type (
	// Subscription for address and asset associations
	Subscription struct {
		ID        uint `gorm:"primaryKey;"`
		CreatedAt time.Time
		Address   string `gorm:"uniqueIndex:idx_address; type:varchar(256); not null;"`
	}

	SubscriptionsAssetAssociation struct {
		CreatedAt      time.Time    `gorm:"index;"`
		Subscription   Subscription `gorm:"ForeignKey:SubscriptionId; not null"`
		SubscriptionId uint         `gorm:"primary_key; autoIncrement:false; index"`

		Asset   Asset `gorm:"ForeignKey:AssetId; not null"`
		AssetId uint  `gorm:"primary_key; autoIncrement:false; index"`
	}
)
