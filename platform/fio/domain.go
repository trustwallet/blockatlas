package fio

import (
	"strings"

	"github.com/trustwallet/blockatlas/coin"
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

// Supported tlds
var tlds = map[string]int{
	"@trust":       CoinType.FIO,
	"@trustwallet": CoinType.FIO,
	"@binance":     CoinType.FIO,
	"@fiomembers":  CoinType.FIO,
}

func (p *Platform) Match(name string) bool {
	tld := strings.ToLower(getTLD(name))
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

// Obtain tld from then name, e.g. "@trust" from "nick@trust"
func getTLD(name string) string {
	lastSeparatorIdx := strings.LastIndex(name, "@")
	if lastSeparatorIdx < 0 || lastSeparatorIdx >= len(name)-1 {
		// no separator inside string
		return ""
	}
	// return tail including separator
	return name[lastSeparatorIdx:]
}
