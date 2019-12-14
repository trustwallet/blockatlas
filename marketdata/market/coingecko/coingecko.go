package coingecko

import (
	"github.com/trustwallet/blockatlas/marketdata/coingecko"
	"github.com/trustwallet/blockatlas/marketdata/market"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
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

func (m *Market) normalizeTicker(price coingecko.CoinPrice, provider string) (tickers blockatlas.Tickers) {
	tokenId := ""
	coinName := strings.ToUpper(price.Symbol)
	coinType := blockatlas.TypeCoin

	cgCoin, err := m.client.GetCoinsBySymbol(price.Id)
	if err != nil {
		tickers = append(tickers, &blockatlas.Ticker{
			CoinName: coinName,
			CoinType: coinType,
			TokenId:  tokenId,
			Price: blockatlas.TickerPrice{
				Value:     price.CurrentPrice,
				Change24h: price.PriceChangePercentage24h,
				Currency:  blockatlas.DefaultCurrency,
				Provider:  provider,
			},
			LastUpdate: price.LastUpdated,
		})
		return
	}

	for _, cg := range cgCoin {
		coinName = strings.ToUpper(cg.Symbol)
		if cg.CoinType == blockatlas.TypeCoin {
			tokenId = ""
		} else if len(cg.TokenId) > 0 {
			tokenId = cg.TokenId
		}
		tickers = append(tickers, &blockatlas.Ticker{
			CoinName: coinName,
			CoinType: cg.CoinType,
			TokenId:  tokenId,
			Price: blockatlas.TickerPrice{
				Value:     price.CurrentPrice,
				Change24h: price.PriceChange24h,
				Currency:  blockatlas.DefaultCurrency,
				Provider:  provider,
			},
			LastUpdate: price.LastUpdated,
		})
	}
	return
}

func (m *Market) normalizeTickers(prices coingecko.CoinPrices, provider string) (tickers blockatlas.Tickers) {
	for _, price := range prices {
		t := m.normalizeTicker(price, provider)
		tickers = append(tickers, t...)
	}
	return
}
