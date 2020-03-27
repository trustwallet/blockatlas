package ethereum

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/ethereum/blockbook"
	"github.com/trustwallet/blockatlas/platform/ethereum/collection"
	"github.com/trustwallet/blockatlas/platform/ethereum/ens"
	"github.com/trustwallet/blockatlas/platform/ethereum/trustray"
)

type Platform struct {
	CoinIndex   uint
	RpcURL      string
	trustray    trustray.Client
	blockbook   blockbook.Client
	collectible collection.Client
	ens         ens.RpcClient
}

func Init(coin uint, api, rpc string) *Platform {
	return &Platform{
		CoinIndex: coin,
		RpcURL:    rpc,
		trustray:  trustray.Client{blockatlas.InitClient(api)},
		ens:       ens.RpcClient{blockatlas.InitJSONClient(rpc)},
	}
}

func InitWitCollection(coin uint, api, rpc, collectionApi, collectionKey string) *Platform {
	p := Init(coin, api, rpc)
	p.collectible = collection.Client{blockatlas.InitClient(collectionApi)}
	p.collectible.Headers["X-API-KEY"] = collectionKey
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
