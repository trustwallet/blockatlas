package blockbook

import (
	"github.com/trustwallet/golibs/tokentype"
	"github.com/trustwallet/golibs/txtype"
)

func (c *Client) GetTokenList(address string, coinIndex uint) (txtype.TokenPage, error) {
	tokens, err := c.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(tokens, coinIndex), nil
}

func NormalizeTokens(srcTokens []Token, coinIndex uint) []txtype.Token {
	tokenPage := make([]txtype.Token, 0, len(srcTokens))
	for _, srcToken := range srcTokens {
		if srcToken.Balance == "0" || srcToken.Balance == "" {
			continue
		}
		token := NormalizeToken(&srcToken, coinIndex)
		tokenPage = append(tokenPage, token)
	}
	return tokenPage
}

func NormalizeToken(srcToken *Token, coinIndex uint) txtype.Token {
	return txtype.Token{
		Name:     srcToken.Name,
		Symbol:   srcToken.Symbol,
		TokenID:  srcToken.Contract,
		Coin:     coinIndex,
		Decimals: srcToken.Decimals,
		Type:     tokentype.GetEthereumTokenTypeByIndex(coinIndex),
	}
}
