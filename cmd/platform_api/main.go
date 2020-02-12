package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/api"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/platform"
)

const (
	defaultPort       = "8420"
	defaultConfigPath = "../../config.yml"
	allPlatforms      = "all"
)

var (
	port, confPath, chosenPlatform string
	sg                             *gin.HandlerFunc
)

func init() {
	port, confPath, sg = internal.InitAPI(defaultPort, defaultConfigPath)
	chosenPlatform = viper.GetString("platform")
	platform.Init(chosenPlatform)
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

	api.LoadPlatforms(engine)

	internal.SetupGracefulShutdown(port, engine)
}
