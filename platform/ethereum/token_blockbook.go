package ethereum

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/ethereum/blockbook"
)

func (p *Platform) GetTokenListByAddressFromBlockbook(address string) (blockatlas.TokenPage, error) {
	account, err := p.blockbook.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokensBlockbook(account.Tokens, *p), nil
}

// NormalizeTxs converts multiple Ethereum tokens
func NormalizeTokensBlockbook(srcTokens []blockbook.Token, p Platform) []blockatlas.Token {
	tokenPage := make([]blockatlas.Token, 0)
	for _, srcToken := range srcTokens {
		token, ok := NormalizeTokenBlockbook(&srcToken, p.CoinIndex)
		if !ok {
			continue
		}
		tokenPage = append(tokenPage, token)
	}
	return tokenPage
}

// NormalizeToken converts a Blockbook Ethereum token into the generic model
func NormalizeTokenBlockbook(srcToken *blockbook.Token, coinIndex uint) (t blockatlas.Token, ok bool) {
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
