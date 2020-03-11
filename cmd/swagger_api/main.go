package main

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
)

const (
	defaultPort       = "8423"
	defaultConfigPath = "../../config.yml"
)

var (
	port, confPath string
	engine         *gin.Engine
)

func init() {
	port, confPath = internal.ParseArgs(defaultPort, defaultConfigPath)
	tmp := sentrygin.New(sentrygin.Options{})
	sg := &tmp
	internal.InitConfig(confPath)
	logger.InitLogger()
	engine = internal.InitEngine(sg, viper.GetString("gin.mode"))
}

func main() {
	logger.Info("Loading Swagger API")
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "swagger/index.html")
	})
	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	internal.SetupGracefulShutdown(port, engine)
}
