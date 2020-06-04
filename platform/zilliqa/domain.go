package zilliqa

import (
	"strings"

	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type ZNSResponse struct {
	Addresses map[string]string
}

// Supported tlds
var tlds = map[string]int{
	".zil":    CoinType.ZIL,
	".crypto": CoinType.ZIL,
}

func (p *Platform) Match(name string) bool {
	tld := address.GetTLD(name, ".")
	if len(tld) == 0 {
		return false
	}
	_, ok := tlds[strings.ToLower(tld)]
	return ok
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
