package coin

type ExternalCoin struct {
	Coin     uint   `json:"coin"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals uint   `json:"decimals"`
}

func (c *Coin) External() *ExternalCoin {
	return &ExternalCoin{
		Coin:     c.ID,
		Name:     c.Name,
		Symbol:   c.Symbol,
		Decimals: c.Decimals,
	}
}
