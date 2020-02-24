package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/api"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/storage"
)

const (
	defaultPort       = "8422"
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
	gin.SetMode(viper.GetString("gin.mode"))

	engine := gin.New()
	engine.Use(ginutils.CheckReverseProxy, *sg)
	engine.Use(ginutils.CORSMiddleware())
	engine.Use(gin.Logger())
	engine.OPTIONS("/*path", ginutils.CORSMiddleware())

	api.SetupObserverAPI(engine, cache)
	internal.SetupGracefulShutdown(port, engine)
}
