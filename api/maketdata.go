package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/market"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/pkg/ginutils/gincache"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/assets"
	"github.com/trustwallet/blockatlas/storage"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	defaultMaxChartItems = 64
)

type TickerRequest struct {
	Currency string `json:"currency"`
	Assets   []Coin `json:"assets"`
}

type Coin struct {
	Coin     uint                `json:"coin"`
	CoinType blockatlas.CoinType `json:"type"`
	TokenId  string              `json:"token_id,omitempty"`
}

func SetupMarketAPI(router gin.IRouter, db storage.Market) {
	router.Use(ginutils.TokenAuthMiddleware(viper.GetString("market.auth")))
	// Ticker
	router.GET("/ticker", getTickerHandler(db))
	router.POST("/ticker", getTickersHandler(db))
	// Charts
	router.GET("/charts", gincache.CacheMiddleware(time.Minute*5, getChartsHandler()))
	router.GET("/info", getCoinInfoHandler())
}

// @Summary Get ticker value for a specific market
// @Id get_ticker
// @Description Get the ticker value from an market and coin/token
// @Accept json
// @Produce json
// @Tags Market
// @Param coin query int true "coin id"
// @Param token query string false "token id"
// @Param currency query string false "the currency to show the quote" default(USD)
// @Success 200 {object} blockatlas.Ticker
// @Router /v1/market/ticker [get]
func getTickerHandler(storage storage.Market) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		coinQuery := c.Query("coin")
		coinId, err := strconv.Atoi(coinQuery)
		if err != nil {
			ginutils.RenderError(c, http.StatusBadRequest, "Invalid coin")
			return
		}
		token := strings.ToUpper(c.Query("token"))

		currency := c.DefaultQuery("currency", blockatlas.DefaultCurrency)
		rate, err := storage.GetRate(strings.ToUpper(currency))
		if err != nil {
			logger.Error(err, "Failed to retrieve currency", logger.Params{"coin": coinId, "currency": currency})
			ginutils.RenderError(c, http.StatusInternalServerError, "Failed to retrieve currency")
			return
		}

		symbol := coin.Coins[uint(coinId)].Symbol
		result, err := storage.GetTicker(symbol, token)
		if err != nil {
			logger.Error(err, "Failed to retrieve ticker", logger.Params{"coin": coinId, "currency": currency})
			ginutils.RenderError(c, http.StatusInternalServerError, err.Error())
			return
		}
		result.ApplyRate(currency, rate.Rate, rate.PercentChange24h)
		ginutils.RenderSuccess(c, result)
	}
}

// @Summary Get ticker values for a specific market
// @Id get_tickers
// @Description Get the ticker values from many market and coin/token
// @Accept json
// @Produce json
// @Tags Market
// @Param tickers body api.TickerRequest true "Ticker"
// @Success 200 {object} blockatlas.Tickers
// @Router /v1/market/ticker [post]
func getTickersHandler(storage storage.Market) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		md := TickerRequest{Currency: blockatlas.DefaultCurrency}
		if err := c.BindJSON(&md); err != nil {
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		rate, err := storage.GetRate(strings.ToUpper(md.Currency))
		if err != nil {
			logger.Error(err, "Failed to retrieve rate", logger.Params{"currency": md.Currency})
			ginutils.RenderError(c, http.StatusInternalServerError, "Failed to retrieve rate")
			return
		}

		tickers := make(blockatlas.Tickers, 0)
		for _, coinRequest := range md.Assets {
			exchangeRate := rate.Rate
			percentChange := rate.PercentChange24h

			coinObj, ok := coin.Coins[coinRequest.Coin]
			if !ok {
				continue
			}
			r, err := storage.GetTicker(coinObj.Symbol, strings.ToUpper(coinRequest.TokenId))
			if err != nil {
				// TODO: either fail the request or signal to requester that ticker for coin couldn't be retrieved
				logger.Error(err, "Failed to retrieve ticker", logger.Params{"coin": coinObj.Symbol, "currency": md.Currency})
				continue
			}
			if r.Price.Currency != blockatlas.DefaultCurrency {
				newRate, err := storage.GetRate(strings.ToUpper(r.Price.Currency))
				if err == nil {
					exchangeRate *= newRate.Rate
					percentChange = newRate.PercentChange24h
				} else {
					tickerRate, err := storage.GetTicker(strings.ToUpper(r.Price.Currency), "")
					if err == nil {
						exchangeRate *= tickerRate.Price.Value
						percentChange = big.NewFloat(tickerRate.Price.Change24h)
					}
				}
			}

			r.ApplyRate(md.Currency, exchangeRate, percentChange)
			r.SetCoinId(coinRequest.Coin)
			tickers = append(tickers, r)
		}

		ginutils.RenderSuccess(c, blockatlas.TickerResponse{Currency: md.Currency, Docs: tickers})
	}
}

// @Summary Get charts data for a specific coin
// @Id get_charts_data
// @Description Get the charts data from an market and coin/token
// @Accept json
// @Produce json
// @Tags Market
// @Param coin query int true "Coin ID" default(60)
// @Param token query string false "Token ID"
// @Param time_start query int false "Start timestamp" default(1574483028)
// @Param max_items query int false "Max number of items in result prices array" default(64)
// @Param currency query string false "The currency to show charts" default(USD)
// @Success 200 {object} blockatlas.ChartData
// @Router /v1/market/charts [get]
func getChartsHandler() func(c *gin.Context) {
	var charts = market.InitCharts()
	return func(c *gin.Context) {
		coinQuery := c.Query("coin")
		coinId, err := strconv.Atoi(coinQuery)
		if err != nil {
			ginutils.RenderError(c, http.StatusBadRequest, "Invalid coin")
			return
		}
		token := c.Query("token")

		timeStart, err := strconv.ParseInt(c.Query("time_start"), 10, 64)
		if err != nil {
			ginutils.RenderError(c, http.StatusBadRequest, "Invalid time_start")
			return
		}
		maxItems, err := strconv.Atoi(c.Query("max_items"))
		if err != nil || maxItems <= 0 {
			maxItems = defaultMaxChartItems
		}

		currency := c.DefaultQuery("currency", blockatlas.DefaultCurrency)

		chart, err := charts.GetChartData(uint(coinId), token, currency, timeStart, maxItems)
		if err != nil {
			logger.Error(err, "Failed to retrieve chart", logger.Params{"coin": coinId, "currency": currency})
			ginutils.RenderError(c, http.StatusInternalServerError, err.Error())
			return
		}
		ginutils.RenderSuccess(c, chart)
	}
}

// @Summary Get charts coin info data for a specific coin
// @Id get_charts_coin_info
// @Description Get the charts coin info data from an market and coin/contract
// @Accept json
// @Produce json
// @Tags Market
// @Param coin query int true "Coin ID" default(60)
// @Param token query string false "Token ID"
// @Param time_start query int false "Start timestamp" default(1574483028)
// @Param currency query string false "The currency to show coin info in" default(USD)
// @Success 200 {object} blockatlas.ChartCoinInfo
// @Router /v1/market/info [get]
func getCoinInfoHandler() func(c *gin.Context) {
	var charts = market.InitCharts()
	return func(c *gin.Context) {
		coinQuery := c.Query("coin")
		coinId, err := strconv.Atoi(coinQuery)
		if err != nil {
			ginutils.RenderError(c, http.StatusBadRequest, "Invalid coin")
			return
		}

		token := c.Query("token")
		currency := c.DefaultQuery("currency", blockatlas.DefaultCurrency)
		chart, err := charts.GetCoinInfo(uint(coinId), token, currency)
		if err != nil {
			logger.Error(err, "Failed to retrieve coin info", logger.Params{"coin": coinId, "currency": currency})
			ginutils.RenderError(c, http.StatusInternalServerError, err.Error())
			return
		}
		chart.Info, _ = assets.GetCoinInfo(coinId, token)
		ginutils.RenderSuccess(c, chart)
	}
}
