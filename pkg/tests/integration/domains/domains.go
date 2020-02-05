// +build integration

package domains

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/domains"
	"testing"
)

func TestDomains(t *testing.T) {
	var tests = []struct {
		name     string
		domain   string
		coins    []uint64
		Expected []blockatlas.Resolved
		wantErr  bool
	}{
		{
			"test .eth domain",
			"vitalik.eth",
			[]uint64{coin.ETH},
			[]blockatlas.Resolved{{Result: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045", Coin: coin.ETH}},
			false,
		},
		{
			"test .luxe domain",
			"vitalik.luxe",
			[]uint64{coin.ETH},
			[]blockatlas.Resolved{{Result: "0xD8A667312D5260F12a306Ae7730C754d938da86c", Coin: coin.ETH}},
			false,
		},
		{
			"test .xyz domain",
			"ourxyzwallet.xyz",
			[]uint64{coin.ETH},
			[]blockatlas.Resolved{{Result: "0x0C54eEAd78d555bE3cbCD451424F9A27a7843935", Coin: coin.ETH}},
			false,
		},
		{
			"test .zil domain",
			"dpantani.zil",
			[]uint64{coin.ZIL},
			[]blockatlas.Resolved{{Result: "zil1vdntvlk47j9kh9a85klqcd9rvgze06ruhmna64", Coin: coin.ZIL}},
			false,
		},
		{
			"test batch .zil domains",
			"dpantani.zil",
			[]uint64{coin.BTC, coin.ETH, coin.ZIL, coin.LTC, coin.BNB, coin.BCH, coin.DOGE, coin.XRP},
			[]blockatlas.Resolved{
				{Result: "bc1qd7eystu9xl53hkyxm4kyg7h5yk4p436sqx6f27", Coin: coin.BTC},
				{Result: "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB", Coin: coin.ETH},
				{Result: "zil1vdntvlk47j9kh9a85klqcd9rvgze06ruhmna64", Coin: coin.ZIL},
				{Result: "ltc1qz6nd472gx5gl3urfeldkrhg3h83c8tp2m7m6sd", Coin: coin.LTC},
				{Result: "bnb1h4vyuuytu4rm86ust29wwlevt95du52383cctm", Coin: coin.BNB},
				{Result: "qzpjlfnzudeu83krv0yk0r2kys67qptj6ys6eg6dms", Coin: coin.BCH},
				{Result: "DP9VmQyDMyB1TWwgXkyRpBa7rTfPYgMvjy", Coin: coin.DOGE},
				{Result: "rUvXBttEXhdwaKjEM2MxbtswHU6AMhUTgJ", Coin: coin.XRP},
			},
			false,
		},
		{
			"test .crypto domain",
			"dpantani.crypto",
			[]uint64{coin.ZIL},
			[]blockatlas.Resolved{
				{Result: "zil1vdntvlk47j9kh9a85klqcd9rvgze06ruhmna64", Coin: coin.ZIL},
			},
			false,
		},
		{
			"test batch .crypto domains",
			"dpantani.crypto",
			[]uint64{coin.BTC, coin.ETH, coin.ZIL, coin.LTC, coin.BNB, coin.BCH, coin.DOGE, coin.XRP},
			[]blockatlas.Resolved{
				{Result: "bc1qd7eystu9xl53hkyxm4kyg7h5yk4p436sqx6f27", Coin: coin.BTC},
				{Result: "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB", Coin: coin.ETH},
				{Result: "zil1vdntvlk47j9kh9a85klqcd9rvgze06ruhmna64", Coin: coin.ZIL},
				{Result: "ltc1qz6nd472gx5gl3urfeldkrhg3h83c8tp2m7m6sd", Coin: coin.LTC},
				{Result: "bnb1h4vyuuytu4rm86ust29wwlevt95du52383cctm", Coin: coin.BNB},
				{Result: "qzpjlfnzudeu83krv0yk0r2kys67qptj6ys6eg6dms", Coin: coin.BCH},
				{Result: "DP9VmQyDMyB1TWwgXkyRpBa7rTfPYgMvjy", Coin: coin.DOGE},
				{Result: "rUvXBttEXhdwaKjEM2MxbtswHU6AMhUTgJ", Coin: coin.XRP},
			},
			false,
		},
		{
			"test batch with invalid coin",
			"vitalik.eth",
			[]uint64{44, coin.ETH},
			[]blockatlas.Resolved{{Result: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045", Coin: coin.ETH}},
			false,
		},
		{"test invalid coin", "vitalik.eth", []uint64{44}, nil, false},
		{"test invalid name", "z9z9z900s982jidhwallet.eth", []uint64{coin.ALGO}, nil, true},
		{"test invalid domain", "vitalik.blabla", []uint64{coin.ALGO}, nil, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := domains.HandleLookup(test.domain, test.coins)
			if test.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.EqualValues(t, test.Expected, got)
		})
	}
}
