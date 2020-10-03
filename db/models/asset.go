package models

type Asset struct {
	ID    uint   `gorm:"primary_key"`
	Asset string `gorm:"type:varchar(128); uniqueIndex"`

	TokenTypeID uint      `gorm:"primary_key; autoIncrement:false; index:,"`
	TokenType   TokenType `gorm:"ForeignKey:TokenTypeID; not null"`
}
