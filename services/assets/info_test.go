package assets

import (
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

func Test_getCoinInfoUrl(t *testing.T) {
	type args struct {
		c     coin.Coin
		token string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test Ethereum coin", args{coin.Ethereum(), ""}, AssetsURL + coin.Ethereum().Handle + "/info"},
		{"test Ethereum token", args{coin.Ethereum(), "0x0000000000b3F879cb30FE243b4Dfee438691c04"}, AssetsURL + coin.Ethereum().Handle + "/assets/" + "0x0000000000b3F879cb30FE243b4Dfee438691c04"},
		{"test Binance coin", args{coin.Binance(), ""}, AssetsURL + coin.Binance().Handle + "/info"},
		{"test Binance token", args{coin.Binance(), "BUSD-BD1"}, AssetsURL + coin.Binance().Handle + "/assets/" + "BUSD-BD1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCoinInfoUrl(tt.args.c, tt.args.token); got != tt.want {
				t.Errorf("getCoinInfoUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
