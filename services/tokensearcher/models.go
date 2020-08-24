package tokensearcher

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/watchmarket/pkg/watchmarket"
	"strconv"
	"sync"
)

type NodesResponse struct {
	sync.Mutex
	AssetsByAddress AssetsByAddress
}

func (nr *NodesResponse) UpdateAssetsByAddress(tokens blockatlas.TokenPage, coin int, address string) {
	nr.Lock()
	for _, t := range tokens {
		key := strconv.Itoa(coin) + "_" + address
		r := nr.AssetsByAddress[key]
		nr.AssetsByAddress[key] = append(r, watchmarket.BuildID(t.Coin, t.TokenID))
	}
	nr.Unlock()
}
