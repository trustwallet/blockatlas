package main

import (
	"github.com/gin-gonic/gin"
	"github.com/trustwallet/blockatlas/api"
	"github.com/trustwallet/blockatlas/config"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
)

const (
	defaultPort       = "8421"
	defaultConfigPath = "../../config.yml"
)

var (
	port, confPath string
	cache          *storage.Storage
	sg             *gin.HandlerFunc
)

func init() {
	port, confPath, sg, cache = internal.InitAPIWithRedis(defaultPort, defaultConfigPath)
}

func main() {
	gin.SetMode(config.Configuration.Gin.Mode)

	engine := gin.New()
	engine.Use(ginutils.CheckReverseProxy, *sg)
	engine.Use(ginutils.CORSMiddleware())
	engine.Use(gin.Logger())

	engine.OPTIONS("/*path", ginutils.CORSMiddleware())
	engine.GET("/", api.GetRoot)
	engine.GET("/status", func(c *gin.Context) {
		ginutils.RenderSuccess(c, map[string]interface{}{
			"status": true,
		})
	})

	logger.Info("Loading market API")

	marketAPI := engine.Group("/v1/market")

	api.SetupMarketAPI(marketAPI, cache)
	internal.SetupGracefulShutdown(port, engine)
}
