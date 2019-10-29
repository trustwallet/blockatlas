package storage

import (
	"time"
)

type Block struct {
	Coin        interface{} `gorm:"type:varchar(20);primary_key"`
	BlockHeight int64
}

type Xpub struct {
	ID        uint        `json:"-" gorm:"primary_key"`
	Xpub      string      `json:"xpub" gorm:"type:varchar(150)"`
	Address   string      `json:"address" gorm:"type:varchar(150)"`
	Coin      interface{} `json:"coin" gorm:"type:varchar(20)"`
	CreatedAt time.Time
}

type Subscription struct {
	ID        uint        `json:"-" gorm:"primary_key"`
	Coin      interface{} `json:"coin" gorm:"type:varchar(20)"`
	Address   string      `json:"address" gorm:"type:varchar(150)"`
	Webhook   string      `json:"webhook" gorm:"type:varchar(150)"`
	IsXpub    bool        `json:"is_xpub"`
	CreatedAt time.Time
}

func (s *Subscription) Equal(sub Subscription) bool {
	if s.Coin != sub.Coin {
		return false
	}
	if s.Address != sub.Address {
		return false
	}
	if s.Webhook != sub.Webhook {
		return false
	}
	return true
}
