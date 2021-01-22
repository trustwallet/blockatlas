package binance

func (p *Platform) GetTokenListByAddress(address string) ([]string, error) {
	account, err := p.client.FetchAccountMeta(address)
	if err != nil || len(account.Balances) == 0 {
		return nil, nil
	}
	tokens, err := p.client.FetchTokens()
	if err != nil {
		return nil, err
	}
	return normalizeTokens(account.Balances, tokens), nil
}
