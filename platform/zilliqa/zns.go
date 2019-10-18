package zilliqa

import (
	"github.com/spf13/viper"
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type ZNSResponse struct {
	Addresses map[string]string
}

func (p *Platform) Lookup(coin uint64, name string) (*blockatlas.Resolved, error) {
	client := blockatlas.InitClient(viper.GetString("zilliqa.lookup"))
	var resp ZNSResponse
	err := client.Get(&resp, "/"+name, nil)
	if err != nil {
		return nil, err
	}
	result := blockatlas.Resolved{
		Coin: coin,
	}
	symbol := CoinType.Coins[uint(coin)].Symbol
	result.Result = resp.Addresses[symbol]
	return &result, nil
}
