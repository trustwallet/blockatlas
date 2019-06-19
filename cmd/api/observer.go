package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/observer"
	observerStorage "github.com/trustwallet/blockatlas/observer/storage"
	"net/http"
	"strconv"
)

func setupObserverAPI(router gin.IRouter) {
	router.Use(requireAuth)
	router.POST("/:coin", addObserver)
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
	coinStr := c.Param("coin")
	coin, err := strconv.Atoi(coinStr)
	if err != nil {
		c.String(http.StatusNotFound, "404 page not found")
		return
	}

	var req struct {
		Address []string `json:"addresses"`
		Webhook string `json:"webhook"`
	}
	if c.BindJSON(&req) != nil {
		return
	}

	if len(req.Address) == 0 {
		c.String(http.StatusOK, "Added")
		return
	}

	subs := make([]observer.Subscription, len(req.Address))

	for i, addr := range req.Address {
		subs[i] = observer.Subscription{
			Coin: uint(coin),
			Address: addr,
			Webhook: req.Webhook,
		}
	}

	err = observerStorage.App.Add(subs)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.String(http.StatusOK, "Added")
}
