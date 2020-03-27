package main

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/api"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/services/observer/healthcheck"
	"time"
)

const (
	defaultPort       = "3333"
	defaultConfigPath = "../../config.yml"
)

var (
	port, confPath string
	engine         *gin.Engine
)

func init() {
	port, confPath = internal.ParseArgs(defaultPort, defaultConfigPath)

	internal.InitConfig(confPath)

	logger.InitLogger()
	tmp := sentrygin.New(sentrygin.Options{})
	sg := &tmp

	engine = internal.InitEngine(sg, viper.GetString("gin.mode"))

	platform.Init(viper.GetString("platform"))
}

func main() {
	for _, blockAPI := range platform.BlockAPIs {
		go healthcheck.Worker(blockAPI)
		time.Sleep(time.Millisecond * 200)
	}

	api.SetupHealthCheckApi(engine)
	internal.SetupGracefulShutdown(port, engine)
}
