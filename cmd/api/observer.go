package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/observer"
	observerStorage "github.com/trustwallet/blockatlas/observer/storage"
	"github.com/trustwallet/blockatlas/platform"
	"net/http"
	"strconv"
)

type setParams struct {
	Subscriptions map[string][]string `json:"subscriptions"`
	Webhook       string              `json:"webhook"`
}

func setupObserverAPI(router gin.IRouter) {
	router.Use(requireAuth)
	router.POST("/webhook/register", addCall)
	router.POST("/webhook/delete", deleteCall)
	router.GET("/status", statusCall)
}

func requireAuth(c *gin.Context) {
	auth := fmt.Sprintf("Bearer %s", viper.GetString("observer.auth"))
	if c.GetHeader("Authorization") == auth {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func (s *setParams) ToSubscriptions() (subs []observer.Subscription) {
	for coinStr, perCoin := range s.Subscriptions {
		coin, _ := strconv.Atoi(coinStr)
		if coin == 0 {
			continue
		}
		for _, addr := range perCoin {
			subs = append(subs, observer.Subscription{
				Coin:    uint(coin),
				Address: addr,
				WebHook: s.Webhook,
			})
		}
	}
	return
}

func addCall(c *gin.Context) {
	var req setParams
	if c.BindJSON(&req) != nil {
		return
	}

	if len(req.Subscriptions) == 0 {
		c.String(http.StatusOK, "Added")
		return
	}

	err := observerStorage.App.Add(req.ToSubscriptions())
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, "Added")
}

func deleteCall(c *gin.Context) {
	var req setParams
	if c.BindJSON(&req) != nil {
		return
	}

	if len(req.Subscriptions) == 0 {
		c.String(http.StatusOK, "Added")
		return
	}

	err := observerStorage.App.Delete(req.ToSubscriptions())
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusOK, "Deleted")
}

func statusCall(c *gin.Context) {
	type coinStatus struct {
		Height int64  `json:"height"`
		Error  string `json:"error,omitempty"`
	}

	result := make(map[string]coinStatus)

	for _, api := range platform.BlockAPIs {
		coin := api.Coin()
		num, err := observerStorage.App.GetBlockNumber(coin.ID)
		var status coinStatus
		if err != nil {
			status = coinStatus{Error: err.Error()}
		} else if num == 0 {
			status = coinStatus{Error: "no blocks"}
		} else {
			status = coinStatus{Height: num}
		}
		result[coin.Handle] = status
	}

	c.JSON(http.StatusOK, result)
}
