package fio

import (
	"strings"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

// Supported tlds
var tlds = map[string]interface{}{
	"@trust":       nil,
	"@trustwallet": nil,
	"@binance":     nil,
	"@fiomembers":  nil,
}

func (p *Platform) Match(name string) bool {
	tld := strings.ToLower(address.GetTLD(name, "@"))
	if len(tld) == 0 {
		return false
	}
	if _, ok := tlds[strings.ToLower(tld)]; ok {
		return true
	}
	// we match any @xxx domain
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
