package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/marketdata"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/storage"
	"net/http"
	"strconv"
	"strings"
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
	router.GET("/ticker", getTickerHandler(db))
	router.POST("/ticker", getTickersHandler(db))
}

// @Summary Get ticker value for a specific market
// @Id get_ticker
// @Description Get the ticker value from an market and coin/token
// @Accept json
// @Produce json
// @Tags ticker
// @Param coin query int true "coin id"
// @Param token query string false "token id"
// @Param currency query string false "the currency to show the quote" default(USD)
// @Success 200 {object} blockatlas.Ticker
// @Router /market/v1/ticker [get]
func getTickerHandler(storage storage.Market) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		coinQuery := c.Query("coin")
		coinId, err := strconv.Atoi(coinQuery)
		if err != nil {
			ginutils.RenderError(c, http.StatusInternalServerError, "Invalid coin")
			return
		}
		token := c.Query("token")

		currency := c.DefaultQuery("currency", blockatlas.DefaultCurrency)
		rate, err := storage.GetRate(strings.ToUpper(currency))
		if err != nil {
			ginutils.RenderError(c, http.StatusInternalServerError, "Invalid currency")
			return
		}

		coinObj := coin.Coins[uint(coinId)]
		result, err := storage.GetTicker(coinObj.Symbol, strings.ToUpper(token))
		if err != nil {
			ginutils.RenderError(c, http.StatusInternalServerError, err.Error())
			return
		}
		result.ApplyRate(rate.Rate, currency)
		ginutils.RenderSuccess(c, result)
	}
}

// @Summary Get ticker values for a specific markets
// @Id get_tickers
// @Description Get the ticker values from many markets and coin/token
// @Accept json
// @Produce json
// @Tags ticker
// @Param tickers body api.TickerRequest true "Ticker"
// @Success 200 {object} blockatlas.Tickers
// @Router /market/v1/tickers [post]
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
			ginutils.RenderError(c, http.StatusInternalServerError, "Invalid currency")
			return
		}

		tickers := make(blockatlas.Tickers, 0)
		for _, coinRequest := range md.Assets {
			coinObj, ok := coin.Coins[coinRequest.Coin]
			if !ok {
				continue
			}
			r, err := storage.GetTicker(coinObj.Symbol, strings.ToUpper(coinRequest.TokenId))
			if err != nil {
				continue
			}
			r.ApplyRate(rate.Rate, md.Currency)
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
// @Tags charts
// @Param coin query int true "coin id"
// @Param token query string false "token id"
// @Param currency query string false "the currency to show charts" default(USD)
// @Success 200 {object} blockatlas.ChartData
// @Router /v1/market/charts [get]
func makeChartsRoute(router gin.IRouter) {
	var charts = marketdata.InitCharts()
	router.GET("/market/charts", func(c *gin.Context) {
		coinQuery := c.Query("coin")
		coinId, err := strconv.Atoi(coinQuery)
		if err != nil {
			ginutils.RenderError(c, http.StatusInternalServerError, "Invalid coin")
			return
		}
		token := c.Query("token")

		timeStart, err := strconv.ParseInt(c.Query("timeStart"), 10, 64)
		if err != nil {
			ginutils.RenderError(c, http.StatusInternalServerError, "Invalid timeStart")
			return
		}

		currency := c.DefaultQuery("currency", blockatlas.DefaultCurrency)

		chart, err := charts.GetChartData(uint(coinId), token, currency, timeStart)
		if err != nil {
			ginutils.RenderError(c, http.StatusInternalServerError, err.Error())
			return
		}
		ginutils.RenderSuccess(c, chart)
	})
}

// @Summary Get charts coin info data for a specific coin
// @Id get_charts_coin_info
// @Description Get the charts coin info data from an market and coin/contract
// @Accept json
// @Produce json
// @Tags charts
// @Param coin query int true "coin id"
// @Param token query string false "token id"
// @Param currency query string false "the currency to show coin info in" default(USD)
// @Success 200 {object} blockatlas.ChartCoinInfo
// @Router /v1/market/info [get]
func makeCoinInfoRoute(router gin.IRouter) {
	var charts = marketdata.InitCharts()
	router.GET("/market/info", func(c *gin.Context) {
		coinQuery := c.Query("coin")
		coinId, err := strconv.Atoi(coinQuery)
		if err != nil {
			ginutils.RenderError(c, http.StatusInternalServerError, "Invalid coin")
			return
		}
		token := c.Query("token")

		currency := c.DefaultQuery("currency", blockatlas.DefaultCurrency)

		chart, err := charts.GetCoinInfo(uint(coinId), token, currency)
		if err != nil {
			ginutils.RenderError(c, http.StatusInternalServerError, err.Error())
			return
		}
		ginutils.RenderSuccess(c, chart)
	})
}
