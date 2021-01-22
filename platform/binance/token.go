package binance

import (
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTokenListByAddress(address string) ([]types.Token, error) {
	account, err := p.client.FetchAccountMeta(address)
	if err != nil || len(account.Balances) == 0 {
		return nil, nil
	}
	tokens, err := p.client.FetchTokens()
	if err != nil {
		return nil, err
	}
	return normalizeTokens(account.Balances, tokens), nil
}

func (p *Platform) GetTokenListIdsByAddress(address string) ([]string, error) {
	assets, err := p.GetTokenListByAddress(address)
	if err != nil {
		return []string{}, err
	}
	return types.GetAssetsIds(assets), nil
}
