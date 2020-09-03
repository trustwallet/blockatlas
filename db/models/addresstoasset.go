package models

import (
	"time"
)

type AddressToAssetAssociation struct {
	CreatedAt time.Time  `sql:"index"`
	DeletedAt *time.Time `sql:"index"`

	Address   Address `gorm:"ForeignKey:AddressID; not null"`
	AddressID uint    `gorm:"primary_key; auto_increment:false" sql:"index"`

	Asset   Asset `gorm:"ForeignKey:AssetID; not null"`
	AssetID uint  `gorm:"primary_key; auto_increment:false" sql:"index"`
}
