package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/api/endpoint"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/services/tokenindexer"
	"github.com/trustwallet/golibs/network/middleware"
)

func RegisterTransactionsAPI(router gin.IRouter, api blockatlas.Platform) {
	handle := api.Coin().Handle
	txUtxoAPI, ok := api.(blockatlas.TxUtxoAPI)
	if ok {
		router.GET("/v1/"+handle+"/address/:address", func(c *gin.Context) {
			endpoint.GetTransactionsHistory(c, txUtxoAPI, nil)
		})
		router.GET("/v1/"+handle+"/xpub/:xpub", func(c *gin.Context) {
			endpoint.GetTransactionsByXpub(c, txUtxoAPI)
		})
		router.GET("/v2/"+handle+"/transactions/xpub/:xpub", func(c *gin.Context) {
			endpoint.GetTransactionsByXpub(c, txUtxoAPI)
		})
		return
	}
	txAPI, okTxApi := api.(blockatlas.TxAPI)
	tokenTxAPI, okTokenTxApi := api.(blockatlas.TokenTxAPI)
	if okTxApi || okTokenTxApi {
		router.GET("/v1/"+handle+"/:address", func(c *gin.Context) {
			endpoint.GetTransactionsHistory(c, txAPI, tokenTxAPI)
		})
		router.GET("/v2/"+handle+"/transactions/:address", func(c *gin.Context) {
			endpoint.GetTransactionsHistory(c, txAPI, tokenTxAPI)
		})
	}
}

func RegisterBlockAPI(router gin.IRouter, api blockatlas.Platform) {
	handle := api.Coin().Handle
	if blockAPI, ok := api.(blockatlas.BlockAPI); ok {
		router.GET("/v2/"+handle+"/blocks/:block", func(c *gin.Context) {
			endpoint.GetBlock(c, blockAPI)
		})
	}
}

func RegisterTokensAPI(router gin.IRouter, api blockatlas.Platform) {
	tokenAPI, ok := api.(blockatlas.TokensAPI)
	if !ok {
		return
	}
	handle := tokenAPI.Coin().Handle
	router.GET("/v2/"+handle+"/tokens/:address", func(c *gin.Context) {
		endpoint.GetTokensByAddress(c, tokenAPI)
	})
	router.GET("/v2/"+handle+"/tokens/:address/ids", func(c *gin.Context) {
		endpoint.GetTokensIdsByAddress(c, tokenAPI)
	})
}

func RegisterStakeAPI(router gin.IRouter, api blockatlas.Platform) {
	stakeAPI, ok := api.(blockatlas.StakeAPI)
	if !ok {
		return
	}
	handle := api.Coin().Handle
	router.GET("/v2/"+handle+"/staking/validators", middleware.CacheMiddleware(time.Hour, func(c *gin.Context) {
		endpoint.GetValidators(c, stakeAPI)
	}))
	router.GET("/v2/"+handle+"/staking/delegations/:address", func(c *gin.Context) {
		endpoint.GetStakingDelegationsForSpecificCoin(c, stakeAPI)
	})
}

func RegisterCollectionsAPI(router gin.IRouter, api blockatlas.CollectionsAPI) {
	handle := api.Coin().Handle
	router.GET("/v4/"+handle+"/collections/:owner/collection/:collection_id", func(c *gin.Context) {
		endpoint.GetCollectiblesForSpecificCollectionAndOwner(c, api)
	})
}

func RegisterBatchAPI(router gin.IRouter) {
	router.GET("/v3/staking/list", middleware.CacheMiddleware(time.Hour*10, func(c *gin.Context) {
		endpoint.GetStakeInfoForCoins(c, platform.StakeAPIs)
	}))
	router.POST("/v2/staking/delegations", func(c *gin.Context) {
		endpoint.GetStakeDelegationsWithAllInfoForBatch(c, platform.StakeAPIs)
	})
	router.POST("/v2/staking/list", middleware.CacheMiddleware(time.Hour, func(c *gin.Context) {
		endpoint.GetStakeInfoForBatch(c, platform.StakeAPIs)
	}))
	router.POST("/v4/collectibles/categories", func(c *gin.Context) {
		endpoint.GetCollectionCategoriesFromList(c, platform.CollectionsAPIs)
	})
}

func RegisterBasicAPI(router gin.IRouter) {
	router.GET("/", endpoint.GetStatus)
}

func RegisterTokensIndexAPI(router gin.IRouter, instance tokenindexer.Instance) {
	router.GET("/v3/tokens/new", func(c *gin.Context) {
		endpoint.GetNewTokens(c, instance)
	})
	router.POST("/v1/assets/associations", func(c *gin.Context) {
		endpoint.GetTokensByAddressV3(c, instance)
	})
}
