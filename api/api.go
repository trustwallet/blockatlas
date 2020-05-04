package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/platform"
)

func SetupPlatformAPI(router gin.IRouter) {
	for _, api := range platform.Platforms {
		RegisterTransactionsAPI(router, api)
		RegisterTokensAPI(router, api)
		RegisterStakeAPI(router, api)
	}
	for _, api := range platform.CollectionsAPIs {
		RegisterCollectionsAPI(router, api)
	}

	RegisterBatchAPI(router)
	RegisterDomainAPI(router)
	RegisterBasicAPI(router)
}
