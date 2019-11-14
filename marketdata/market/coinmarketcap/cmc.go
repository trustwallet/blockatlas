package cmc

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/market"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"math/big"
	"net/url"
	"time"
)

type Market struct {
	market.Market
}

func InitMarket() market.Provider {
	m := &Market{
		Market: market.Market{
			Id:         "cmc",
			Name:       "CoinMarketCap",
			URL:        "https://coinmarketcap.com/",
			Request:    blockatlas.InitClient(viper.GetString("market.cmc_api")),
			UpdateTime: time.Second * 30,
		},
	}
	m.Headers["X-CMC_PRO_API_KEY"] = viper.GetString("market.cmc_api_key")
	return m
}

func (m *Market) GetData() (blockatlas.Tickers, error) {
	var prices CoinPrices
	err := m.Get(&prices, "v1/cryptocurrency/listings/latest", url.Values{"limit": {"5000"}, "convert": {"USD"}})
	if err != nil {
		return nil, err
	}
	return normalizeTickers(prices, m.GetId()), nil
}

func normalizeTicker(price Data, provider string) (blockatlas.Ticker, error) {
	tokenId := ""
	symbol := price.Symbol
	coinType := blockatlas.TypeCoin
	if price.Platform != nil {
		coinType = blockatlas.TypeToken
		symbol = price.Platform.Symbol
		tokenId = price.Symbol
	}
	return blockatlas.Ticker{
		CoinName: symbol,
		CoinType: coinType,
		TokenId:  tokenId,
		Price: blockatlas.TickerPrice{
			Value:     big.NewFloat(price.Quote.USD.Price),
			Change24h: big.NewFloat(price.Quote.USD.PercentChange24h),
			Currency:  "USD",
			Provider:  provider,
		},
		LastUpdate: price.LastUpdated,
	}, nil
}

func normalizeTickers(prices CoinPrices, provider string) (tickers blockatlas.Tickers) {
	for _, price := range prices.Data {
		t, err := normalizeTicker(price, provider)
		if err != nil {
			logger.Error(err)
			continue
		}
		tickers = append(tickers, t)
	}
	return
}

func percentageChange(value, percent float64) float64 {
	return value * (percent / 100)
}
