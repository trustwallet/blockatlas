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
	tokens      EthereumClient
	collectible collection.Client
	ens         ens.RpcClient
}

func Init(coinType uint, api, rpc string) *Platform {
	return &Platform{
		CoinIndex: coinType,
		RpcURL:    rpc,
		ens:       ens.RpcClient{blockatlas.InitJSONClient(rpc)},
		tokens:    &trustray.Client{blockatlas.InitClient(api)},
	}
}

func InitWitCollection(coinType uint, api, rpc, blockbookApi, collectionApi, collectionKey string) *Platform {
	platform := Init(coinType, api, rpc)
	if coinType == coin.ETH {
		platform.client = &blockbook.Client{blockatlas.InitClient(blockbookApi)}
	} else {
		platform.client = &trustray.Client{blockatlas.InitClient(api)}
	}
	platform.collectible = collection.Client{blockatlas.InitClient(collectionApi)}
	platform.collectible.Headers["X-API-KEY"] = collectionKey
	return platform
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
