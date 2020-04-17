package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/api/endpoint"
	"github.com/trustwallet/blockatlas/pkg/servicerepo"
	"github.com/trustwallet/blockatlas/platform"
)

func SetupPlatformAPI(serviceRepo *servicerepo.ServiceRepo, root gin.IRouter) {
	endpoint.Init(serviceRepo)

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
