package dex

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/market"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"math/big"
	"net/url"
	"strconv"
	"time"
)

const (
	quoteAsset = "BNB"
)

type Market struct {
	market.Market
}

func InitMarket() market.Provider {
	m := &Market{
		Market: market.Market{
			Id:         "dex",
			Request:    blockatlas.InitClient(viper.GetString("market.dex_api")),
			UpdateTime: viper.GetString("market.dex_quote_update_time"),
		},
	}
	return m
}

func (m *Market) GetData() (blockatlas.Tickers, error) {
	var prices []CoinPrice
	err := m.Get(&prices, "v1/ticker/24hr", url.Values{"limit": {"1000"}})
	if err != nil {
		return nil, err
	}
	rate, err := m.Storage.GetRate(quoteAsset)
	if err != nil {
		return nil, errors.E(err, "rate not found", errors.Params{"asset": quoteAsset})
	}
	result := normalizeTickers(prices, m.GetId())
	result.ApplyRate(rate.Rate.Quo(big.NewFloat(1), rate.Rate), "USD")
	return result, nil
}

func normalizeTicker(price CoinPrice, provider string) (blockatlas.Ticker, error) {
	if price.QuoteAssetName != quoteAsset {
		return blockatlas.Ticker{}, errors.E("invalid quote asset",
			errors.Params{"Symbol": price.BaseAssetName, "QuoteAsset": price.QuoteAssetName})
	}
	value, err := strconv.ParseFloat(price.LastPrice, 64)
	if err != nil {
		return blockatlas.Ticker{}, errors.E(err, "normalizeTicker parse value error",
			errors.Params{"LastPrice": price.LastPrice, "Symbol": price.BaseAssetName})
	}
	value24h, err := strconv.ParseFloat(price.PriceChangePercent, 64)
	if err != nil {
		return blockatlas.Ticker{}, errors.E(err, "normalizeTicker parse value24h error",
			errors.Params{"PriceChange": price.PriceChangePercent, "Symbol": price.BaseAssetName})
	}
	return blockatlas.Ticker{
		CoinName: price.QuoteAssetName,
		CoinType: blockatlas.TypeToken,
		TokenId:  price.BaseAssetName,
		Price: blockatlas.TickerPrice{
			Value:     big.NewFloat(value),
			Change24h: big.NewFloat(value24h),
			Currency:  "BNB",
			Provider:  provider,
		},
		LastUpdate: time.Now(),
	}, nil
}

func normalizeTickers(prices []CoinPrice, provider string) (tickers blockatlas.Tickers) {
	for _, price := range prices {
		t, err := normalizeTicker(price, provider)
		if err != nil {
			continue
		}
		tickers = append(tickers, t)
	}
	return
}
