package models

import "time"

type (
	NotificationSubscription struct {
		DeletedAt *time.Time `gorm:"default:NULL; index"`
		Address   Address    `gorm:"ForeignKey:AddressID; not null"`
		AddressID uint       `gorm:"primary_key; type:int4; autoIncrement:false"`
	}

	AssetSubscription struct {
		DeletedAt *time.Time `gorm:"default:NULL; index"`
		Address   Address    `gorm:"ForeignKey:AddressID; not null"`
		AddressID uint       `gorm:"primary_key; type:int4; autoIncrement:false"`
	}
)
