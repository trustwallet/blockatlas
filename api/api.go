package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/platform"
)

func SetupPlatformAPI(root gin.IRouter) {
	for _, api := range platform.Platforms {
		RegisterCollectionsAPI(root, api)
		RegisterTransactionsAPI(root, api)
		RegisterCustomAPI(root, api)
		RegisterTokensAPI(root, api)
		RegisterStakeAPI(root, api)
	}

	RegisterBatchAPI(root)
	RegisterDomainAPI(root)
	RegisterBasicAPI(root)
}
