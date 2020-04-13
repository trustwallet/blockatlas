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

func RegisterCollectionsAPI(root gin.IRouter, api blockatlas.Platform) {
	collectionAPI, ok := api.(blockatlas.CollectionAPI)
	if !ok {
		return
	}

	handle := collectionAPI.Coin().Handle

	root.GET("/v3/"+handle+"/collections/:owner/collection/:collection_id", func(c *gin.Context) {
		endpoint.GetCollectiblesForSpecificCollectionAndOwnerV3(c, collectionAPI)
	})
	root.GET("/v3/"+handle+"/collections/:owner", func(c *gin.Context) {
		endpoint.GetCollectiblesForOwnerV3(c, collectionAPI)
	})
	root.GET("/v4/"+handle+"/collections/:owner/collection/:collection_id", func(c *gin.Context) {
		endpoint.GetCollectiblesForSpecificCollectionAndOwner(c, collectionAPI)
	})
}

func RegisterTokensAPI(root gin.IRouter, api blockatlas.Platform) {
	tokenAPI, ok := api.(blockatlas.TokenAPI)
	if !ok {
		return
	}

	handle := api.Coin().Handle

	root.GET("/v2/"+handle+"/tokens/:address", func(c *gin.Context) {
		endpoint.GetTokensByAddress(c, tokenAPI)
	})
}

func RegisterTransactionsAPI(root gin.IRouter, api blockatlas.Platform) {
	txAPI, _ := api.(blockatlas.TxAPI)
	tokenTxAPI, _ := api.(blockatlas.TokenTxAPI)

	handle := api.Coin().Handle

	if IsForCustomAPI(handle) {
		return
	}

	root.GET("/v1/"+handle+"/:address", func(c *gin.Context) {
		endpoint.GetTransactionsHistory(c, txAPI, tokenTxAPI)
	})
	root.GET("/v2/"+handle+"/transactions/:address", func(c *gin.Context) {
		endpoint.GetTransactionsHistory(c, txAPI, tokenTxAPI)
	})
}

func RegisterStakeAPI(root gin.IRouter, api blockatlas.Platform) {
	stakingAPI, ok := api.(blockatlas.StakeAPI)
	if !ok {
		return
	}
	handle := api.Coin().Handle

	root.GET("/v2/"+handle+"/staking/validators", middleware.CacheMiddleware(time.Hour, func(c *gin.Context) {
		endpoint.GetValidators(c, stakingAPI)
	}))
	root.GET("/v2/"+handle+"/staking/delegations/:address", func(c *gin.Context) {
		endpoint.GetStakingDelegationsForSpecificCoin(c, stakingAPI)
	})
}

func RegisterBatchAPI(root gin.IRouter) {
	root.POST("v2/staking/delegations", func(c *gin.Context) {
		endpoint.GetStakeDelegationsWithAllInfoForBatch(c, platform.StakeAPIs)
	})
	root.POST("v2/staking/list", middleware.CacheMiddleware(time.Hour, func(c *gin.Context) {
		endpoint.GetStakeInfoForBatch(c, platform.StakeAPIs)
	}))
	root.POST("/v3/collectibles/categories", func(c *gin.Context) {
		endpoint.GetCollectionCategoriesFromListV3(c, platform.CollectionAPIs)
	})
	root.POST("/v4/collectibles/categories", func(c *gin.Context) {
		endpoint.GetCollectionCategoriesFromList(c, platform.CollectionAPIs)
	})
}

// CustomAPI must be removed and all handlers needs to be migrated to the transactions, tokens api
func RegisterCustomAPI(root gin.IRouter, api blockatlas.Platform) {
	customAPI, ok := api.(blockatlas.CustomAPI)
	if !ok {
		return
	}
	handle := api.Coin().Handle

	customRouter := root.Group("/v1/" + handle)
	customAPI.RegisterRoutes(customRouter)
}

func RegisterDomainAPI(root gin.IRouter) {
	root.GET("/ns/lookup", endpoint.GetAddressByCoinAndDomain)
	root.GET("v2/ns/lookup", endpoint.GetAddressByCoinAndDomainBatch)
}

func RegisterBasicAPI(root gin.IRouter) {
	root.GET("/", endpoint.GetStatus)
	root.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
}

func IsForCustomAPI(handle string) bool {
	switch handle {
	case coin.Gochain().Handle,
		coin.Thundertoken().Handle,
		coin.Classic().Handle,
		coin.Poa().Handle,
		coin.Callisto().Handle,
		coin.Wanchain().Handle,
		coin.Tomochain().Handle,
		coin.Ethereum().Handle,
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
