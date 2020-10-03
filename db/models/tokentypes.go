package models

import "time"

type TokenType struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"index:,"`
	Type      string    `gorm:"varchar(32); uniqueIndex"`
	Decimals  uint      `gorm:"int(4)"`
	Name      string    `gorm:"type:varchar(128)"`
	Symbol    string    `gorm:"type:varchar(128)"`
}
