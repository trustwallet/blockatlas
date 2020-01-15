package cmc

import "time"

type Charts struct {
	Data ChartQuotes `json:"data"`
}

type ChartQuotes map[string]ChartQuoteValues

type ChartQuoteValues map[string][]float64

type ChartInfo struct {
	Data ChartInfoData `json:"data"`
}

type ChartInfoData struct {
	Rank              uint32                    `json:"rank"`
	CirculatingSupply float64                   `json:"circulating_supply"`
	TotalSupply       float64                   `json:"total_supply"`
	Quotes            map[string]ChartInfoQuote `json:"quotes"`
}

type ChartInfoQuote struct {
	Price     float64 `json:"price"`
	Volume24  float64 `json:"volume_24h"`
	MarketCap float64 `json:"market_cap"`
}

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
