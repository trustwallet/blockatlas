package ethereum_blockbook

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	account, err := p.client.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(account.Tokens, *p), nil
}

// NormalizeTxs converts multiple Ethereum tokens
func NormalizeTokens(srcTokens []Token, p Platform) []blockatlas.Token {
	tokenPage := make([]blockatlas.Token, 0)
	for _, srcToken := range srcTokens {
		token, ok := NormalizeToken(&srcToken, p.CoinIndex)
		if !ok {
			continue
		}
		tokenPage = append(tokenPage, token)
	}
	return tokenPage
}

// NormalizeToken converts a Blockbook Ethereum token into the generic model
func NormalizeToken(srcToken *Token, coinIndex uint) (t blockatlas.Token, ok bool) {
	t = blockatlas.Token{
		Name:     srcToken.Name,
		Symbol:   srcToken.Symbol,
		TokenID:  srcToken.Contract,
		Coin:     coinIndex,
		Decimals: srcToken.Decimals,
		Type:     blockatlas.TokenTypeERC20,
	}

	return t, true
}
