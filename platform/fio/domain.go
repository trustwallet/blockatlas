package fio

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/naming"
	"github.com/trustwallet/golibs/coin"
)

func (p *Platform) CanHandle(name string) bool {
	domain := naming.GetTopDomain(name, "@")
	if len(domain) == 0 {
		return false
	}
	switch domain {
	case "@trust":
		return true
	case "@trustwallet":
		return true
	case "@binance":
		return true
	case "@fiomembers":
		return true
	}
	// we match any @xxx domain!
	return len(domain) >= 2
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
