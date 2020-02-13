package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	cc "github.com/hewigovens/go-coincodec"
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	ens "github.com/wealdtech/go-ens/v3"
)

func (p *Platform) Lookup(coins []uint64, name string) ([]blockatlas.Resolved, error) {
	var result []blockatlas.Resolved
	client, err := ethclient.Dial(p.RpcURL)
	if err != nil {
		return result, errors.E(err, "can't dial to ethereum rpc")
	}
	defer client.Close()
	resolver, err := ens.NewResolver(client, name)
	if err != nil {
		return result, errors.E(err, "new ens resolver failed")
	}
	for _, coin := range coins {
		// try to get multi coin address
		address, err := addressForCoin(resolver, coin, name)
		if err != nil {
			logger.Error(errors.E(err, errors.Params{"coin": coin, "name": name}))
			continue
		}
		result = append(result, blockatlas.Resolved{Coin: coin, Result: address})
	}

	return result, nil
}

func addressForCoin(resolver *ens.Resolver, coin uint64, name string) (string, error) {
	address, err := resolver.MultiAddress(coin)
	if err != nil {
		if coin == CoinType.ETH {
			// user may not set multi coin address
			result, err := lookupLegacyETH(resolver, name)
			if err != nil {
				return "", errors.E(err, "query legacy address failed")
			}
			return result, nil
		}
		return "", errors.E(err, "query multi coin address failed")
	}
	encoded, err := cc.ToString(address, uint32(coin))
	if err != nil {
		return "", errors.E(err, "encode to address failed")
	}
	return encoded, nil
}

func lookupLegacyETH(resolver *ens.Resolver, name string) (string, error) {
	address, err := resolver.Address()
	if err != nil {
		return "", errors.E(err, "query address failed")
	}
	return address.Hex(), nil
}
