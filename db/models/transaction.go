package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type Transaction struct {
	ID         string `gorm:"primary_key; autoIncrement:false; index"`
	Coin       uint   `gorm:"primary_key; autoIncrement:false; index"`
	AssetID    string
	From       string
	To         string
	FeeAssetID string
	Fee        string
	Date       time.Time
	Block      uint64
	Sequence   uint64
	Status     string
	Memo       string
	Type       string
	Metadata   postgres.Jsonb
}
