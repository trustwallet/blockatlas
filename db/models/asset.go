package models

type Asset struct {
	ID    uint   `gorm:"primary_key; uniqueIndex"`
	Asset string `gorm:"type:varchar(128); uniqueIndex"`
	//Decimals uint   `gorm:"int(4)"`
	//Name     string `gorm:"type:varchar(128)"`
	//Symbol   string `gorm:"type:varchar(128)"`
	//
	//TokenTypeID uint      `gorm:"primary_key; autoIncrement:false; index:,"`
	//TokenType   TokenType `gorm:"foreignKey:TokenTypeID; not null"`
}
