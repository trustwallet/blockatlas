package ethereum

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	account, err := p.client.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(account.Docs, *p), nil
}

// NormalizeToken converts a Ethereum token into the generic model
func NormalizeToken(srcToken *Contract, coinIndex uint) (t blockatlas.Token, ok bool) {
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

	t = blockatlas.Token{
		Name:     srcToken.Name,
		Symbol:   srcToken.Symbol,
		TokenID:  srcToken.Address,
		Coin:     coinIndex,
		Decimals: srcToken.Decimals,
		Type:     tokenType,
	}

	return t, true
}

// NormalizeTxs converts multiple Ethereum tokens
func NormalizeTokens(srcTokens []Contract, p Platform) []blockatlas.Token {
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
