package models

// An AddressTokenTracker marks existence of address <=> token mappings in the database.
type AddressTokenTracker struct {
	Coin    uint   `gorm:"primary_key;auto_increment:false"`
	Address string `gorm:"primary_key"`
}

// AddressToken is a single mapping between an address and a token.
type AddressToken struct {
	Coin    uint   `gorm:"primary_key;auto_increment:false"`
	Address string `gorm:"primary_key"`
	Token   string `gorm:"primary_key"`
}
