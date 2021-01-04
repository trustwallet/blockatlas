package models

import (
	"errors"
	"time"
	"unicode/utf8"
)

type Asset struct {
	CreatedAt time.Time `gorm:"index:,"`
	ID        uint      `gorm:"primary_key; uniqueIndex"`
	Asset     string    `gorm:"type:varchar(128); uniqueIndex"`

	Decimals uint   `gorm:"int(4)"`
	Name     string `gorm:"type:varchar(128)"`
	Symbol   string `gorm:"type:varchar(128)"`
	Type     string `gorm:"type:varchar(12)"`
	Coin     uint
}

func (asset *Asset) IsValid() error {
	if len(asset.Name) >= 32 {
		return errors.New("name should be less than 32")
	}
	if len(asset.Symbol) >= 32 {
		return errors.New("name should be less than 32")
	}
	stringValues := []string{asset.Asset, asset.Type, asset.Symbol, asset.Name}

	for _, value := range stringValues {
		if value == "" {
			return errors.New("empty value for asset: " + asset.Asset)
		}
		if !utf8.ValidString(value) {
			return errors.New("not valid utf8 string: " + value)
		}
	}
	return nil
}

func AssetIDs(assets []Asset) []string {
	result := make([]string, 0, len(assets))
	for _, a := range assets {
		result = append(result, a.Asset)
	}
	return result
}
