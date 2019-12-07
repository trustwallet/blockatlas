package coingecko

import (
	"github.com/trustwallet/blockatlas/marketdata/coingecko"
	"github.com/trustwallet/blockatlas/marketdata/market"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

const (
	id = "coingecko"
)

type Market struct {
	client *coingecko.Client
	market.Market
}

func InitMarket(api string, updateTime string) market.Provider {
	m := &Market{
		client: coingecko.NewClient(api),
		Market: market.Market{
			Id:         id,
			UpdateTime: updateTime,
		},
	}
	return m
}

func (m *Market) GetData() (result blockatlas.Tickers, err error) {
	coins, err := m.client.FetchLatestRates()
	if err != nil {
		return
	}
	result = m.normalizeTickers(coins, m.GetId())
	return result, nil
}

func (m *Market) normalizeTicker(price coingecko.CoinPrice, provider string) (*blockatlas.Ticker, error) {
	tokenId := ""
	coinName := price.Symbol
	coinType := blockatlas.TypeCoin
	if len(price.CoinDetails.AssetPlatformId) > 0 {
		coinType = blockatlas.TypeToken
		coin, err := m.client.GetCoinById(price.CoinDetails.AssetPlatformId)
		if err != nil {
			return nil, err
		}
		coinName = coin.Symbol
		tokenId = price.CoinDetails.ContractAddress
		if len(tokenId) == 0 {
			tokenId = price.Symbol
		}
	}

	return &blockatlas.Ticker{
		CoinName: coinName,
		CoinType: coinType,
		TokenId:  tokenId,
		Price: blockatlas.TickerPrice{
			Value:     price.CurrentPrice,
			Currency:  blockatlas.DefaultCurrency,
			Change24h: price.PriceChange24h,
			Provider:  provider,
		},
		LastUpdate: price.LastUpdated,
	}, nil
}

func (m *Market) normalizeTickers(prices coingecko.CoinPrices, provider string) (tickers blockatlas.Tickers) {
	for _, price := range prices {
		t, err := m.normalizeTicker(price, provider)
		if err != nil {
			continue
		}
		tickers = append(tickers, t)
	}
	return
}
