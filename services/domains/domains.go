package domains

import (
	"math"
	"strings"

	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/platform"
)

// TLDMapping Mapping of name TLD's to coin where they are handled
var TLDMapping = map[string]uint64{
	".eth":         CoinType.ETH,
	".xyz":         CoinType.ETH,
	".luxe":        CoinType.ETH,
	".kred":        CoinType.ETH,
	".zil":         CoinType.ZIL,
	".crypto":      CoinType.ZIL,
	"@trust":       CoinType.FIO,
	"@trustwallet": CoinType.FIO,
	"@binance":     CoinType.FIO,
	"@fiomembers":  CoinType.FIO,
	"@":            CoinType.FIO, // any FIO domain
}

func HandleLookup(name string, coins []uint64) ([]blockatlas.Resolved, error) {
	// Assumption: format of the name can be decided (top-level-domain), and at most one naming service is tried
	name = strings.ToLower(name)
	tld, err := getTLD(name)
	if err != nil {
		return nil, errors.E(err, "name format not recognized", errors.Params{"name": name, "coins": coins})
	}
	id, ok := TLDMapping[tld]
	if !ok {
		// special handling for FIO, any fio domain
		if len(tld) >= 2 && tld[0] == '@' {
			tld = string("@")
			id, ok = TLDMapping[tld]
			if !ok {
				return nil, errors.E("name not found", errors.Params{"name": name, "coins": coins, "tld": tld})
			}
		}
	}
	api, ok := platform.NamingAPIs[id]
	if !ok {
		return nil, errors.E("platform not found", errors.Params{"name": name, "coins": coins})
	}
	result, err := api.Lookup(coins, name)
	if err != nil {
		return nil, errors.E(err, "name format not recognized", errors.Params{"name": name, "coins": coins})
	}
	return result, nil
}

// Obtain tld from then name, e.g. ".ens" from "nick.ens"
func getTLD(name string) (tld string, error error) {
	// find last separator
	lastSeparatorIdx := int(math.Max(
		float64(strings.LastIndex(name, ".")),
		float64(strings.LastIndex(name, "@"))))
	if lastSeparatorIdx <= -1 || lastSeparatorIdx >= len(name)-1 {
		// no separator inside string
		return "", errors.E("No TLD found in name", errors.Params{"name": name})
	}
	// return tail including separator
	return name[lastSeparatorIdx:], nil
}
