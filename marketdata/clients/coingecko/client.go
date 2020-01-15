package coingecko

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	bucketSize = 500
)

type Client struct {
	blockatlas.Request
}

func NewClient(api string) *Client {
	c := Client{
		Request: blockatlas.InitClient(api),
	}
	return &c
}

func (c *Client) FetchLatestRates(coins GeckoCoins, currency string) (prices CoinPrices) {
	ci := coins.coinIds()

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
				"vs_currency": {currency},
				"sparkline":   {"false"},
				"ids":         {strings.Join(bucket[:], ",")},
			}

			var cp CoinPrices
			err := c.Get(&cp, "v3/coins/markets", values)
			if err != nil {
				logger.Error(err)
				return
			}
			prChan <- cp
		}(i)

		i += bucketSize
	}

	go func() {
		wg.Wait()
		close(prChan)
	}()

	for bucket := range prChan {
		prices = append(prices, bucket...)
	}

	return
}

func (c *Client) GetChartsData(id string, currency string, timeStart int64, timeEnd int64) (charts Charts, err error) {
	values := url.Values{
		"vs_currency": {currency},
		"from":        {strconv.FormatInt(timeStart, 10)},
		"to":          {strconv.FormatInt(timeEnd, 10)},
	}
	err = c.GetWithCache(&charts, fmt.Sprintf("v3/coins/%s/market_chart/range", id), values, time.Minute*5)
	return
}

func (c *Client) FetchCoinsList() (coins GeckoCoins, err error) {
	values := url.Values{
		"include_platform": {"true"},
	}
	err = c.GetWithCache(&coins, "v3/coins/list", values, time.Hour)
	return
}

func (coins GeckoCoins) coinIds() []string {
	coinIds := make([]string, 0)
	for _, coin := range coins {
		coinIds = append(coinIds, coin.Id)
	}
	return coinIds
}
