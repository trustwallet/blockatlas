package blockbook

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (c *Client) GetTokenList(address string, coinIndex uint) (blockatlas.TokenPage, error) {
	account, err := c.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(account.Tokens, coinIndex), nil
}

// NormalizeTxs converts multiple Ethereum tokens
func NormalizeTokens(srcTokens []Token, coinIndex uint) []blockatlas.Token {
	tokenPage := make([]blockatlas.Token, 0)
	for _, srcToken := range srcTokens {
		token, ok := NormalizeToken(&srcToken, coinIndex)
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
