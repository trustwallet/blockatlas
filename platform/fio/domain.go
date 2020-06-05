package fio

import (
	"strings"

	"github.com/minio/minio-go/pkg/set"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/naming"
)

var domains = set.CreateStringSet(
	"@trust",
	"@trustwallet",
	"@binance",
	"@fiomembers",
)

func (p *Platform) CanHandle(name string) bool {
	tld := strings.ToLower(naming.GetTLD(name, "@"))
	if len(tld) == 0 {
		return false
	}
	if domains.Contains(strings.ToLower(tld)) {
		return true
	}
	// we match any @xxx domain!
	if len(tld) >= 2 {
		return true
	}
	return false
}

func (p *Platform) Lookup(coins []uint64, name string) ([]blockatlas.Resolved, error) {
	var result []blockatlas.Resolved
	for _, coinId := range coins {
		coinObj := coin.Coins[uint(coinId)]
		address, err := p.client.lookupPubAddress(name, coinObj.Symbol)
		if err != nil {
			return result, err
		}
		result = append(result, blockatlas.Resolved{Coin: coinId, Result: address})
	}

	return result, nil
}
