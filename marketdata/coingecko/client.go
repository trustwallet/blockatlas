package coingecko

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	m map[string][]CoinResult
	blockatlas.Request
}

func NewClient(api string) *Client {
	c := Client{
		Request: blockatlas.InitClient(api),
	}
	c.m = make(map[string][]CoinResult)
	return &c
}

func (c *Client) FetchLatestRates() (prices CoinPrices, err error) {
	coins, err := c.fetchCoinsList()
	if err != nil {
		return
	}

	coinsMap := make(map[string]GeckoCoin)
	for _, coin := range coins {
		coinsMap[coin.Id] = coin
	}

	for _, coin := range coins {
		for platform, address := range coin.Platforms {
			if len(platform) == 0 || len(address) == 0 {
				continue
			}
			platformCoin, ok := coinsMap[platform]
			if !ok {
				continue
			}

			_, ok = c.m[coin.Id]
			if !ok {
				c.m[coin.Id] = make([]CoinResult, 0)
			}
			c.m[coin.Id] = append(c.m[coin.Id], CoinResult{
				Symbol:   platformCoin.Symbol,
				TokenId:  address,
				CoinType: blockatlas.TypeToken,
			})
		}
	}

	coinIds := make([]string, 0)
	for _, coin := range coins {
		coinIds = append(coinIds, coin.Id)
	}

	bucketSize := 500
	i := 0
	for i < len(coinIds) {
		var end = len(coinIds)
		if len(coinIds) > i+bucketSize {
			end = i + bucketSize
		}
		bucket := coinIds[i:end]
		values := url.Values{
			"vs_currency": {blockatlas.DefaultCurrency},
			"sparkline":   {"false"},
			"ids":         {strings.Join(bucket[:], ",")},
		}

		var cp CoinPrices
		err = c.Get(&cp, "v3/coins/markets", values)
		if err != nil {
			return
		}
		prices = append(prices, cp...)
		i += bucketSize
	}

	return
}

func (c *Client) GetCoinsBySymbol(id string) (coins []CoinResult, err error) {
	coins, ok := c.m[id]
	if !ok {
		err = errors.E("No coin found by id", errors.Params{"id": id}).Err
	}
	return
}

func (c *Client) fetchCoinsList() (coins GeckoCoins, err error) {
	values := url.Values{
		"include_platform": {"true"},
	}
	err = c.GetWithCache(&coins, "v3/coins/list", values, time.Hour)
	return
}
