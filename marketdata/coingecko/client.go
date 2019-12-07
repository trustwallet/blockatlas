package coingecko

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	batchSize     = 40
	batchWaitTime = 1
)

type Client struct {
	pricesMap map[string]*Coin
	m         sync.Mutex
	blockatlas.Request
}

func NewClient(api string) *Client {
	c := Client{
		Request: blockatlas.InitClient(api),
	}
	c.pricesMap = make(map[string]*Coin)
	go c.initCoins()
	return &c
}

func (c *Client) fetchCoins() map[string]*Coin {
	c.m.Lock()
	defer c.m.Unlock()

	return c.pricesMap

}

func (c *Client) initCoins() {
	c.m.Lock()
	defer c.m.Unlock()

	coins, err := c.FetchCoinsList()
	if err != nil {
		logger.Error(err)
		return
	}
	prices := make(chan Coin)

	var wg sync.WaitGroup
	wg.Add(len(coins))
	logger.Info(len(coins))
	go func(coins CoingeckoCoins) {
		for i, coin := range coins {
			go func(coin CoingeckoCoin) {
				defer wg.Done()
				details, err := c.FetchCoinById(coin.Id)
				if err != nil {
					logger.Error(err)
					return
				}
				prices <- Coin{
					details,
					coin,
				}
			}(coin)
			if i%batchSize == 0 && i > 0 {
				time.Sleep(batchWaitTime * time.Minute)
			}
		}
	}(coins)

	go func() {
		for val := range prices {
			logger.Info(val.Id)
			c.pricesMap[val.Id] = &val
		}
	}()
	wg.Wait()
	close(prices)
}

func (c *Client) FetchLatestRates() (prices CoinPrices, err error) {
	coins := c.fetchCoins()

	var coinIds []string
	for coin := range coins {
		coinIds = append(coinIds, coin)
	}

	values := url.Values{
		"vs_currency": {blockatlas.DefaultCurrency},
		"sparkline":   {"false"},
		"ids":         {strings.Join(coinIds[:], ",")},
	}

	var cgPrices CoingeckoPrices
	err = c.Get(&cgPrices, "v3/coins/markets", values)
	if err != nil {
		return
	}

	for _, price := range cgPrices {
		prices = append(prices, CoinPrice{
			c.pricesMap[price.Id].CoinDetails,
			price,
		})
	}
	return
}

func (c *Client) GetCoinById(id string) (coin *Coin, err error) {
	coin, ok := c.pricesMap[id]
	if !ok {
		err = errors.E("No coin found by id", errors.Params{"id": id}).Err
	}
	return
}

func (c *Client) FetchCoinsList() (coins CoingeckoCoins, err error) {
	err = c.Get(&coins, "v3/coins/list", nil)
	return
}

func (c *Client) FetchCoinById(id string) (details CoinDetails, err error) {
	err = c.Get(&details, fmt.Sprintf("v3/coins/%s", id), nil)
	return
}
