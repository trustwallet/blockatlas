package models

import (
	"testing"
	"time"
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
