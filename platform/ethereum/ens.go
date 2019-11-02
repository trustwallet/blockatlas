package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	cc "github.com/hewigovens/go-coincodec"
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	ens "github.com/wealdtech/go-ens/v3"
)

func (p *Platform) Lookup(coin uint64, name string) (blockatlas.Resolved, error) {
	client, err := ethclient.Dial(p.RpcURL)
	result := blockatlas.Resolved{
		Coin: coin,
	}
	if err != nil {
		return result, errors.E(err, "can't dial to ethereum rpc")
	}
	defer client.Close()
	// try to get multi coin address
	ensName, err := ens.NewName(client, name)
	if err != nil {
		if coin == CoinType.ETH {
			// https://github.com/wealdtech/go-ens#management-of-subdomains
			// subdomains have their own registrars they do not work with the Name interface.
			return lookupLegacyETH(client, name)
		}
		return result, errors.E(err, "query ens failed")
	}

	address, err := ensName.Address(coin)
	if err != nil {
		if coin == CoinType.ETH {
			// user may not set multi coin address
			return lookupLegacyETH(client, ensName.Name)
		}
		return result, err
	}

	encoded, err := cc.ToString(address, uint32(coin))
	result.Result = encoded
	if err != nil {
		return result, errors.E(err, "encode to address failed")
	}
	return result, nil
}

func lookupLegacyETH(client *ethclient.Client, name string) (blockatlas.Resolved, error) {
	result := blockatlas.Resolved{
		Coin: CoinType.ETH,
	}
	resolver, err := ens.NewResolver(client, name)
	if err != nil {
		return result, errors.E(err, "new ens resolver failed")
	}
	address, err := resolver.Address()
	if err != nil {
		return result, errors.E(err, "query address failed")
	}
	result.Result = address.Hex()
	return result, nil
}
