package api

import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/trustwallet/blockatlas/api/endpoint"
	"github.com/trustwallet/blockatlas/api/middleware"
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
	if _, ok := api.(blockatlas.TxByAddrAndXPubAPI); ok {
		// this is XPUB style
		return
	}
	txAPI, _ := api.(blockatlas.TxAPI)
	tokenTxAPI, _ := api.(blockatlas.TokenTxAPI)

	handle := api.Coin().Handle

	router.GET("/v1/"+handle+"/:address", func(c *gin.Context) {
		endpoint.GetTransactionsHistory(c, txAPI, tokenTxAPI)
	})
	router.GET("/v2/"+handle+"/transactions/:address", func(c *gin.Context) {
		endpoint.GetTransactionsHistory(c, txAPI, tokenTxAPI)
	})
}

func RegisterTxByAddrAndXPubAPI(router gin.IRouter, api blockatlas.TxByAddrAndXPubAPI) {
	handle := api.Coin().Handle
	router.GET("/v1/" + handle + "/address/:address", func(c *gin.Context) {
		endpoint.GetTransactionsHistory(c, api, nil)
	})
	router.GET("/v1/" + handle + "/xpub/:xpub", func(c *gin.Context) {
		endpoint.GetTransactionsByXPub(c, api)
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

func RegisterDomainAPI(router gin.IRouter) {
	router.GET("/ns/lookup", endpoint.GetAddressByCoinAndDomain)
	router.GET("v2/ns/lookup", endpoint.GetAddressByCoinAndDomainBatch)
}

func RegisterBasicAPI(router gin.IRouter) {
	router.GET("/", endpoint.GetStatus)
	router.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
}
