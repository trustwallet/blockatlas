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
	ens               RpcClient
}

func Init(coin uint, api, rpc string) *Platform {
	p := &Platform{
		CoinIndex: coin,
		RpcURL:    rpc,
		client:    Client{blockatlas.InitClient(api)},
		ens:       RpcClient{blockatlas.InitJSONClient(rpc)},
	}
	return p
}

func InitWitCollection(coin uint, api, rpc, collectionApi, collectionKey string) *Platform {
	p := Init(coin, api, rpc)
	p.collectionsClient = CollectionsClient{blockatlas.InitClient(collectionApi)}
	p.collectionsClient.Headers["X-API-KEY"] = collectionKey
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
