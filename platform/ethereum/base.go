package ethereum

import (
	"github.com/trustwallet/blockatlas/platform/bitcoin/blockbook"
	"github.com/trustwallet/blockatlas/platform/ethereum/collection"
	"github.com/trustwallet/blockatlas/platform/ethereum/trustray"
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/network/middleware"
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
		client:    &trustray.Client{Request: client.InitClient(api, middleware.SentryErrorHandler)},
	}
}

func InitWithBlockbook(coinType uint, blockbookApi, rpc string) *Platform {
	return &Platform{
		CoinIndex: coinType,
		RpcURL:    rpc,
		client:    &blockbook.Client{Request: client.InitClient(blockbookApi, middleware.SentryErrorHandler)},
	}
}

func InitWithCollection(coinType uint, rpc, blockbookApi, collectionApi, collectionKey string) *Platform {
	platform := InitWithBlockbook(coinType, blockbookApi, rpc)
	platform.collectible = collection.Client{Request: client.InitClient(collectionApi, middleware.SentryErrorHandler)}
	platform.collectible.Headers["X-API-KEY"] = collectionKey
	return platform
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
