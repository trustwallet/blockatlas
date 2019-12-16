package coingecko

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"net/url"
	"strings"
	"sync"
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
	return &c
}

func (c *Client) FetchLatestRates() (prices CoinPrices, err error) {
	coins, err := c.fetchCoinsList()
	if err != nil {
		return
	}
	c.m = prepareCache(coins)
	ci := coinIds(coins)

	bucketSize := 500
	i := 0
	prChan := make(chan CoinPrices)
	var wg sync.WaitGroup
	for i < len(ci) {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var end = len(ci)
			if len(ci) > i+bucketSize {
				end = i + bucketSize
			}
			bucket := ci[i:end]
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
			prChan <- cp
		}(i)

		i += bucketSize
	}

	go func() {
		for bucket := range prChan {
			prices = append(prices, bucket...)
		}
	}()
	wg.Wait()
	close(prChan)

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

func prepareCache(coins GeckoCoins) map[string][]CoinResult {
	m := make(map[string][]CoinResult)
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

			_, ok = m[coin.Id]
			if !ok {
				m[coin.Id] = make([]CoinResult, 0)
			}
			m[coin.Id] = append(m[coin.Id], CoinResult{
				Symbol:   platformCoin.Symbol,
				TokenId:  address,
				CoinType: blockatlas.TypeToken,
			})
		}
	}
	return m
}

func coinIds(coins GeckoCoins) []string {
	coinIds := make([]string, 0)
	for _, coin := range coins {
		coinIds = append(coinIds, coin.Id)
	}
	return coinIds
}
