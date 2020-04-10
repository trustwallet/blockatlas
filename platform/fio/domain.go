package fio

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

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

func (p *Platform) ReverseLookup(coin uint64, publicKey string) ([]blockatlas.Resolved, error) {
	var results []blockatlas.Resolved
	addresses, err := p.client.lookupPublicKey(publicKey)
	if err != nil {
		return results, err
	}
	for _, address := range addresses {
		results = append(results, blockatlas.Resolved{Coin: coin, Result: address})
	}
	return results, nil
}
