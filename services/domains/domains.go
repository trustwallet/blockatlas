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
	matchingCount := 0
	for _, api := range platform.NamingAPIs {
		if !api.Match(name) {
			continue
		}
		result, err := api.Lookup(coins, name)
		if err != nil {
			return nil, errors.E(err, "name format not recognized", errors.Params{"name": name, "coins": coins})
		}
		for _, r := range result {
			ret = append(ret, r)
		}
		matchingCount++
	}
	if matchingCount == 0 {
		return nil, errors.E("platform not found", errors.Params{"name": name, "coins": coins})
	}
	return ret, nil
}
