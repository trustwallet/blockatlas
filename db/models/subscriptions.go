package models

import (
	"github.com/jinzhu/gorm"
)

type (
	NotificationSubscription struct {
		gorm.Model
		Address   Address `gorm:"ForeignKey:AddressID; not null"`
		AddressID uint    `gorm:"unique" sql:"index"`
	}

	AssetSubscription struct {
		gorm.Model
		Address   Address `gorm:"ForeignKey:AddressID; not null"`
		AddressID uint    `gorm:"unique" sql:"index"`
	}
)
