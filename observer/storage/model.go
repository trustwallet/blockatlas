package storage

import "time"

type Block struct {
	Coin        int `gorm:"primary_key" sql:"TYPE:int not null"`
	BlockHeight int64
}

type Xpub struct {
	ID        int `gorm:"auto_increment;not null"`
	Coin      int
	Address   string `gorm:"primary_key;type:varchar(150)"`
	Xpub      string `gorm:"primary_key;type:varchar(150)"`
	CreatedAt time.Time
}

type Subscription struct {
	ID        int    `json:"-" gorm:"auto_increment;not null"`
	Coin      int    `json:"coin"`
	Address   string `json:"address" gorm:"primary_key;type:varchar(150)"`
	Webhook   string `json:"webhook" gorm:"primary_key;type:varchar(150)"`
	Xpub      string `json:"-" sql:"-" gorm:"-"`
	CreatedAt time.Time
}
