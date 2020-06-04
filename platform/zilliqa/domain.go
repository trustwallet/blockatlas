package zilliqa

import (
	"strings"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type ZNSResponse struct {
	Addresses map[string]string
}

var domains = map[string]interface{}{
	".zil":    nil,
	".crypto": nil,
}

func (p *Platform) Match(name string) bool {
	tld := address.GetTLD(name, ".")
	if len(tld) == 0 {
		return false
	}
	_, ok := domains[strings.ToLower(tld)]
	return ok
}

func (p *Platform) Lookup(coins []uint64, name string) ([]blockatlas.Resolved, error) {
	var result []blockatlas.Resolved
	resp, err := p.udClient.LookupName(name)
	if err != nil {
		return result, err
	}
	for _, c := range coins {
		symbol := coin.Coins[uint(c)].Symbol
		address := resp.Addresses[symbol]
		if len(address) == 0 {
			continue
		}
		result = append(result, blockatlas.Resolved{Coin: c, Result: address})
	}
	return result, nil
}
