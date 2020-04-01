package trustray

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (c *Client) GetTokenList(address string, coinIndex uint) (blockatlas.TokenPage, error) {
	account, err := c.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(account.Docs, coinIndex), nil
}

func GetTokenTypeByIndex(coinIndex uint) blockatlas.TokenType {
	var tokenType blockatlas.TokenType
	switch coinIndex {
	case coin.Ethereum().ID:
		tokenType = blockatlas.TokenTypeERC20
	case coin.Classic().ID:
		tokenType = blockatlas.TokenTypeETC20
	case coin.Poa().ID:
		tokenType = blockatlas.TokenTypePOA20
	case coin.Callisto().ID:
		tokenType = blockatlas.TokenTypeCLO20
	case coin.Wanchain().ID:
		tokenType = blockatlas.TokenTypeWAN20
	case coin.Thundertoken().ID:
		tokenType = blockatlas.TokenTypeTT20
	case coin.Gochain().ID:
		tokenType = blockatlas.TokenTypeGO20
	default:
		tokenType = "unknown"
	}
	return tokenType
}

// NormalizeToken converts a Ethereum token into the generic model
func NormalizeToken(srcToken *Contract, coinIndex uint) blockatlas.Token {
	tokenType := GetTokenTypeByIndex(coinIndex)

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
