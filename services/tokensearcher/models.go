package tokensearcher

import (
	"github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/watchmarket/pkg/watchmarket"
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
		nr.AssetsByAddress[key] = append(r, watchmarket.BuildID(t.Coin, t.TokenID))
	}
	nr.Unlock()
}
