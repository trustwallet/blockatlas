package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/platform"
	"net/http"

	"github.com/gin-gonic/gin"
)

var routers = make(map[string]gin.IRouter)

func loadPlatforms(root gin.IRouter) {
	v1 := root.Group("/v1")

	for _, txAPI := range platform.TxAPIs {
		router := getRouter(v1, txAPI.Coin().Handle)
		makeGenericAPI(router, txAPI)
	}
	for _, customAPI := range platform.CustomAPIs {
		router := getRouter(v1, customAPI.Coin().Handle)
		customAPI.RegisterRoutes(router)
	}

	logrus.WithField("routes", len(routers)).
		Info("Routes set up")

	v1.GET("/", getEnabledEndpoints)
}

func setupEmpty(router gin.IRouter) {
	var emptyPage blockatlas.TxPage
	emptyResponse, _ := emptyPage.MarshalJSON()
	mkEmpty := func(c *gin.Context) {
		c.Writer.Header().Set("content-type", "application/json")
		c.Writer.WriteHeader(http.StatusOK)
		_, _ = c.Writer.Write(emptyResponse)
	}
	router.GET("/:address", mkEmpty)
	router.GET("/:address/transactions/:token", mkEmpty)
}

// getRouter lazy loads routers
func getRouter(router gin.IRouter, handle string) gin.IRouter {
	if group, ok := routers[handle]; ok {
		return group
	} else {
		path := fmt.Sprintf("/%s", handle)
		logrus.Debugf("Registering %s", path)
		group := router.Group(path)
		routers[handle] = group
		return group
	}
}
