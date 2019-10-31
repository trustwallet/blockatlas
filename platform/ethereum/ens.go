package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	cc "github.com/hewigovens/go-coincodec"
	"github.com/pkg/errors"
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	ens "github.com/wealdtech/go-ens/v3"
)

func (p *Platform) Lookup(coin uint64, name string) blockatlas.Resolved {
	client, err := ethclient.Dial(p.RpcURL)
	result := blockatlas.Resolved{
		Coin: coin,
	}
	if err != nil {
		result.Error = errors.Wrap(err, "can't dial to ethereum rpc").Error()
		return result
	}
	defer client.Close()
	ensName, err := ens.NewName(client, name)
	if err != nil {
		result.Error = errors.Wrap(err, "query ens failed").Error()
		return result
	}

	address, err := ensName.Address(coin)
	if err != nil {
		if coin == CoinType.ETH {
			// will remove this later
			return lookupLegacyETH(client, ensName.Name)
		}
		result.Error = err.Error()
		return result
	}

	encoded, err := cc.ToString(address, uint32(coin))
	result.Result = encoded
	if err != nil {
		result.Error = errors.Wrap(err, "encode to address failed").Error()
	}
	return result
}

func lookupLegacyETH(client *ethclient.Client, name string) blockatlas.Resolved {
	result := blockatlas.Resolved{
		Coin: CoinType.ETH,
	}
	resolver, err := ens.NewResolver(client, name)
	if err != nil {
		result.Error = errors.Wrap(err, "new ens resolver failed").Error()
		return result
	}
	address, err := resolver.Address()
	if err != nil {
		result.Error = errors.Wrap(err, "query address failed").Error()
		return result
	}
	result.Result = address.Hex()
	return result
}
