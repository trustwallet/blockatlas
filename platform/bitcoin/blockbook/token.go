package blockbook

import (
	"github.com/trustwallet/golibs/types"
)

func (c *Client) GetTokenList(address string, coinIndex uint) ([]types.Token, error) {
	tokens, err := c.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(tokens, coinIndex), nil
}

func NormalizeTokens(tokens []Token, coinIndex uint) []types.Token {
	assets := make([]types.Token, 0)
	for _, srcToken := range tokens {
		if srcToken.Balance == "0" || srcToken.Balance == "" {
			continue
		}
		token := NormalizeToken(&srcToken, coinIndex)
		assets = append(assets, token)
	}
	return assets
}

func NormalizeToken(srcToken *Token, coinIndex uint) types.Token {
	return types.Token{
		Name:     srcToken.Name,
		Symbol:   srcToken.Symbol,
		TokenID:  srcToken.Contract,
		Coin:     coinIndex,
		Decimals: srcToken.Decimals,
		Type:     types.GetEthereumTokenTypeByIndex(coinIndex),
	}
}
