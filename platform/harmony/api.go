package harmony

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client    Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("harmony.rpc"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ONE]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	_, err := p.client.GetTxsOfAddress(address)
	return nil, err
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}
