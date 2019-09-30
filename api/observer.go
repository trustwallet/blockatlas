package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/observer"
	observerStorage "github.com/trustwallet/blockatlas/observer/storage"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/platform/bitcoin"
	"net/http"
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
	router.Use(requireAuth)
	router.POST("/webhook/register", addCall)
	router.DELETE("/webhook/register", deleteCall)
	router.GET("/status", statusCall)
}

func requireAuth(c *gin.Context) {
	auth := fmt.Sprintf("Bearer %s", viper.GetString("observer.auth"))
	if c.GetHeader("Authorization") == auth {
		c.Next()
	} else {
		ErrorResponse(c).Code(http.StatusUnauthorized).Render()
	}
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
		RenderSuccess(c, ObserverResponse{Status: "Added"})
		return
	}

	var subs []observer.Subscription
	for coinStr, perCoin := range req.Subscriptions {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}
		for _, addr := range perCoin {
			subs = append(subs, observer.Subscription{
				Coin:     uint(coin),
				Address:  addr,
				Webhooks: []string{req.Webhook},
			})
		}
	}

	var xpubSubs []observer.Subscription
	for coinStr, perCoin := range req.XpubSubscriptions {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}

		for _, xpub := range perCoin {
			xpubSubs = append(xpubSubs, observer.Subscription{
				Coin:     uint(coin),
				Address:  xpub,
				Webhooks: []string{req.Webhook},
			})
			go cacheXPubAddress(xpub, uint(coin))
		}
	}
	subs = append(subs, xpubSubs...)
	err := observerStorage.App.Add(subs)
	if err != nil {
		ErrorResponse(c).Message(err.Error()).Render()
		return
	}

	RenderSuccess(c, ObserverResponse{Status: "Added"})
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
		RenderSuccess(c, ObserverResponse{Status: "Deleted"})
		return
	}

	var subs []observer.Subscription
	for coinStr, perCoin := range req.Subscriptions {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}
		for _, addr := range perCoin {
			subs = append(subs, observer.Subscription{
				Coin:     uint(coin),
				Address:  addr,
				Webhooks: []string{req.Webhook},
			})
		}
	}

	var xpubSubs []observer.Subscription
	for coinStr, perCoin := range req.XpubSubscriptions {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}
		for _, addr := range perCoin {
			xpubSubs = append(xpubSubs, observer.Subscription{
				Coin:     uint(coin),
				Address:  addr,
				Webhooks: []string{req.Webhook},
			})
		}
	}

	subs = append(subs, xpubSubs...)
	err := observerStorage.App.Delete(subs)
	if err != nil {
		ErrorResponse(c).Message(err.Error()).Render()
		return
	}

	RenderSuccess(c, ObserverResponse{Status: "Deleted"})
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
	RenderSuccess(c, result)
}
