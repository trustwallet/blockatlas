package compound

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/marketdata/market"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"time"
)

const (
	compound = "compound"
)

type Market struct {
	market.Market
}

func InitMarket(api string, updateTime string) market.Provider {
	m := &Market{
		Market: market.Market{
			Id:         compound,
			Request:    blockatlas.InitClient(api),
			UpdateTime: updateTime,
		},
	}
	return m
}

func (m *Market) GetData() (result blockatlas.Tickers, err error) {
	var coinPrices CoinPrices
	err = m.Get(&coinPrices, "v2/ctoken", nil)
	if err != nil {
		return
	}
	result = normalizeTickers(coinPrices, m.GetId())
	return result, nil
}

func normalizeTicker(ctoken CToken, provider string) (*blockatlas.Ticker, error) {
	// TODO: add value24 calculation
	return &blockatlas.Ticker{
		CoinName: coin.Ethereum().Symbol,
		CoinType: blockatlas.TypeToken,
		TokenId:  ctoken.TokenAddress,
		Price: blockatlas.TickerPrice{
			Value:    ctoken.ExchangeRate.Value,
			Currency: blockatlas.DefaultCurrency,
			Provider: provider,
		},
		LastUpdate: time.Now(),
	}, nil
}

func normalizeTickers(prices CoinPrices, provider string) (tickers blockatlas.Tickers) {
	for _, price := range prices.Data {
		t, err := normalizeTicker(price, provider)
		if err != nil {
			continue
		}
		tickers = append(tickers, t)
	}
	return
}
