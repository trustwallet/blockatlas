package storage

import (
	"database/sql"
	"time"
)

type Block struct {
	Coin        interface{} `gorm:"type:varchar(20);primary_key"`
	BlockHeight int64
}

type Subscription struct {
	Coin      interface{}    `json:"coin" gorm:"type:varchar(20);primary_key"`
	Address   string         `json:"address" gorm:"type:varchar(150);primary_key"`
	Webhook   string         `json:"webhook" gorm:"type:varchar(150);primary_key"`
	Xpub      sql.NullString `json:"-" gorm:"type:varchar(150)"`
	CreatedAt time.Time
}

func (s *Subscription) Equal(sub Subscription) bool {
	if s.Coin != sub.Coin {
		return false
	}
	if s.Address != sub.Address {
		return false
	}
	if s.Xpub.String != sub.Xpub.String {
		return false
	}
	if s.Webhook != sub.Webhook {
		return false
	}
	return true
}
