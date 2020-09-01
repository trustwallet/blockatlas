package models

import (
	"time"
)

type AddressToAssetAssociation struct {
	CreatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Address   Address `gorm:"ForeignKey:AddressID; not null"`
	AddressID uint    `gorm:"primary_key"`

	Asset   Asset `gorm:"ForeignKey:AssetID; not null"`
	AssetID uint  `gorm:"primary_key"`
}
