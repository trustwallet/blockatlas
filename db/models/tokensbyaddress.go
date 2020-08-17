package models

import "github.com/jinzhu/gorm"

type (
	Address struct {
		gorm.Model
		Address string `gorm:"type:varchar(128); primary_key; unique_index"`
	}

	Asset struct {
		gorm.Model
		AssetID string `gorm:"type:varchar(128); primary_key; unique_index"`
	}

	AddressToTokenAssociation struct {
		gorm.Model

		Address   Address `gorm:"ForeignKey:AddressID; not null"`
		AddressID uint    `sql:"index"`

		Asset   Asset `gorm:"ForeignKey:AssetID; not null"`
		AssetID uint  `sql:"index"`
	}
)
