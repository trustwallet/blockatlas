package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/platform"
)

var routers = make(map[string]gin.IRouter)

func loadPlatforms(root gin.IRouter) {
	v1 := root.Group("/v1")

	for _, txAPI := range platform.TxAPIs {
		router := getRouter(v1, txAPI.Coin().Handle)
		makeTxRoute(router, txAPI)
	}
	for _, customAPI := range platform.CustomAPIs {
		router := getRouter(v1, customAPI.Coin().Handle)
		customAPI.RegisterRoutes(router)
	}

	logrus.WithField("routes", len(routers)).
		Info("Routes set up")

	v1.GET("/", getEnabledEndpoints)
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
