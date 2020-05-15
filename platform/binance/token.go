package binance

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
)

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	account, err := p.rpcClient.fetchAccountMetadata(address)
	if err != nil || len(account.Balances) == 0 {
		return []blockatlas.Token{}, nil
	}
	tokens, err := p.rpcClient.fetchTokens()
	if err != nil {
		return nil, err
	}
	return normalizeTokens(account.Balances, tokens), nil
}

// NormalizeTxs converts multiple Binance tokens
func normalizeTokens(srcBalance []Balance, tokens *TokenList) (tokenPage []blockatlas.Token) {
	for _, srcToken := range srcBalance {
		token, ok := normalizeToken(&srcToken, tokens)
		if !ok {
			continue
		}
		tokenPage = append(tokenPage, token)
	}
	return
}

// normalizeToken converts a Binance token into the generic model
func normalizeToken(srcToken *Balance, tokens *TokenList) (blockatlas.Token, bool) {
	var result blockatlas.Token
	if srcToken.isAllZeroBalance() {
		return result, false
	}

	token := tokens.findToken(srcToken.Symbol)
	if token == nil {
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

// decimalPlaces count the decimals places.
func decimalPlaces(v string) int {
	s := strings.Split(v, ".")
	if len(s) < 2 {
		return 0
	}
	return len(s[1])
}
