package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/observer"
	observerStorage "github.com/trustwallet/blockatlas/observer/storage"
	"net/http"
)

func setupObserverAPI(router gin.IRouter) {
	router.Use(requireAuth)
	router.POST("/", addObserver)
}

func requireAuth(c *gin.Context) {
	auth := fmt.Sprintf("Bearer %s", viper.GetString("observer.auth"))
	if c.GetHeader("Authorization") == auth {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func addObserver(c *gin.Context) {
	var req struct {
		Subscriptions map[uint][]string `json:"subscriptions"`
		Webhook string `json:"webhook"`
	}
	if c.BindJSON(&req) != nil {
		return
	}

	if len(req.Subscriptions) == 0 {
		c.String(http.StatusOK, "Added")
		return
	}

	var subs []observer.Subscription
	for coin, perCoin := range req.Subscriptions {
		for _, addr := range perCoin {
			subs = append(subs, observer.Subscription{
				Coin:    uint(coin),
				Address: addr,
				Webhook: req.Webhook,
			})
		}
	}

	err := observerStorage.App.Add(subs)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.String(http.StatusOK, "Added")
}
