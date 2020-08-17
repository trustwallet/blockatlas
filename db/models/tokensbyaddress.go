package models

type (
	AddressToTokenAssociation struct {
		Address          string `gorm:"type:varchar(128)"  sql:"index"`
		TokenID          uint   `sql:"index"`
		Token            `gorm:"ForeignKey:TokenID; not null"`
		LastUpdatedBlock uint
	}
)
