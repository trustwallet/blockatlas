package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/observer/storage"
	"github.com/trustwallet/blockatlas/platform"
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

func SetupObserverAPI(router gin.IRouter, db *storage.Storage) {
	router.Use(requireAuth)
	router.POST("/webhook/register", addCall(db))
	router.DELETE("/webhook/register", deleteCall(db))
	router.GET("/status", statusCall(db))
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
func addCall(storage *storage.Storage) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		var req Webhook
		if c.BindJSON(&req) != nil {
			return
		}

		if len(req.Subscriptions) == 0 && len(req.XpubSubscriptions) == 0 {
			RenderSuccess(c, ObserverResponse{Status: "Added"})
			return
		}

		subs := parseSubscriptions(req.Subscriptions, req.Webhook)
		xpubSubs := parseSubscriptions(req.XpubSubscriptions, req.Webhook)
		subs = append(subs, xpubSubs...)

		go storage.AddSubscriptions(subs)
		go cacheXpub(req.XpubSubscriptions, storage)
		RenderSuccess(c, ObserverResponse{Status: "Added"})
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
func deleteCall(storage *storage.Storage) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		var req Webhook
		if c.BindJSON(&req) != nil {
			return
		}

		if len(req.Subscriptions) == 0 && len(req.XpubSubscriptions) == 0 {
			RenderSuccess(c, ObserverResponse{Status: "Deleted"})
			return
		}

		subs := parseSubscriptions(req.Subscriptions, req.Webhook)
		xpubSubs := parseSubscriptions(req.XpubSubscriptions, req.Webhook)
		subs = append(subs, xpubSubs...)

		go storage.DeleteSubscriptions(subs)
		RenderSuccess(c, ObserverResponse{Status: "Deleted"})
	}
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
func statusCall(storage *storage.Storage) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		result := make(map[string]CoinStatus)
		for _, api := range platform.BlockAPIs {
			coin := api.Coin()
			num, err := storage.GetBlockNumber(coin.ID)
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
}

func cacheXpub(subscriptions map[string][]string, storage *storage.Storage) {
	for coinStr, perCoin := range subscriptions {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}
		for _, xpub := range perCoin {
			go storage.CacheXPubAddress(xpub, uint(coin))
		}
	}
}

func parseSubscriptions(subscriptions map[string][]string, webhook string) (subs []interface{}) {
	for coinStr, perCoin := range subscriptions {
		coin, err := strconv.Atoi(coinStr)
		if err != nil {
			continue
		}
		for _, addr := range perCoin {
			subs = append(subs, &storage.Subscription{
				Coin:    coin,
				Address: addr,
				Webhook: webhook,
			})
		}
	}
	return
}
