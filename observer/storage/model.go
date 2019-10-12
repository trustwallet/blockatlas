package storage

import "time"

type Block struct {
	ID          uint `gorm:"primary_key"`
	Coin        uint `gorm:"primary_key"`
	BlockHeight int64
}

type Xpub struct {
	ID        uint `gorm:"auto_increment;not null"`
	Coin      uint
	Address   string `gorm:"primary_key;type:varchar(150)"`
	Xpub      string `gorm:"primary_key;type:varchar(150)"`
	CreatedAt time.Time
}

type Subscription struct {
	ID        uint   `json:"-" gorm:"auto_increment;not null"`
	Coin      uint   `json:"coin"`
	Address   string `json:"address" gorm:"primary_key;type:varchar(150)"`
	Webhook   string `json:"webhook" gorm:"primary_key;type:varchar(150)"`
	Xpub      string `json:"-" sql:"-" gorm:"-"`
	CreatedAt time.Time
}
