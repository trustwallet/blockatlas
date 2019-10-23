package ethereum

import (
	"encoding/hex"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	cc "github.com/hewigovens/go-coincodec"
	"github.com/spf13/viper"
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
		return nil, errors.E(http.StatusInternalServerError, err.Error())
	}

	address, err := ensName.Address(coin)
	if err != nil {
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
