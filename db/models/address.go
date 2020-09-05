package models

type Address struct {
	ID      uint   `gorm:"primary_key"`
	Address string `gorm:"type:varchar(128); unique_index"`
}

// Use such model in future
// Coin    uint   `gorm:"index:idx_coin;" sql:"unique_index:idx_ca"`
// Address string `gorm:"index:idx_address; type:varchar(128)" sql:"unique_index:idx_ca"`
