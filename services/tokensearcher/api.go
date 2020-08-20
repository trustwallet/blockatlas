package tokensearcher

import (
	"context"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Instance struct {
	database *db.Instance
	apis     map[uint]blockatlas.TokensAPI
}

func Init(database *db.Instance, apis map[uint]blockatlas.TokensAPI) Instance {
	return Instance{database: database, apis: apis}
}

func (i Instance) HandleTokensRequest(request map[string][]string, ctx context.Context) error {
	var addresses []string
	for coinID, requestAddresses := range request {
		for _, a := range requestAddresses {
			addresses = append(addresses, coinID+"_"+a)
		}
	}
	assetsByAddresses, err := i.database.GetAssetsMapByAddresses(addresses, ctx)
	if err != nil {
		return err
	}
	return nil
}
