package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/api"
	"github.com/trustwallet/blockatlas/build"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/platform"
)

const (
	defaultPort       = "8420"
	defaultConfigPath = "../../config.yml"
)

var (
	port, confPath, chosenPlatform string
	sg                             *gin.HandlerFunc
)

func init() {
	build.LogVersionInfo()
	port, confPath, sg = internal.InitAPI(defaultPort, defaultConfigPath)
	platform.Init(viper.GetString("platform"))
}

func main() {
	gin.SetMode(viper.GetString("gin.mode"))
	engine := gin.New()
	engine.Use(ginutils.CheckReverseProxy, *sg)
	engine.Use(ginutils.CORSMiddleware())
	engine.Use(gin.Logger())

	engine.OPTIONS("/*path", ginutils.CORSMiddleware())
	engine.GET("/", api.GetRoot)
	engine.GET("/status", func(c *gin.Context) {
		ginutils.RenderSuccess(c, map[string]interface{}{
			"status":  true,
			"version": build.Version,
			"build":   build.Build,
			"date":    build.Date,
		})
	})

	api.LoadPlatforms(engine)

	internal.SetupGracefulShutdown(port, engine)
}
