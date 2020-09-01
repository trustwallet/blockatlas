package models

import (
	"time"
)

type AddressToAssetAssociation struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Address   Address `gorm:"ForeignKey:AddressID; not null"`
	AddressID uint    `sql:"index"`

	Asset   Asset `gorm:"ForeignKey:AssetID; not null"`
	AssetID uint  `sql:"index"`
}
