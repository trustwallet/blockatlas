package assets

import (
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

func Test_getCoinInfoUrl(t *testing.T) {
	assetsURL := "http://localhost:8420/"
	type args struct {
		c     coin.Coin
		token string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test Ethereum coin", args{coin.Ethereum(), ""}, assetsURL + coin.Ethereum().Handle},
		{"test Ethereum token", args{coin.Ethereum(), "0x0000000000b3F879cb30FE243b4Dfee438691c04"}, assetsURL + coin.Ethereum().Handle + "/assets/" + "0x0000000000b3F879cb30FE243b4Dfee438691c04"},
		{"test Binance coin", args{coin.Binance(), ""}, assetsURL + coin.Binance().Handle},
		{"test Binance token", args{coin.Binance(), "0x0000000000b3F879cb30FE243b4Dfee438691c04"}, assetsURL + coin.Binance().Handle + "/assets/" + "0x0000000000b3F879cb30FE243b4Dfee438691c04"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCoinInfoUrl(tt.args.c, tt.args.token, assetsURL); got != tt.want {
				t.Errorf("getCoinInfoUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
