package zilliqa

import (
	"strings"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/naming"

	"github.com/minio/minio-go/pkg/set"
)

type ZNSResponse struct {
	Addresses map[string]string
}

var domains = set.CreateStringSet(
	".zil",
	".crypto",
)

func (p *Platform) CanHandle(name string) bool {
	tld := naming.GetTLD(name, ".")
	if len(tld) == 0 {
		return false
	}
	return domains.Contains(strings.ToLower(tld))
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
