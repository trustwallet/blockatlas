package models

import (
	"time"
)

type AddressToAssetAssociation struct {
	CreatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Address   Address `gorm:"ForeignKey:AddressID; not null"`
	AddressID uint    `gorm:"index:idx_address" sql:"unique_index:idx_aa"`

	Asset   Asset `gorm:"ForeignKey:AssetID; not null"`
	AssetID uint  `gorm:"index:idx_asset" sql:"unique_index:idx_aa"`
}
