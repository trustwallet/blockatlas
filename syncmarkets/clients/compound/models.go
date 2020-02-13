package compound

type CoinPrices struct {
	Data []CToken `json:"cToken"`
}

type CToken struct {
	UnderlyingPrice Amount `json:"underlying_price"`
	ExchangeRate    Amount `json:"exchange_rate"`
	Symbol          string `json:"symbol"`
	TokenAddress    string `json:"token_address"`
}

type Amount struct {
	Value float64 `json:"value,string"`
}
