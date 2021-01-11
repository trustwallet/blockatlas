package models

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/trustwallet/golibs/asset"
	"github.com/trustwallet/golibs/txtype"
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

func AssetFrom(t txtype.Tx) (Asset, bool) {
	var a Asset
	switch t.Meta.(type) {
	case txtype.TokenTransfer:
		a.Asset = asset.BuildID(t.Coin, t.Meta.(txtype.TokenTransfer).TokenID)
		a.Decimals = t.Meta.(txtype.TokenTransfer).Decimals
		a.Name = t.Meta.(txtype.TokenTransfer).Name
		a.Symbol = t.Meta.(txtype.TokenTransfer).Symbol
		tp, ok := txtype.GetTokenType(t.Coin, t.Meta.(txtype.TokenTransfer).TokenID)
		if !ok {
			return Asset{}, false
		}
		a.Type = tp
	case *txtype.TokenTransfer:
		a.Asset = asset.BuildID(t.Coin, t.Meta.(*txtype.TokenTransfer).TokenID)
		a.Decimals = t.Meta.(*txtype.TokenTransfer).Decimals
		a.Name = t.Meta.(*txtype.TokenTransfer).Name
		a.Symbol = t.Meta.(*txtype.TokenTransfer).Symbol
		tp, ok := txtype.GetTokenType(t.Coin, t.Meta.(*txtype.TokenTransfer).TokenID)
		if !ok {
			return Asset{}, false
		}
		a.Type = tp
	case txtype.NativeTokenTransfer:
		a.Asset = asset.BuildID(t.Coin, t.Meta.(txtype.NativeTokenTransfer).TokenID)
		a.Decimals = t.Meta.(txtype.NativeTokenTransfer).Decimals
		a.Name = t.Meta.(txtype.NativeTokenTransfer).Name
		a.Symbol = t.Meta.(txtype.NativeTokenTransfer).Symbol
		tp, ok := txtype.GetTokenType(t.Coin, t.Meta.(txtype.NativeTokenTransfer).TokenID)
		if !ok {
			return Asset{}, false
		}
		a.Type = tp
	case *txtype.NativeTokenTransfer:
		a.Asset = asset.BuildID(t.Coin, t.Meta.(*txtype.NativeTokenTransfer).TokenID)
		a.Decimals = t.Meta.(*txtype.NativeTokenTransfer).Decimals
		a.Name = t.Meta.(*txtype.NativeTokenTransfer).Name
		a.Symbol = t.Meta.(*txtype.NativeTokenTransfer).Symbol
		tp, ok := txtype.GetTokenType(t.Coin, t.Meta.(*txtype.NativeTokenTransfer).TokenID)
		if !ok {
			return Asset{}, false
		}
		a.Type = tp
	case txtype.AnyAction:
		a.Asset = asset.BuildID(t.Coin, t.Meta.(txtype.AnyAction).TokenID)
		a.Decimals = t.Meta.(txtype.AnyAction).Decimals
		a.Name = t.Meta.(txtype.AnyAction).Name
		a.Symbol = t.Meta.(txtype.AnyAction).Symbol
		tp, ok := txtype.GetTokenType(t.Coin, t.Meta.(txtype.AnyAction).TokenID)
		if !ok {
			return Asset{}, false
		}
		a.Type = tp
	case *txtype.AnyAction:
		a.Asset = asset.BuildID(t.Coin, t.Meta.(*txtype.AnyAction).TokenID)
		a.Decimals = t.Meta.(*txtype.AnyAction).Decimals
		a.Name = t.Meta.(*txtype.AnyAction).Name
		a.Symbol = t.Meta.(*txtype.AnyAction).Symbol
		tp, ok := txtype.GetTokenType(t.Coin, t.Meta.(*txtype.AnyAction).TokenID)
		if !ok {
			return Asset{}, false
		}
		a.Type = tp
	default:
		return Asset{}, false
	}
	a.Coin = t.Coin

	if a.IsValid() != nil {
		return Asset{}, false
	}

	return a, true
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
