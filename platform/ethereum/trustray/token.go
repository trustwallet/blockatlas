package trustray

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (c *Client) GetTokenList(address string, coinIndex uint) (blockatlas.TokenPage, error) {
	account, err := c.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(account.Docs, coinIndex), nil
}

// NormalizeToken converts a Ethereum token into the generic model
func NormalizeToken(srcToken *Contract, coinIndex uint) blockatlas.Token {
	tokenType := blockatlas.GetEthereumTokenTypeByIndex(coinIndex)

	return blockatlas.Token{
		Name:     srcToken.Name,
		Symbol:   srcToken.Symbol,
		TokenID:  srcToken.Address,
		Coin:     coinIndex,
		Decimals: srcToken.Decimals,
		Type:     tokenType,
	}
}

// NormalizeTxs converts multiple Ethereum tokens
func NormalizeTokens(srcTokens []Contract, coinIndex uint) []blockatlas.Token {
	tokenPage := make([]blockatlas.Token, 0)
	for _, srcToken := range srcTokens {
		token := NormalizeToken(&srcToken, coinIndex)
		tokenPage = append(tokenPage, token)
	}
	return tokenPage
}
