package coingecko

import (
	"github.com/trustwallet/blockatlas/market/clients/coingecko"
	"github.com/trustwallet/blockatlas/market/market"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
)

const (
	id = "coingecko"
)

type Market struct {
	client *coingecko.Client
	cache  *coingecko.Cache
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
	coins, err := m.client.FetchCoinsList()
	if err != nil {
		return
	}
	m.cache = coingecko.NewCache(coins)

	rates := m.client.FetchLatestRates(coins, blockatlas.DefaultCurrency)
	result = m.normalizeTickers(rates, m.GetId())
	return
}

func (m *Market) normalizeTicker(price coingecko.CoinPrice, provider string) (tickers blockatlas.Tickers) {
	tokenId := ""
	coinName := strings.ToUpper(price.Symbol)
	coinType := blockatlas.TypeCoin

	cgCoin, err := m.cache.GetCoinsById(price.Id)
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
				Change24h: price.PriceChangePercentage24h,
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
