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
	client      EthereumClient
	collectible collection.Client
	ens         ens.RpcClient
}

func Init(coinType uint, api, rpc string) *Platform {
	return &Platform{
		CoinIndex: coinType,
		RpcURL:    rpc,
		ens:       ens.RpcClient{Request: blockatlas.InitJSONClient(rpc)},
		client:    &trustray.Client{Request: blockatlas.InitClient(api)},
	}
}

func InitWithBlockbook(coinType uint, blockbookApi, rpc string) *Platform {
	return &Platform{
		CoinIndex: coinType,
		RpcURL:    rpc,
		ens:       ens.RpcClient{Request: blockatlas.InitJSONClient(rpc)},
		client:    &blockbook.Client{Request: blockatlas.InitClient(blockbookApi)},
	}
}

func InitWitCollection(coinType uint, api, rpc, blockbookApi, collectionApi, collectionKey string) *Platform {
	platform := InitWithBlockbook(coinType, blockbookApi, rpc)
	platform.collectible = collection.Client{Request: blockatlas.InitClient(collectionApi)}
	platform.collectible.Headers["X-API-KEY"] = collectionKey
	return platform
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
