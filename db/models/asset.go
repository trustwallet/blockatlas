package models

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/trustwallet/golibs/asset"
	"github.com/trustwallet/golibs/types"
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

func AssetsFrom(t types.Tx) (assets []Asset) {
	var a Asset
	var ok bool

	if t.Type == types.TxContractCall && len(t.TokenTransfers) > 0 {
		for _, transfer := range t.TokenTransfers {
			a, ok = AssetFromTokenTransfer(&t, &transfer)
			if !ok || a.IsValid() != nil {
				assets = append(assets, a)
			}
		}
	} else {
		if t.Type == types.TxTokenTransfer {
			transfer := t.Meta.(types.TokenTransfer)
			a, ok = AssetFromTokenTransfer(&t, &transfer)
		} else if t.Type == types.TxNativeTokenTransfer {
			transfer := t.Meta.(types.NativeTokenTransfer)
			a, ok = AssetFromNativeTokenTransfer(&t, &transfer)
		} else if t.Type == types.TxAnyAction {
			action := t.Meta.(types.AnyAction)
			a, ok = AssetFromAnyAction(&t, &action)
		}
		if !ok || a.IsValid() != nil {
			assets = append(assets, a)
		}
	}
	return
}

func AssetFromTokenTransfer(t *types.Tx, transfer *types.TokenTransfer) (a Asset, ok bool) {
	tp, ok := types.GetTokenType(t.Coin, transfer.TokenID)
	if !ok {
		return
	}
	a.Asset = asset.BuildID(t.Coin, transfer.TokenID)
	a.Decimals = transfer.Decimals
	a.Name = transfer.Name
	a.Symbol = transfer.Symbol
	a.Type = tp
	a.Coin = t.Coin
	return
}

func AssetFromNativeTokenTransfer(t *types.Tx, transfer *types.NativeTokenTransfer) (a Asset, ok bool) {
	tp, ok := types.GetTokenType(t.Coin, transfer.TokenID)
	if !ok {
		return
	}
	a.Asset = asset.BuildID(t.Coin, transfer.TokenID)
	a.Decimals = transfer.Decimals
	a.Name = transfer.Name
	a.Symbol = transfer.Symbol
	a.Type = tp
	a.Coin = t.Coin
	return
}

func AssetFromAnyAction(t *types.Tx, action *types.AnyAction) (a Asset, ok bool) {
	tp, ok := types.GetTokenType(t.Coin, action.TokenID)
	if !ok {
		return
	}
	a.Asset = asset.BuildID(t.Coin, action.TokenID)
	a.Decimals = action.Decimals
	a.Name = action.Name
	a.Symbol = action.Symbol
	a.Type = tp
	a.Coin = t.Coin
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
