package bitcoin

func (p *Platform) GetAddressesFromXpub(xpub string) ([]string, error) {
	tokens, err := p.client.GetAddressesFromXpub(xpub)
	addresses := make([]string, 0)
	for _, token := range tokens {
		addresses = append(addresses, token.Name)
	}
	return addresses, err
}
