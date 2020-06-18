package api

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/trustwallet/blockatlas/docs"
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

func SetupSwaggerAPI(router gin.IRouter) {
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
