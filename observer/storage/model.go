package storage

import "time"

type Block struct {
	Coin        interface{} `gorm:"type:varchar(20);primary_key"`
	BlockHeight int64
}

type Xpub struct {
	Coin      interface{} `gorm:"primary_key;type:varchar(20)"`
	Address   string      `gorm:"primary_key;type:varchar(150)"`
	Xpub      string      `gorm:"primary_key;type:varchar(150)"`
	CreatedAt time.Time
}

type Subscription struct {
	Coin      interface{} `json:"coin" gorm:"primary_key;type:varchar(20)"`
	Address   string      `json:"address" gorm:"primary_key;type:varchar(150)"`
	Webhook   string      `json:"webhook" gorm:"primary_key;type:varchar(150)"`
	Xpub      string      `json:"-" sql:"-" gorm:"-"`
	CreatedAt time.Time
}
