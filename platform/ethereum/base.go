package ethereum

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/ethereum/blockbook"
	"github.com/trustwallet/blockatlas/platform/ethereum/collection"
	"github.com/trustwallet/blockatlas/platform/ethereum/trustray"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	CoinIndex   uint
	RpcURL      string
	client      EthereumClient
	collectible collection.Client
}

func Init(coinType uint, api, rpc string) *Platform {
	return &Platform{
		CoinIndex: coinType,
		RpcURL:    rpc,
		client:    &trustray.Client{Request: blockatlas.InitClient(api)},
	}
}

func InitWithBlockbook(coinType uint, blockbookApi, rpc string) *Platform {
	return &Platform{
		CoinIndex: coinType,
		RpcURL:    rpc,
		client:    &blockbook.Client{Request: blockatlas.InitClient(blockbookApi)},
	}
}

func InitWithCollection(coinType uint, rpc, blockbookApi, collectionApi, collectionKey string) *Platform {
	platform := InitWithBlockbook(coinType, blockbookApi, rpc)
	platform.collectible = collection.Client{Request: blockatlas.InitClient(collectionApi)}
	platform.collectible.Headers["X-API-KEY"] = collectionKey
	return platform
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
