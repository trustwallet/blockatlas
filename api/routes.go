package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
)

var routers = make(map[string]gin.IRouter)

func SetupPlatformAPI(root gin.IRouter) {
	root.GET("/", GetRoot)
	root.GET("/status", GetStatus)

	v1 := root.Group("/v1")
	v3 := root.Group("/v3")
	v2 := root.Group("/v2")
	v4 := root.Group("/v4")

	v1.GET("/", GetSupportedEndpoints)

	for _, txAPI := range platform.Platforms {
		router := getRouter(v1, txAPI.Coin().Handle)
		makeTxRouteV1(router, txAPI)

		routerV2 := getRouter(v2, txAPI.Coin().Handle)
		makeTxRouteV2(routerV2, txAPI)
	}

	for _, tokenAPI := range platform.Platforms {
		router := getRouter(v2, tokenAPI.Coin().Handle)
		makeTokenRoute(router, tokenAPI)
	}

	for _, stakeAPI := range platform.Platforms {
		router := getRouter(v2, stakeAPI.Coin().Handle)
		makeStakingValidatorsRoute(router, stakeAPI)
		makeStakingDelegationsRoute(router, stakeAPI)
	}

	for _, collectionAPI := range platform.Platforms {

		routerV3 := getRouter(v3, collectionAPI.Coin().Handle)
		makeCollectionsRoute(routerV3, collectionAPI)
		makeCollectionRoute(routerV3, collectionAPI)

		routerV4 := getRouter(v4, collectionAPI.Coin().Handle)
		makeCollectionRouteV4(routerV4, collectionAPI)
	}

	for _, customAPI := range platform.CustomAPIs {
		router := getRouter(v1, customAPI.Coin().Handle)
		customAPI.RegisterRoutes(router)
	}

	ns := root.Group("/ns")
	batchNs := v2.Group("/ns")
	MakeLookupRoute(ns)
	MakeLookupBatchRoute(batchNs)

	makeStakingDelegationsBatchRoute(v2)
	makeStakingDelegationsSimpleBatchRoute(v2)

	makeCategoriesBatchRoute(v3)
	makeCategoriesBatchRouteV4(v4)

	logger.Info("Routes set up", logger.Params{"routes": len(routers)})
}

// getRouter lazy loads routers
func getRouter(router *gin.RouterGroup, handle string) gin.IRouter {
	key := fmt.Sprintf("%s/%s", router.BasePath(), handle)
	if group, ok := routers[key]; ok {
		return group
	} else {
		path := fmt.Sprintf("/%s", handle)
		logger.Debug("Registering route", logger.Params{"path": len(path)})
		group := router.Group(path)
		routers[key] = group
		return group
	}
}

func GetSupportedEndpoints(c *gin.Context) {
	var resp struct {
		Endpoints []string `json:"endpoints,omitempty"`
	}
	for handle := range routers {
		resp.Endpoints = append(resp.Endpoints, handle)
	}
	ginutils.RenderSuccess(c, &resp)
}

func GetRoot(c *gin.Context) { ginutils.RenderSuccess(c, "Welcome to the Block Atlas API") }

func GetStatus(c *gin.Context) {
	ginutils.RenderSuccess(c, map[string]interface{}{
		"status": true,
		"build":  internal.Build,
		"date":   internal.Date,
	})
}
