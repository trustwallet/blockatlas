package models

import (
	"github.com/jinzhu/gorm"
)

type Notification struct {
	gorm.Model
	Address   Address `gorm:"ForeignKey:AddressID; not null"`
	AddressID uint    `sql:"index"`
}
