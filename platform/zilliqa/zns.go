package zilliqa

import (
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type ZNSResponse struct {
	Addresses map[string]string
}

func (p *Platform) Lookup(coins []int, name string) ([]blockatlas.Resolved, error) {
	var result []blockatlas.Resolved
	resp, err := p.udClient.LookupName(name)
	if err != nil {
		return result, err
	}
	for _, coin := range coins {
		symbol := CoinType.Coins[uint(coin)].Symbol
		result = append(result, blockatlas.Resolved{Coin: uint64(coin), Result: resp.Addresses[symbol]})
	}
	return result, nil
}
