package naming

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/platform"
)

func HandleLookup(name string, coins []uint64) ([]blockatlas.Resolved, error) {
	addresses := make([]blockatlas.Resolved, 0)
	apis := findHandlerApis(name, platform.NamingAPIs)
	if len(apis) == 0 {
		return nil, errors.E("platform not found", errors.Params{"name": name, "coins": coins})
	}
	for _, api := range apis {
		provAddresses, err := api.Lookup(coins, name)
		if err != nil {
			return nil, errors.E(err, "name format not recognized", errors.Params{"name": name, "coins": coins})
		}
		addresses = append(addresses, provAddresses...)
	}
	return addresses, nil
}

func findHandlerApis(name string, allApis map[uint]blockatlas.NamingServiceAPI) []blockatlas.NamingServiceAPI {
	apis := []blockatlas.NamingServiceAPI{}
	for _, api := range allApis {
		if api.CanHandle(name) {
			apis = append(apis, api)
		}
	}
	return apis
}
