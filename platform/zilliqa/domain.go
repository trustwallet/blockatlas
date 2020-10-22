package zilliqa

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/naming"
	"github.com/trustwallet/golibs/coin"
)

type ZNSResponse struct {
	Addresses map[string]string
}

func (p *Platform) CanHandle(name string) bool {
	switch naming.GetTopDomain(name, ".") {
	case ".zil":
		return true
	case ".crypto":
		return true
	}
	return false
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
