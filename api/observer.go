package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/observer"
	observerStorage "github.com/trustwallet/blockatlas/observer/storage"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"strconv"
)

type ObserverResponse struct {
	Status string `json:"status"`
}

type Webhook struct {
	Subscriptions     map[string][]string `json:"subscriptions"`
	XpubSubscriptions map[string][]string `json:"xpub_subscriptions"`
	Webhook           string              `json:"webhook"`
}

type CoinStatus struct {
	Height int64  `json:"height"`
	Error  string `json:"error,omitempty"`
}

func SetupObserverAPI(router gin.IRouter) {
	router.Use(ginutils.TokenAuthMiddleware(viper.GetString("observer.auth")))
	router.POST("/webhook/register", addCall)
	router.DELETE("/webhook/register", deleteCall)
	router.GET("/status", statusCall)
}

// @Summary Create a webhook
// @ID create_webhook
// @Description Create a webhook for addresses transactions
// @Accept json
// @Produce json
// @Tags observer,subscriptions
// @Param subscriptions body api.Webhook true "Accounts subscriptions"
// @Param Authorization header string true "Bearer authorization header" default(Bearer test)
// @Header 200 {string} Authorization {token}
// @Success 200 {object} api.ObserverResponse
// @Router /observer/v1/webhook/register [post]
func addCall(c *gin.Context) {
	var req Webhook
	if c.BindJSON(&req) != nil {
		return
	}

	if len(req.Subscriptions) == 0 && len(req.XpubSubscriptions) == 0 {
		ginutils.RenderSuccess(c, ObserverResponse{Status: "Added"})
		return
	}

	subs := parseSubscriptions(req.Subscriptions, req.Webhook)
	xpubSubs := parseSubscriptions(req.XpubSubscriptions, req.Webhook)
	subs = append(subs, xpubSubs...)
	err := observerStorage.App.Add(subs)
	if err != nil {
		ginutils.ErrorResponse(c).Message(err.Error()).Render()
		return
	}
	go cacheXpub(req.XpubSubscriptions)
	ginutils.RenderSuccess(c, ObserverResponse{Status: "Added"})
}

func cacheXPubAddress(xpub string, coin uint) {
	platform := bitcoin.UtxoPlatform(coin)
	addresses, err := platform.GetAddressesFromXpub(xpub)
	if err != nil || len(addresses) == 0 {
		logger.Error("GetAddressesFromXpub", err, logger.Params{
			"xpub":      xpub,
			"coin":      coin,
			"addresses": addresses,
		})
		return
	}
	err = observerStorage.App.SaveXpubAddresses(coin, addresses, xpub)
	if err != nil {
		logger.Error("SaveXpubAddresses", err, logger.Params{
			"xpub": xpub,
			"coin": coin,
		})
	}
}

// @Summary Delete a webhook
// @ID delete_webhook
// @Description Delete a webhook for addresses transactions
// @Accept json
// @Produce json
// @Tags observer,subscriptions
// @Param subscriptions body api.Webhook true "Accounts subscriptions"
// @Param Authorization header string true "Bearer authorization header" default(Bearer test)
// @Header 200 {string} Authorization {token}
// @Success 200 {object} api.ObserverResponse
// @Router /observer/v1/webhook/register [delete]
func deleteCall(c *gin.Context) {
	var req Webhook
	if c.BindJSON(&req) != nil {
		return
	}

	if len(req.Subscriptions) == 0 && len(req.XpubSubscriptions) == 0 {
		ginutils.RenderSuccess(c, ObserverResponse{Status: "Deleted"})
		return
	}

	subs := parseSubscriptions(req.Subscriptions, req.Webhook)
	xpubSubs := parseSubscriptions(req.XpubSubscriptions, req.Webhook)
	subs = append(subs, xpubSubs...)

	err := observerStorage.App.Delete(subs)
	if err != nil {
		ginutils.ErrorResponse(c).Message(err.Error()).Render()
		return
	}

	ginutils.RenderSuccess(c, ObserverResponse{Status: "Deleted"})
}

// @Summary Get coin status
// @ID coin_status
// @Description Get coin status
// @Accept json
// @Produce json
// @Tags observer,subscriptions
// @Param Authorization header string true "Bearer authorization header" default(Bearer test)
// @Header 200 {string} Authorization {token}
// @Success 200 {object} api.CoinStatus
// @Router /observer/v1/status [get]
func statusCall(c *gin.Context) {
	result := make(map[string]CoinStatus)
	for _, api := range platform.BlockAPIs {
		coin := api.Coin()
		num, err := observerStorage.App.GetBlockNumber(coin.ID)
		var status CoinStatus
		if err != nil {
			status = CoinStatus{Error: err.Error()}
		} else if num == 0 {
			status = CoinStatus{Error: "no blocks"}
		} else {
			status = CoinStatus{Height: num}
		}
		result[coin.Handle] = status
	}
	ginutils.RenderSuccess(c, result)
}

func cacheXpub(subscriptions map[string][]string) {
	for coinStr, perCoin := range subscriptions {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}
		for _, xpub := range perCoin {
			go cacheXPubAddress(xpub, uint(coin))
		}
	}
}

func parseSubscriptions(subscriptions map[string][]string, webhook string) (subs []observer.Subscription) {
	for coinStr, perCoin := range subscriptions {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}
		for _, addr := range perCoin {
			subs = append(subs, observer.Subscription{
				Coin:     uint(coin),
				Address:  addr,
				Webhooks: []string{webhook},
			})
		}
	}
	return
}
