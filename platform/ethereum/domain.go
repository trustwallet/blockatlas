package ethereum

import (
	CoinType "github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform/ethereum/ens"
	AddressEncoder "github.com/trustwallet/ens-coincodec"
)

func (p *Platform) Lookup(coins []uint64, name string) ([]blockatlas.Resolved, error) {
	var result []blockatlas.Resolved
	node, err := ens.NameHash(name)
	if err != nil {
		return result, errors.E(err, "name hash failed")
	}
	for _, coin := range coins {
		resolver, err := p.ens.Resolver(node[:])
		if err != nil {
			return result, errors.E(err, "query resolver failed")
		}
		// try to get multi coin address
		address, err := p.addressForCoin("0x"+resolver, node[:], coin)
		if err != nil {
			logger.Error(errors.E(err, errors.Params{"coin": coin, "name": name}))
			continue
		}
		result = append(result, blockatlas.Resolved{Coin: coin, Result: address})
	}

	return result, nil
}

func (p *Platform) addressForCoin(resovler string, node []byte, coin uint64) (string, error) {
	result, err := p.ens.Addr(resovler, node, coin)
	if err != nil {
		if coin == CoinType.ETH {
			// user may not set multi coin address
			result, err := p.lookupLegacyETH(resovler, node)
			if err != nil {
				return "", errors.E(err, "query legacy address failed")
			}
			return result, nil
		}
		return "", errors.E(err, "query multi coin address failed")
	}
	encoded, err := AddressEncoder.ToString(result, uint32(coin))
	if err != nil {
		return "", errors.E(err, "encode to address failed")
	}
	return encoded, nil
}

func (p *Platform) lookupLegacyETH(resolver string, node []byte) (string, error) {
	return p.ens.LegacyAddr(resolver, node)
}
