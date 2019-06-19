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
	router.DELETE("/:coin/:address", removeObserver)
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
		Address string `form:"address" binding:"required"`
		Webhook string `form:"webhook" binding:"required"`
	}
	if c.Bind(&req) != nil {
		return
	}

	sub := observer.Subscription{
		Coin:    uint(coin),
		Address: req.Address,
		Webhook: req.Webhook,
	}

	err = observerStorage.App.Add(sub)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(200, &sub)
}

func removeObserver(c *gin.Context) {
	coinStr := c.Param("coin")
	coin, err := strconv.Atoi(coinStr)
	if err != nil {
		c.String(http.StatusNotFound, "404 page not found")
		return
	}
	address := c.Param("address")

	err = observerStorage.App.Remove(uint(coin), address)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.String(http.StatusOK, "Removed")
}
