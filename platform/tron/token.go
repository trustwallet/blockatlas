package tron

import (
	"github.com/trustwallet/golibs/asset"
)

func (p *Platform) GetTokenListByAddress(address string) ([]string, error) {
	assetIds := make([]string, 0)
	tokens, err := p.gridClient.fetchAccount(address)
	if err != nil {
		return assetIds, err
	}
	if len(tokens.Data) == 0 {
		return assetIds, nil
	}
	data := tokens.Data[0]

	var tokenIds []string
	for _, v := range data.AssetsV2 {
		tokenIds = append(assetIds, v.Key)
	}
	for _, trc20Tokens := range data.Trc20 {
		for key := range trc20Tokens {
			tokenIds = append(tokenIds, key)
		}
	}
	for _, tokenId := range tokenIds {
		assetIds = append(assetIds, asset.BuildID(p.Coin().ID, tokenId))
	}

	return assetIds, nil
}
