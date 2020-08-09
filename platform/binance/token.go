package binance

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	account, err := p.client.FetchAccountMeta(address)
	if err != nil || len(account.Balances) == 0 {
		return []blockatlas.Token{}, nil
	}
	tokens, err := p.client.FetchTokens()
	if err != nil {
		return nil, err
	}
	return normalizeTokens(account.Balances, tokens), nil
}
