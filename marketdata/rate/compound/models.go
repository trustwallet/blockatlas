package compound

type CoinPrices struct {
	Data []CToken `json:"cToken"`
}

type CToken struct {
	ExchangeRate Amount `json:"exchange_rate"`
	Symbol       string `json:"symbol"`
}

type Amount struct {
	Value float64 `json:"value,string"`
}
