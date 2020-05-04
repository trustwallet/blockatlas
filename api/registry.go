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

func RegisterCollectionsAPI(router gin.IRouter, api blockatlas.CollectionsAPI) {
	handle := api.Coin().Handle
	router.GET("/v3/" + handle + "/collections/:owner/collection/:collection_id", func(c *gin.Context) {
		endpoint.GetCollectiblesForSpecificCollectionAndOwnerV3(c, api)
	})
	router.GET("/v3/" + handle + "/collections/:owner", func(c *gin.Context) {
		endpoint.GetCollectiblesForOwnerV3(c, api)
	})
	router.GET("/v4/" + handle + "/collections/:owner/collection/:collection_id", func(c *gin.Context) {
		endpoint.GetCollectiblesForSpecificCollectionAndOwner(c, api)
	})
}

func RegisterTokensAPI(router gin.IRouter, api blockatlas.TokensAPI) {
	handle := api.Coin().Handle
	router.GET("/v2/" + handle + "/tokens/:address", func(c *gin.Context) {
		endpoint.GetTokensByAddress(c, api)
	})
}

func RegisterTxByAddrAPI(router gin.IRouter, api blockatlas.TxAPI) {
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

func RegisterTxByAddrAndXpubAPI(router gin.IRouter, api blockatlas.TxByAddrAndXpubAPI) {
	handle := api.Coin().Handle
	router.GET("/v1/" + handle + "/address/:address", func(c *gin.Context) {
		endpoint.GetTransactionsHistory(c, api, nil)
	})
	router.GET("/v1/" + handle + "/xpub/:xpub", func(c *gin.Context) {
		endpoint.GetTransactionsByXpub(c, api)
	})
}

func RegisterStakeAPI(router gin.IRouter, api blockatlas.StakeAPI) {
	handle := api.Coin().Handle
	router.GET("/v2/" + handle + "/staking/validators", middleware.CacheMiddleware(time.Hour, func(c *gin.Context) {
		endpoint.GetValidators(c, api)
	}))
	router.GET("/v2/" + handle + "/staking/delegations/:address", func(c *gin.Context) {
		endpoint.GetStakingDelegationsForSpecificCoin(c, api)
	})
}

func RegisterBatchAPI(router gin.IRouter) {
	router.POST("/v2/staking/delegations", func(c *gin.Context) {
		endpoint.GetStakeDelegationsWithAllInfoForBatch(c, platform.StakeAPIs)
	})
	router.POST("/v2/staking/list", middleware.CacheMiddleware(time.Hour, func(c *gin.Context) {
		endpoint.GetStakeInfoForBatch(c, platform.StakeAPIs)
	}))
	router.POST("/v3/collectibles/categories", func(c *gin.Context) {
		endpoint.GetCollectionCategoriesFromListV3(c, platform.CollectionsAPIs)
	})
	router.POST("/v4/collectibles/categories", func(c *gin.Context) {
		endpoint.GetCollectionCategoriesFromList(c, platform.CollectionsAPIs)
	})
	router.POST("/v2/tokens", func(c *gin.Context) {
		endpoint.GetTokens(c, platform.TokensAPIs)
	})
}

func RegisterDomainAPI(router gin.IRouter) {
	router.GET("/ns/lookup", endpoint.GetAddressByCoinAndDomain)
	router.GET("/v2/ns/lookup", endpoint.GetAddressByCoinAndDomainBatch)
}

func RegisterBasicAPI(router gin.IRouter) {
	router.GET("/", endpoint.GetStatus)
	router.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
}
