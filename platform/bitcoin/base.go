package bitcoin

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client    Client
	CoinIndex uint
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString(p.ConfigKey()))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}

func (p *Platform) ConfigKey() string {
	return fmt.Sprintf("%s.api", p.Coin().Handle)
}

func (p *Platform) GetAddressesFromXpub(xpub string) ([]string, error) {
	tokens, err := p.client.GetAddressesFromXpub(xpub)
	addresses := make([]string, 0)
	for _, token := range tokens {
		addresses = append(addresses, token.Name)
	}
	return addresses, err
}
