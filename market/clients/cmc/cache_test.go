package cmc

import (
	"reflect"
	"testing"
)

func TestCmcMapping_cmcToCoinMap(t *testing.T) {
	tests := []struct {
		name  string
		c     CmcSlice
		wantM CmcMapping
	}{
		{
			"parse mapping 1",
			CmcSlice{{Id: 3}, {Id: 10}, {Id: 44}},
			map[uint][]CoinMap{3: {{Id: 3}}, 10: {{Id: 10}}, 44: {{Id: 44}}}},
		{
			"parse mapping 2",
			CmcSlice{{Id: 3}, {Id: 10}},
			map[uint][]CoinMap{3: {{Id: 3}}, 10: {{Id: 10}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := tt.c.cmcToCoinMap(); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("cmcToCoinMap() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestCmcMapping_coinToCmcMap(t *testing.T) {
	tests := []struct {
		name  string
		c     CmcSlice
		wantM CoinMapping
	}{
		{
			"parse mapping 1",
			CmcSlice{{Coin: 60, Id: 3, TokenId: "3211"}, {Coin: 60, Id: 10}, {Coin: 61, Id: 44}},
			CoinMapping{generateId(60, "3211"): {Id: 3, Coin: 60, TokenId: "3211"}, generateId(60, ""): {Coin: 60, Id: 10}, generateId(61, ""): {Coin: 61, Id: 44}}},
		{
			"parse mapping 2",
			CmcSlice{{Coin: 60, Id: 3}, {Coin: 61, Id: 10}},
			CoinMapping{generateId(60, ""): {Coin: 60, Id: 3, TokenId: ""}, generateId(61, ""): {Coin: 61, Id: 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := tt.c.coinToCmcMap(); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("coinToCmcMap() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func Test_generateId(t *testing.T) {
	type args struct {
		coin  uint
		token string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"generate id 1",
			args{coin: 60, token: ""},
			"60:",
		},
		{
			"generate id 2",
			args{coin: 60, token: "12"},
			"60:12",
		},
		{
			"generate id 2",
			args{coin: 185, token: "12"},
			"185:12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateId(tt.args.coin, tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateId() = %v, want %v", got, tt.want)
			}
		})
	}
}
