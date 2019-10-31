package ethereum

import (
	"encoding/hex"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	cc "github.com/hewigovens/go-coincodec"
	"github.com/spf13/viper"
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	ens "github.com/wealdtech/go-ens/v3"
)

func (p *Platform) Lookup(coin uint64, name string) (*blockatlas.Resolved, error) {
	client, err := ethclient.Dial(viper.GetString("ethereum.rpc"))
	result := blockatlas.Resolved{
		Coin: coin,
	}
	if err != nil {
		return nil, errors.E(http.StatusInternalServerError, "can't dial to ethereum rpc")
	}
	defer client.Close()
	ensName, err := ens.NewName(client, name)
	if err != nil {
		logger.Error("Query ens failed", err, logger.Params{
			"name": name,
			"coin": coin,
		})
		return &result, nil
	}

	address, err := ensName.Address(coin)
	if err != nil {
		if coin == CoinType.ETH {
			// will remove this later
			return lookupLegacyETH(client, ensName.Name)
		}
		// return empty result for errors like: no unregistered
		return &result, nil
	}

	encoded, err := cc.ToString(address, uint32(coin))
	result.Result = encoded
	if err != nil {
		logger.Error("Encode to address failed", err, logger.Params{
			"address": hex.EncodeToString(address),
			"coin":    coin,
		})
	}
	return &result, nil
}

func lookupLegacyETH(client *ethclient.Client, name string) (*blockatlas.Resolved, error) {
	result := blockatlas.Resolved{
		Coin: CoinType.ETH,
	}
	resolver, err := ens.NewResolver(client, name)
	if err != nil {
		return &result, nil
	}
	address, err := resolver.Address()
	if err != nil {
		return &result, nil
	}
	result.Result = address.Hex()
	return &result, nil
}
