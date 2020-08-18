package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type (
	Address struct {
		gorm.Model
		Address string `gorm:"type:varchar(128); unique_index"`
	}

	Asset struct {
		gorm.Model
		AssetID string `gorm:"type:varchar(128); unique_index"`
	}

	AddressToTokenAssociation struct {
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time `sql:"index"`

		Address   Address `gorm:"ForeignKey:AddressID; not null"`
		AddressID uint    `sql:"index"`

		Asset   Asset `gorm:"ForeignKey:AssetID; not null"`
		AssetID uint  `sql:"index"`
	}
)
