package binance

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
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

func normalizeTokens(srcBalance []TokenBalance, tokens Tokens) []blockatlas.Token {
	tokensList := make([]blockatlas.Token, 0, len(srcBalance))
	for _, srcToken := range srcBalance {
		token, ok := normalizeToken(srcToken, tokens)
		if !ok {
			continue
		}
		tokensList = append(tokensList, token)
	}
	return tokensList
}

func normalizeToken(srcToken TokenBalance, tokens Tokens) (blockatlas.Token, bool) {
	var result blockatlas.Token
	if srcToken.isAllZeroBalance() {
		return result, false
	}

	token, ok := tokens.findTokenBySymbol(srcToken.Symbol)
	if !ok {
		return result, false
	}

	result = blockatlas.Token{
		Name:     token.Name,
		Symbol:   token.OriginalSymbol,
		TokenID:  token.Symbol,
		Coin:     coin.BNB,
		Decimals: uint(decimalPlaces(token.TotalSupply)),
		Type:     blockatlas.TokenTypeBEP2,
	}

	return result, true
}

func decimalPlaces(v string) int {
	s := strings.Split(v, ".")
	if len(s) < 2 {
		return 0
	}
	return len(s[1])
}
