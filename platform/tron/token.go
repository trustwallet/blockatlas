package tron

import (
	"github.com/trustwallet/golibs/asset"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTokenListByAddress(address string) ([]types.Token, error) {
	return []types.Token{}, nil
}

func (p *Platform) GetTokenListIdsByAddress(address string) ([]string, error) {
	assetIds := make([]string, 0)
	tokens, err := p.client.fetchAccount(address)
	if err != nil {
		return assetIds, err
	}
	if len(tokens.Data) == 0 {
		return assetIds, nil
	}
	for _, trc20Tokens := range tokens.Data[0].Trc20 {
		for assetId := range trc20Tokens {
			assetIds = append(assetIds, asset.BuildID(p.Coin().ID, assetId))
		}
	}

	return assetIds, nil
}
