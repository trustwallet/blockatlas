package compound

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/market"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"time"
)

const (
	ETHAsset = "ETH"
)

type Market struct {
	market.Market
}

func InitMarket() market.Provider {
	m := &Market{
		Market: market.Market{
			Id:         "compound",
			Request:    blockatlas.InitClient(viper.GetString("market.compound.api")),
			UpdateTime: viper.GetString("market.compound.quote_update_time"),
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
		CoinName: ETHAsset,
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
