package blockatlas

import (
	"math/big"
	"time"
)

const (
	TypeCoin  CoinType = "coin"
	TypeToken CoinType = "token"
)

type CoinType string

type TickerResponse struct {
	Currency string  `json:"currency"`
	Result   Tickers `json:"result"`
}

type Ticker struct {
	Coin       uint        `json:"coin"`
	CoinName   string      `json:"coin_name,omitempty"`
	TokenId    string      `json:"token_id,omitempty"`
	CoinType   CoinType    `json:"type,omitempty"`
	Price      TickerPrice `json:"price,omitempty"`
	LastUpdate time.Time   `json:"last_update,omitempty"`
	Error      string      `json:"error,omitempty"`
}

func (t *Ticker) SetCoinId(coinId uint) {
	t.Coin = coinId
	t.CoinName = ""
	t.Price.Provider = ""
	t.Price.Currency = ""
}

type TickerPrice struct {
	Value     *big.Float `json:"value,omitempty"`
	Change24h *big.Float `json:"change_24h,omitempty"`
	Currency  string     `json:"currency,omitempty"`
	Provider  string     `json:"provider,omitempty"`
}

type Rate struct {
	Currency  string     `json:"currency"`
	Rate      *big.Float `json:"rate"`
	Timestamp int64      `json:"timestamp"`
}

type Rates []Rate
type Tickers []Ticker

func (ts Tickers) ApplyRate(rate *big.Float, currency string) {
	for _, t := range ts {
		t.ApplyRate(rate, currency)
	}
}

func (t *Ticker) ApplyRate(rate *big.Float, currency string) {
	if t.Price.Currency == currency {
		return
	}
	t.Price.Value.Mul(t.Price.Value, rate)
	t.Price.Currency = currency
}
