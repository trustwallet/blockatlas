package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/storage"
)

func SetupObserverAPI(router gin.IRouter, db *storage.Storage) {
	router.Use(ginutils.TokenAuthMiddleware(viper.GetString("observer.auth")))
	router.POST("/webhook/register", addCall(db))
	router.DELETE("/webhook/register", deleteCall(db))
	router.GET("/status", statusCall(db))
}

// @Summary Create a webhook
// @ID create_webhook
// @Description Create a webhook for addresses transactions
// @Accept json
// @Produce json
// @Tags Observer
// @Param subscriptions body blockatlas.Webhook true "Accounts subscriptions"
// @Param Authorization header string true "Bearer authorization header" default(Bearer test)
// @Header 200 {string} Authorization {token}
// @Success 200 {object} blockatlas.Observer
// @Router /observer/v1/webhook/register [post]
func addCall(storage storage.Addresses) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		var req blockatlas.Webhook
		if c.BindJSON(&req) != nil {
			return
		}

		if len(req.Subscriptions) == 0 {
			ginutils.RenderSuccess(c, blockatlas.Observer{Message: "Added", Status: true})
			return
		}
		subs := req.ParseSubscriptions()
		go storage.AddSubscriptions(subs)

		ginutils.RenderSuccess(c, blockatlas.Observer{Message: "Added", Status: true})
	}
}

// @Summary Delete a webhook
// @ID delete_webhook
// @Description Delete a webhook for addresses transactions
// @Accept json
// @Produce json
// @Tags Observer
// @Param subscriptions body blockatlas.Webhook true "Accounts subscriptions"
// @Param Authorization header string true "Bearer authorization header" default(Bearer test)
// @Header 200 {string} Authorization {token}
// @Success 200 {object} blockatlas.Observer
// @Router /observer/v1/webhook/register [delete]
func deleteCall(storage storage.Addresses) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		var req blockatlas.Webhook
		if c.BindJSON(&req) != nil {
			return
		}

		if len(req.Subscriptions) == 0 {
			ginutils.RenderSuccess(c, blockatlas.Observer{Message: "Deleted", Status: true})
			return
		}

		subs := req.ParseSubscriptions()
		go storage.DeleteSubscriptions(subs)
		ginutils.RenderSuccess(c, blockatlas.Observer{Message: "Deleted", Status: true})
	}
}

// @Summary Get coin status
// @ID coin_status
// @Description Get coin status
// @Accept json
// @Produce json
// @Tags Observer
// @Param Authorization header string true "Bearer authorization header" default(Bearer test)
// @Header 200 {string} Authorization {token}
// @Success 200 {object} blockatlas.CoinStatus
// @Router /observer/v1/status [get]
func statusCall(storage storage.Tracker) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		result := make(map[string]blockatlas.CoinStatus)
		for _, api := range platform.BlockAPIs {
			coin := api.Coin()
			num, err := storage.GetBlockNumber(coin.ID)
			var status blockatlas.CoinStatus
			if err != nil {
				status = blockatlas.CoinStatus{Error: err.Error()}
			} else if num == 0 {
				status = blockatlas.CoinStatus{Error: "no blocks"}
			} else {
				status = blockatlas.CoinStatus{Height: num}
			}
			result[coin.Handle] = status
		}
		ginutils.RenderSuccess(c, result)
	}
}
