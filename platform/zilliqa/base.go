package zilliqa

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"

	"github.com/spf13/viper"
)

type Platform struct {
	client    Client
	rpcClient RpcClient
	udClient  Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("zilliqa.api"))}
	p.client.Headers["X-APIKEY"] = viper.GetString("zilliqa.key")

	p.rpcClient = RpcClient{blockatlas.InitClient(viper.GetString("zilliqa.rpc"))}
	p.udClient = Client{blockatlas.InitClient(viper.GetString("zilliqa.lookup"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ZIL]
}
