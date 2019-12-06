package cmc

import (
	"github.com/trustwallet/blockatlas/marketdata/cmcmap"
	"github.com/trustwallet/blockatlas/marketdata/market"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
)

const (
	id = "cmc"
)

type Market struct {
	market.Market
}

func InitMarket(api string, apiKey string, updateTime string) market.Provider {
	m := &Market{
		Market: market.Market{
			Id:         id,
			Request:    blockatlas.InitClient(api),
			UpdateTime: updateTime,
		},
	}
	m.Headers["X-CMC_PRO_API_KEY"] = apiKey
	return m
}

func (m *Market) GetData() (blockatlas.Tickers, error) {
	cmap, err := cmcmap.GetCmcMap()
	if err != nil {
		return nil, err
	}
	var prices CoinPrices
	err = m.Get(&prices, "v1/cryptocurrency/listings/latest",
		url.Values{"limit": {"5000"}, "convert": {blockatlas.DefaultCurrency}})
	if err != nil {
		return nil, err
	}
	return normalizeTickers(prices, m.GetId(), cmap), nil
}

func normalizeTicker(price Data, provider string, cmap cmcmap.CmcMapping) (tickers blockatlas.Tickers) {
	tokenId := ""
	coinName := price.Symbol
	coinType := blockatlas.TypeCoin
	if price.Platform != nil {
		coinType = blockatlas.TypeToken
		coinName = price.Platform.Symbol
		tokenId = price.Platform.TokenAddress
		if len(tokenId) == 0 {
			tokenId = price.Symbol
		}
	}

	cmcCoin, err := cmap.GetCoins(price.Id)
	if err != nil {
		tickers = append(tickers, &blockatlas.Ticker{
			CoinName: coinName,
			CoinType: coinType,
			TokenId:  tokenId,
			Price: blockatlas.TickerPrice{
				Value:     price.Quote.USD.Price,
				Change24h: price.Quote.USD.PercentChange24h,
				Currency:  blockatlas.DefaultCurrency,
				Provider:  provider,
			},
			LastUpdate: price.LastUpdated,
		})
		return
	}

	for _, cmc := range cmcCoin {
		coinName = cmc.Coin.Symbol
		if cmc.CoinType == blockatlas.TypeCoin {
			tokenId = ""
		} else if len(cmc.TokenId) > 0 {
			tokenId = cmc.TokenId
		}
		tickers = append(tickers, &blockatlas.Ticker{
			CoinName: coinName,
			CoinType: cmc.CoinType,
			TokenId:  tokenId,
			Price: blockatlas.TickerPrice{
				Value:     price.Quote.USD.Price,
				Change24h: price.Quote.USD.PercentChange24h,
				Currency:  blockatlas.DefaultCurrency,
				Provider:  provider,
			},
			LastUpdate: price.LastUpdated,
		})
	}
	return
}

func normalizeTickers(prices CoinPrices, provider string, cmap cmcmap.CmcMapping) (tickers blockatlas.Tickers) {
	for _, price := range prices.Data {
		t := normalizeTicker(price, provider, cmap)
		tickers = append(tickers, t...)
	}
	return
}
