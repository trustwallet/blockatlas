package ethereum_blockbook

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/ethereum"
)

type Platform struct {
	CoinIndex         uint
	RpcURL            string
	client            Client
	collectionsClient ethereum.CollectionsClient
}

func Init(coin uint, api, rpc string) *Platform {
	return &Platform{
		CoinIndex: coin,
		RpcURL:    rpc,
		client:    Client{blockatlas.InitClient(api)},
	}
}

func InitWitCollection(coin uint, api, rpc, collectionApi, collectionKey string) *Platform {
	p := Init(coin, api, rpc)
	p.collectionsClient = ethereum.CollectionsClient{blockatlas.InitClient(collectionApi)}
	p.collectionsClient.Headers["X-API-KEY"] = collectionKey
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
