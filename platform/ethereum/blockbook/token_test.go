package blockbook

import (
	"reflect"
	"testing"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func TestNormalizeToken(t *testing.T) {
	type args struct {
		srcToken  *Token
		coinIndex uint
	}
	tests := []struct {
		name string
		args args
		want blockatlas.Token
	}{
		{
			name: "Test Normalize Token",
			args: args{
				srcToken: &Token{
					Type:     "ERC20",
					Name:     "USD//C",
					Contract: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					Symbol:   "USDC",
					Decimals: 6,
				},
				coinIndex: 60,
			},
			want: blockatlas.Token{
				Name:     "USD//C",
				Symbol:   "USDC",
				Decimals: 6,
				TokenID:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Type:     "ERC20",
				Coin:     60,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeToken(tt.args.srcToken, tt.args.coinIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NormalizeToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
