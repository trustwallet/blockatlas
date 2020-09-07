package models

import "time"

type (
	NotificationSubscription struct {
		DeletedAt *time.Time `gorm:"default:NULL" sql:"index"`
		Address   Address    `gorm:"ForeignKey:AddressID; not null"`
		AddressID uint       `gorm:"primary_key; auto_increment:false"`
	}

	AssetSubscription struct {
		DeletedAt *time.Time `gorm:"default:NULL" sql:"index"`
		Address   Address    `gorm:"ForeignKey:AddressID; not null"`
		AddressID uint       `gorm:"primary_key; auto_increment:false"`
	}
)
