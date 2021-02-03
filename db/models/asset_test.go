package models

import (
	"reflect"
	"testing"
	"time"

	"github.com/trustwallet/golibs/types"
)

func TestAsset_isValid1(t *testing.T) {
	type fields struct {
		CreatedAt time.Time
		ID        uint
		Asset     string
		Decimals  uint
		Name      string
		Symbol    string
		Type      string
		Coin      uint
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"Valid asset",
			fields{
				Asset:    "c60",
				Decimals: 18,
				Name:     "Ethereum",
				Symbol:   "ETH",
				Type:     "coin",
				Coin:     60,
			},
			false,
		},
		{
			"Bytes32 name",
			fields{
				Asset:    "c60_t0xfFED56a180f23fD32Bc6A1d8d3c09c283aB594A8",
				Decimals: 0,
				Name:     "\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000\\u0000",
				Symbol:   "FL",
				Type:     "ERC20",
				Coin:     60,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset := &Asset{
				CreatedAt: tt.fields.CreatedAt,
				ID:        tt.fields.ID,
				Asset:     tt.fields.Asset,
				Decimals:  tt.fields.Decimals,
				Name:      tt.fields.Name,
				Symbol:    tt.fields.Symbol,
				Type:      tt.fields.Type,
				Coin:      tt.fields.Coin,
			}
			if err := asset.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("isValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAssetsFrom(t *testing.T) {
	type args struct {
		t types.Tx
	}
	tests := []struct {
		name       string
		args       args
		wantAssets []Asset
	}{
		{
			"Token Transfer",
			args{t: types.Tx{
				Coin: 60,
				Type: types.TxTokenTransfer,
				TokenTransfers: []types.TokenTransfer{
					{
						Name:     "Trust Wallet",
						Symbol:   "TWT",
						TokenID:  "TWT-123",
						Decimals: 8,
						Value:    "1",
						From:     "0x1",
						To:       "0x2",
					},
				},
				Meta: types.TokenTransfer{
					Name:     "TW 2",
					Symbol:   "TWT2",
					TokenID:  "TWT2-123",
					Decimals: 18,
					Value:    "0",
					From:     "0x2",
					To:       "0x123",
				},
			}},
			[]Asset{
				{
					Asset:    "c60_tTWT2-123",
					Decimals: 18,
					Name:     "TW 2",
					Symbol:   "TWT2",
					Type:     "ERC20",
					Coin:     60,
				},
				{
					Asset:    "c60_tTWT-123",
					Decimals: 8,
					Name:     "Trust Wallet",
					Symbol:   "TWT",
					Type:     "ERC20",
					Coin:     60,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAssets := AssetsFrom(tt.args.t); !reflect.DeepEqual(gotAssets, tt.wantAssets) {
				t.Errorf("AssetsFrom() = %v, want %v", gotAssets, tt.wantAssets)
			}
		})
	}
}
