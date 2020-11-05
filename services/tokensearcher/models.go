package tokensearcher

import (
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/asset"
	"sync"
)

type NodesResponse struct {
	sync.Mutex
	AssetsByAddress AssetsByAddress
}

func (nr *NodesResponse) UpdateAssetsByAddress(tokens blockatlas.TokenPage, coin int, a string) {
	nr.Lock()
	for _, t := range tokens {
		key := address.PrefixedAddress(uint(coin), a)
		r := nr.AssetsByAddress[key]
		nr.AssetsByAddress[key] = append(r,
			models.Asset{
				Asset:    asset.BuildID(t.Coin, t.TokenID),
				Decimals: t.Decimals,
				Name:     t.Name,
				Symbol:   t.Symbol,
				Type:     string(t.Type),
				ID:       t.Coin,
			},
		)
	}
	nr.Unlock()
}
