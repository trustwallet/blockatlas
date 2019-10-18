package ethereum

import (
	"encoding/hex"
	"net/http"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"

	"github.com/ethereum/go-ethereum/ethclient"
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

	result.Result = mapAddress(coin, address)
	return &result, nil
}

func mapAddress(coin uint64, bytes []byte) string {
	// FIXME: convert bytes to string according to coin
	address := hex.EncodeToString(bytes)
	if address != "" {
		address = "0x" + address
	}
	return address
}
