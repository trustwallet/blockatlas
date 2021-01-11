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

func AssetFrom(t txtype.Tx) (a Asset, ok bool) {
	a.Coin = t.Coin
	switch t.Meta.(type) {
	case txtype.TokenTransfer:
		transfer := t.Meta.(txtype.TokenTransfer)
		a, ok = assetFromTokenTransfer(&t, &transfer)
	case *txtype.TokenTransfer:
		transfer := t.Meta.(*txtype.TokenTransfer)
		a, ok = assetFromTokenTransfer(&t, transfer)
	case txtype.NativeTokenTransfer:
		transfer := t.Meta.(txtype.NativeTokenTransfer)
		a, ok = assetFromNativeTokenTransfer(&t, &transfer)
	case *txtype.NativeTokenTransfer:
		transfer := t.Meta.(*txtype.NativeTokenTransfer)
		a, ok = assetFromNativeTokenTransfer(&t, transfer)
	case txtype.AnyAction:
		action := t.Meta.(txtype.AnyAction)
		a, ok = assetFromAnyAction(&t, &action)
	case *txtype.AnyAction:
		action := t.Meta.(*txtype.AnyAction)
		a, ok = assetFromAnyAction(&t, action)
	default:
		break
	}

	if !ok || a.IsValid() != nil {
		return Asset{}, false
	}
	return
}

func assetFromTokenTransfer(t *txtype.Tx, transfer *txtype.TokenTransfer) (a Asset, ok bool) {
	tp, ok := txtype.GetTokenType(t.Coin, transfer.TokenID)
	if !ok {
		return
	}
	a.Asset = asset.BuildID(t.Coin, transfer.TokenID)
	a.Decimals = transfer.Decimals
	a.Name = transfer.Name
	a.Symbol = transfer.Symbol
	a.Type = tp
	return
}

func assetFromNativeTokenTransfer(t *txtype.Tx, transfer *txtype.NativeTokenTransfer) (a Asset, ok bool) {
	tp, ok := txtype.GetTokenType(t.Coin, transfer.TokenID)
	if !ok {
		return
	}
	a.Asset = asset.BuildID(t.Coin, transfer.TokenID)
	a.Decimals = transfer.Decimals
	a.Name = transfer.Name
	a.Symbol = transfer.Symbol
	a.Type = tp
	return
}

func assetFromAnyAction(t *txtype.Tx, action *txtype.AnyAction) (a Asset, ok bool) {
	tp, ok := txtype.GetTokenType(t.Coin, action.TokenID)
	if !ok {
		return
	}
	a.Asset = asset.BuildID(t.Coin, action.TokenID)
	a.Decimals = action.Decimals
	a.Name = action.Name
	a.Symbol = action.Symbol
	a.Type = tp
	return
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
