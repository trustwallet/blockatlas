package compound

import (
	"github.com/trustwallet/blockatlas/coin"
	c "github.com/trustwallet/blockatlas/market/clients/compound"
	"github.com/trustwallet/blockatlas/market/market"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"time"
)

const (
	id = "compound"
)

type Market struct {
	market.Market
	client *c.Client
}

func InitMarket(api string, updateTime string) market.Provider {
	m := &Market{
		Market: market.Market{
			Id:         id,
			UpdateTime: updateTime,
		},
		client: c.NewClient(api),
	}
	return m
}

func (m *Market) GetData() (result blockatlas.Tickers, err error) {
	coinPrices, err := m.client.GetData()
	if err != nil {
		return
	}
	result = normalizeTickers(coinPrices, m.GetId())
	return result, nil
}

func normalizeTicker(ctoken c.CToken, provider string) (*blockatlas.Ticker, error) {
	// TODO: add value24 calculation
	return &blockatlas.Ticker{
		CoinName: coin.Ethereum().Symbol,
		CoinType: blockatlas.TypeToken,
		TokenId:  ctoken.TokenAddress,
		Price: blockatlas.TickerPrice{
			Value:    ctoken.UnderlyingPrice.Value,
			Currency: coin.Coins[coin.ETH].Symbol,
			Provider: provider,
		},
		LastUpdate: time.Now(),
	}, nil
}

func normalizeTickers(prices c.CoinPrices, provider string) (tickers blockatlas.Tickers) {
	for _, price := range prices.Data {
		t, err := normalizeTicker(price, provider)
		if err != nil {
			continue
		}
		tickers = append(tickers, t)
	}
	return
}
