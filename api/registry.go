package api

import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/trustwallet/blockatlas/api/endpoint"
	"github.com/trustwallet/blockatlas/api/middleware"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform"
	"time"
)

func RegisterCollectionsAPI(router gin.IRouter, api blockatlas.Platform) {
	collectionAPI, ok := api.(blockatlas.CollectionAPI)
	if !ok {
		return
	}

	handle := collectionAPI.Coin().Handle

	router.GET("/v3/"+handle+"/collections/:owner/collection/:collection_id", func(c *gin.Context) {
		endpoint.GetCollectiblesForSpecificCollectionAndOwnerV3(c, collectionAPI)
	})
	router.GET("/v3/"+handle+"/collections/:owner", func(c *gin.Context) {
		endpoint.GetCollectiblesForOwnerV3(c, collectionAPI)
	})
	router.GET("/v4/"+handle+"/collections/:owner/collection/:collection_id", func(c *gin.Context) {
		endpoint.GetCollectiblesForSpecificCollectionAndOwner(c, collectionAPI)
	})
}

func RegisterTokensAPI(router gin.IRouter, api blockatlas.Platform) {
	tokenAPI, ok := api.(blockatlas.TokenAPI)
	if !ok {
		return
	}

	handle := api.Coin().Handle

	router.GET("/v2/"+handle+"/tokens/:address", func(c *gin.Context) {
		endpoint.GetTokensByAddress(c, tokenAPI)
	})
}

func RegisterTransactionsAPI(router gin.IRouter, api blockatlas.Platform) {
	txAPI, _ := api.(blockatlas.TxAPI)
	tokenTxAPI, _ := api.(blockatlas.TokenTxAPI)

	handle := api.Coin().Handle

	if IsForCustomAPI(handle) {
		return
	}

	router.GET("/v1/"+handle+"/:address", func(c *gin.Context) {
		endpoint.GetTransactionsHistory(c, txAPI, tokenTxAPI)
	})
	router.GET("/v2/"+handle+"/transactions/:address", func(c *gin.Context) {
		endpoint.GetTransactionsHistory(c, txAPI, tokenTxAPI)
	})
}

func RegisterStakeAPI(router gin.IRouter, api blockatlas.Platform) {
	stakingAPI, ok := api.(blockatlas.StakeAPI)
	if !ok {
		return
	}
	handle := api.Coin().Handle

	router.GET("/v2/"+handle+"/staking/validators", middleware.CacheMiddleware(time.Hour, func(c *gin.Context) {
		endpoint.GetValidators(c, stakingAPI)
	}))
	router.GET("/v2/"+handle+"/staking/delegations/:address", func(c *gin.Context) {
		endpoint.GetStakingDelegationsForSpecificCoin(c, stakingAPI)
	})
}

func RegisterBatchAPI(router gin.IRouter) {
	router.POST("v2/staking/delegations", func(c *gin.Context) {
		endpoint.GetStakeDelegationsWithAllInfoForBatch(c, platform.StakeAPIs)
	})
	router.POST("v2/staking/list", middleware.CacheMiddleware(time.Hour, func(c *gin.Context) {
		endpoint.GetStakeInfoForBatch(c, platform.StakeAPIs)
	}))
	router.POST("/v3/collectibles/categories", func(c *gin.Context) {
		endpoint.GetCollectionCategoriesFromListV3(c, platform.CollectionAPIs)
	})
	router.POST("/v4/collectibles/categories", func(c *gin.Context) {
		endpoint.GetCollectionCategoriesFromList(c, platform.CollectionAPIs)
	})
	router.POST("/v2/tokens", func(c *gin.Context) {
		endpoint.GetTokens(c, platform.TokensAPIs)
	})
}

// CustomAPI must be removed and all handlers needs to be migrated to the transactions, tokens api
func RegisterCustomAPI(router gin.IRouter, api blockatlas.Platform) {
	customAPI, ok := api.(blockatlas.CustomAPI)
	if !ok {
		return
	}
	handle := api.Coin().Handle

	customRouter := router.Group("/v1/" + handle)
	customAPI.RegisterRoutes(customRouter)
}

func RegisterDomainAPI(router gin.IRouter) {
	router.GET("/ns/lookup", endpoint.GetAddressByCoinAndDomain)
	router.GET("v2/ns/lookup", endpoint.GetAddressByCoinAndDomainBatch)
}

func RegisterBasicAPI(router gin.IRouter) {
	router.GET("/", endpoint.GetStatus)
	router.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
}

func IsForCustomAPI(handle string) bool {
	switch handle {
	case
		coin.Bitcoin().Handle,
		coin.Litecoin().Handle,
		coin.Bitcoincash().Handle,
		coin.Zcash().Handle,
		coin.Zcoin().Handle,
		coin.Viacoin().Handle,
		coin.Ravencoin().Handle,
		coin.Groestlcoin().Handle,
		coin.Zelcash().Handle,
		coin.Decred().Handle,
		coin.Digibyte().Handle,
		coin.Dash().Handle,
		coin.Doge().Handle,
		coin.Qtum().Handle:
		return true
	default:
		return false
	}
}
