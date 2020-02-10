package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/trustwallet/blockatlas/api"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
)

const (
	defaultPort       = "8423"
	defaultConfigPath = "../../config.yml"
	allPlatforms      = "all"
)

var (
	port, confPath string
	sg             *gin.HandlerFunc
)

func init() {
	port, confPath, sg = internal.InitAPI(defaultPort, defaultConfigPath)
}

func main() {
	gin.SetMode(viper.GetString("gin.mode"))

	engine := gin.New()
	engine.Use(ginutils.CheckReverseProxy, *sg)
	engine.Use(ginutils.CORSMiddleware())

	engine.OPTIONS("/*path", ginutils.CORSMiddleware())
	engine.GET("/", api.GetRoot)
	engine.GET("/status", func(c *gin.Context) {
		ginutils.RenderSuccess(c, map[string]interface{}{
			"status": true,
		})
	})

	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	internal.SetupGracefulShutdown(port, engine)
}
