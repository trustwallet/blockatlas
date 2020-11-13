package api

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/services/tokenindexer"
	"github.com/trustwallet/blockatlas/services/tokensearcher"
)

func SetupPlatformAPI(router gin.IRouter) {
	for _, api := range platform.Platforms {
		RegisterTransactionsAPI(router, api)
		RegisterTokensAPI(router, api)
		RegisterStakeAPI(router, api)
		RegisterBlockAPI(router, api)
	}
	for _, api := range platform.CollectionsAPIs {
		RegisterCollectionsAPI(router, api)
	}

	RegisterBatchAPI(router)
	RegisterBasicAPI(router)
}

func SetupTokensSearcherAPI(router gin.IRouter, instance tokensearcher.Instance) {
	RegisterTokensSearcherAPI(router, instance)
}

func SetupTokensIndexAPI(router gin.IRouter, instance tokenindexer.Instance) {
	RegisterTokensIndexAPI(router, instance)
}
func SetupSwaggerAPI(router gin.IRouter) {
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
