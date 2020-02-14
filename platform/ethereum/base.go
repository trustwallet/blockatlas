package ethereum

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client            Client
	collectionsClient CollectionsClient
	CoinIndex         uint
	RpcURL            string
}

func (p *Platform) Init() error {
	handle := coin.Coins[p.CoinIndex].Handle

	coinVar := fmt.Sprintf("%s.api", handle)
	p.client = Client{blockatlas.InitClient(viper.GetString(coinVar))}

	collectionsApiVar := fmt.Sprintf("%s.collections_api", handle)
	p.collectionsClient = CollectionsClient{blockatlas.InitClient(viper.GetString(collectionsApiVar))}

	collectionsApiKeyVar := fmt.Sprintf("%s.collections_api_key", handle)
	p.collectionsClient.Headers["X-API-KEY"] = viper.GetString(collectionsApiKeyVar)

	p.RpcURL = viper.GetString("ethereum.rpc")
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
