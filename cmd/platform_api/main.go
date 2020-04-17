package main

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/api"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/servicerepo"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/services"
)

const (
	defaultPort       = "8420"
	defaultConfigPath = "../../config.yml"
)

var (
	port, confPath string
	serviceRepo    *servicerepo.ServiceRepo
	engine         *gin.Engine
)

func init() {
	port, confPath = internal.ParseArgs(defaultPort, defaultConfigPath)
	serviceRepo = servicerepo.New()

	internal.InitConfig(confPath)
	logger.InitLogger()
	tmp := sentrygin.New(sentrygin.Options{})
	sg := &tmp

	services.Init(serviceRepo)

	engine = internal.InitEngine(sg, viper.GetString("gin.mode"))

	platform.Init(serviceRepo, viper.GetString("platform"))
}

func main() {
	api.SetupPlatformAPI(serviceRepo, engine)
	internal.SetupGracefulShutdown(port, engine)
}
