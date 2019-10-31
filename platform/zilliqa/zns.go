package zilliqa

import (
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type ZNSResponse struct {
	Addresses map[string]string
}

func (p *Platform) Lookup(coin uint64, name string) (blockatlas.Resolved, error) {
	var resp ZNSResponse
	result := blockatlas.Resolved{
		Coin: coin,
	}
	err := p.udClient.Get(&resp, "/"+name, nil)
	if err != nil {
		return result, err
	}
	symbol := CoinType.Coins[uint(coin)].Symbol
	result.Result = resp.Addresses[symbol]
	return result, nil
}
