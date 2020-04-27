package zilliqa

import (
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type ZNSResponse struct {
	Addresses map[string]string
}

func (p *Platform) Lookup(coins []uint64, name string) ([]blockatlas.Resolved, error) {
	var result []blockatlas.Resolved
	resp, err := p.udClient.LookupName(name)
	if err != nil {
		return result, err
	}
	for _, coin := range coins {
		symbol := CoinType.Coins[uint(coin)].Symbol
		address := resp.Addresses[symbol]
		if len(address) == 0 {
			continue
		}
		result = append(result, blockatlas.Resolved{Coin: coin, Result: address})
	}
	return result, nil
}
