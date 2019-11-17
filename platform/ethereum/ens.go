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
	resolver, err := ens.NewResolver(client, name)
	if err != nil {
		return result, errors.E(err, "new ens resolver failed")
	}
	// try to get multi coin address
	address, err := resolver.MultiAddress(coin)
	if err != nil {
		if coin == CoinType.ETH {
			// user may not set multi coin address
			return lookupLegacyETH(resolver, name)
		}
		return result, errors.E(err, "query multi coin address failed")
	}
	encoded, err := cc.ToString(address, uint32(coin))
	if err != nil {
		return result, errors.E(err, "encode to address failed")
	}
	result.Result = encoded
	return result, nil
}

func lookupLegacyETH(resolver *ens.Resolver, name string) (blockatlas.Resolved, error) {
	result := blockatlas.Resolved{
		Coin: CoinType.ETH,
	}
	address, err := resolver.Address()
	if err != nil {
		return result, errors.E(err, "query address failed")
	}
	result.Result = address.Hex()
	return result, nil
}
