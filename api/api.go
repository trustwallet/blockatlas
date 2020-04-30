package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/platform"
)

func SetupPlatformAPI(router gin.IRouter) {
	for _, api := range platform.TxByAddrAPIs {
		RegisterTxByAddrAPI(router, api)
	}
	for _, api := range platform.TxByAddrAndXPubAPIs {
		RegisterTxByAddrAndXPubAPI(router, api)
	}
	for _, api := range platform.TokensAPIs {
		RegisterTokensAPI(router, api)
	}
	for _, api := range platform.StakeAPIs {
		RegisterStakeAPI(router, api)
	}
	for _, api := range platform.CollectionsAPIs {
		RegisterCollectionsAPI(router, api)
	}

	RegisterBatchAPI(router)
	RegisterDomainAPI(router)
	RegisterBasicAPI(router)
}
