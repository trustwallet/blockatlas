package cmc

import "time"

type CoinPrices struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
	} `json:"status"`
	Data []Data `json:"data"`
}

type Coin struct {
	Id     uint   `json:"id"`
	Symbol string `json:"symbol"`
}

type Data struct {
	Coin
	LastUpdated time.Time `json:"last_updated"`
	Platform    *Platform `json:"platform"`
	Quote       Quote     `json:"quote"`
}

type Platform struct {
	Coin
	TokenAddress string `json:"token_address"`
}

type Quote struct {
	USD USD `json:"USD"`
}

type USD struct {
	Price            float64 `json:"price"`
	PercentChange24h float64 `json:"percent_change_24h"`
}
