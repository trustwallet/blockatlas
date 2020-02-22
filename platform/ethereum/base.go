package ethereum

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	CoinIndex         uint
	RpcURL            string
	client            Client
	collectionsClient CollectionsClient
}

func Init(coin uint, api, rpc, collectionApi, collectionKey string) *Platform {
	p := Platform{
		CoinIndex:         coin,
		RpcURL:            rpc,
		client:            Client{blockatlas.InitClient(api)},
		collectionsClient: CollectionsClient{blockatlas.InitClient(collectionApi)},
	}
	p.collectionsClient.Headers["X-API-KEY"] = collectionKey
	return &p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
