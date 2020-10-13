package models

type TokenType struct {
	ID   uint   `gorm:"primary_key"`
	Type string `gorm:"varchar(32); uniqueIndex"`
}
