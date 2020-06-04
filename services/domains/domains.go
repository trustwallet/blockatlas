package domains

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/platform"
)

func HandleLookup(name string, coins []uint64) ([]blockatlas.Resolved, error) {
	// Visit all providers, try lookup with all matching ones
	// There must be at least one provider (normally ==1)
	ret := []blockatlas.Resolved{}
	providerCount := 0 // to count number of providers visited, may be different than number of results
	for _, api := range platform.NamingAPIs {
		if !api.Match(name) {
			continue
		}
		providerCount++
		result, err := api.Lookup(coins, name)
		if err != nil {
			return nil, errors.E(err, "name format not recognized", errors.Params{"name": name, "coins": coins})
		}
		ret = append(ret, result...)
	}
	if providerCount == 0 {
		return nil, errors.E("platform not found", errors.Params{"name": name, "coins": coins})
	}
	return ret, nil
}
