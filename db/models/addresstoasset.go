package models

import (
	"time"
)

type AddressToAssetAssociation struct {
	CreatedAt time.Time  `gorm:"index:,"`
	DeletedAt *time.Time `gorm:"index:,; default:NULL"`

	Address   Address `gorm:"ForeignKey:AddressID; not null"`
	AddressID uint    `gorm:"primary_key; autoIncrement:false; index:,"`

	Asset   Asset `gorm:"ForeignKey:AssetID; not null"`
	AssetID uint  `gorm:"primary_key; autoIncrement:false; index:,"`
}
