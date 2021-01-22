package blockbook

import (
	"github.com/trustwallet/golibs/asset"
)

func (c *Client) GetTokenList(address string, coinIndex uint) ([]string, error) {
	tokens, err := c.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(tokens, coinIndex), nil
}

func NormalizeTokens(srcTokens []Token, coinIndex uint) []string {
	tokenPage := make([]string, 0)
	for _, srcToken := range srcTokens {
		if srcToken.Balance == "0" || srcToken.Balance == "" {
			continue
		}
		tokenPage = append(tokenPage, asset.BuildID(coinIndex, srcToken.Contract))
	}
	return tokenPage
}
