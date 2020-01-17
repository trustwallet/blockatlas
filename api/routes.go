package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
)

var routers = make(map[string]gin.IRouter)

func LoadPlatforms(root gin.IRouter) {
	v1 := root.Group("/v1")
	v2 := root.Group("/v2")
	v3 := root.Group("/v3")

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
		routerv2 := getRouter(v2, collectionAPI.Coin().Handle)
		routerv3 := getRouter(v3, collectionAPI.Coin().Handle)

		makeCollectionsRoute(routerv3, collectionAPI)
		makeCollectionRoute(routerv3, collectionAPI)

		//TODO: remove once most of the clients will be updated (deadline: March 17th)
		oldMakeCollectionRoute(routerv2, collectionAPI)
		oldMakeCollectionsRoute(routerv2, collectionAPI)
	}

	for _, customAPI := range platform.CustomAPIs {
		router := getRouter(v1, customAPI.Coin().Handle)
		customAPI.RegisterRoutes(router)
	}

	{
		ns := root.Group("/ns")
		batchNs := v2.Group("/ns")
		MakeLookupRoute(ns)
		MakeLookupBatchRoute(batchNs)
	}

	//TODO: remove once most of the clients will be updated (deadline: March 17th)
	oldMakeCategoriesBatchRoute(v2)
	makeCategoriesBatchRoute(v3)
	makeStakingDelegationsBatchRoute(v2)
	makeStakingDelegationsSimpleBatchRoute(v2)

	makeChartsRoute(v1)
	makeCoinInfoRoute(v1)
	logger.Info("Routes set up", logger.Params{"routes": len(routers)})

	v1.GET("/", getEnabledEndpoints)
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
