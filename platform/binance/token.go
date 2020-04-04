package binance

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
)

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	account, err := p.client.GetAccountMetadata(address)
	if err != nil || len(account.Balances) == 0 {
		return []blockatlas.Token{}, nil
	}
	tokens, err := p.client.GetTokens()
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(account.Balances, tokens), nil
}

// NormalizeTxs converts multiple Binance tokens
func NormalizeTokens(srcBalance []Balance, tokens *TokenPage) (tokenPage []blockatlas.Token) {
	for _, srcToken := range srcBalance {
		token, ok := NormalizeToken(&srcToken, tokens)
		if !ok {
			continue
		}
		tokenPage = append(tokenPage, token)
	}
	return
}

// NormalizeToken converts a Binance token into the generic model
func NormalizeToken(srcToken *Balance, tokens *TokenPage) (t blockatlas.Token, ok bool) {
	if srcToken.isAllZeroBalance() {
		return t, false
	}

	tk := tokens.findToken(srcToken.Symbol)
	if tk == nil {
		return t, false
	}

	return blockatlas.Token{
		Name:     tk.Name,
		Symbol:   tk.OriginalSymbol,
		TokenID:  tk.Symbol,
		Coin:     coin.BNB,
		Decimals: uint(decimalPlaces(tk.TotalSupply)),
		Type:     blockatlas.TokenTypeBEP2,
	}, true
}

// decimalPlaces count the decimals places.
func decimalPlaces(v string) int {
	s := strings.Split(v, ".")
	if len(s) < 2 {
		return 0
	}
	return len(s[1])
}
