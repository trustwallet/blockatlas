package blockbook

import (
	"reflect"
	"testing"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func TestNormalizeToken(t *testing.T) {
	type args struct {
		srcToken  Token
		coinIndex uint
	}
	tests := []struct {
		name string
		args args
		want []blockatlas.Token
	}{
		{
			name: "Should normalize and return token with balance",
			args: args{srcToken: Token{
				Balance:  "100",
				Type:     "ERC20",
				Name:     "USD//C",
				Contract: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Symbol:   "USDC",
				Decimals: 6},
				coinIndex: 60},
			want: []blockatlas.Token{
				{
					Type:    "ERC20",
					Name:    "USD//C",
					TokenID: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					Symbol:  "USDC", Decimals: 6, Coin: 60}},
		},
		{
			name: "Should not return token with zero balance",
			args: args{srcToken: Token{
				Balance:  "0",
				Type:     "ERC20",
				Name:     "USD//C",
				Contract: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Symbol:   "USDC",
				Decimals: 6},
				coinIndex: 60},
			want: []blockatlas.Token{},
		}, {
			name: "Should not return token with zero balance",
			args: args{srcToken: Token{
				Balance:  "",
				Type:     "ERC20",
				Name:     "USD//C",
				Contract: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Symbol:   "USDC",
				Decimals: 6},
				coinIndex: 60},
			want: []blockatlas.Token{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeTokens([]Token{tt.args.srcToken}, tt.args.coinIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NormalizeToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
