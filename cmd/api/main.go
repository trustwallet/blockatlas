package main

import (
	"context"
	"flag"
	"github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/trustwallet/blockatlas/api"
	"github.com/trustwallet/blockatlas/config"
	_ "github.com/trustwallet/blockatlas/docs"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform"
	"github.com/trustwallet/blockatlas/storage"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

const (
	defaultPort       = "8420"
	defaultConfigPath = "../../config.yml"
)

var (
	port, confPath string
	cache          *storage.Storage
	sg             gin.HandlerFunc
)

func init() {
	sg = sentrygin.New(sentrygin.Options{})
	cache = storage.New()

	flag.StringVar(&port, "p", defaultPort, "port for api")
	flag.StringVar(&confPath, "c", defaultConfigPath, "config file for api")

	flag.Parse()

	confPath, err := filepath.Abs(confPath)
	if err != nil {
		logger.Fatal(err)
	}

	config.LoadConfig(confPath)
	logger.InitLogger()
	platform.Init()
}

func main() {
	gin.SetMode(viper.GetString("gin.mode"))
	engine := gin.New()

	engine.Use(gin.Recovery())
	engine.Use(ginutils.CheckReverseProxy, sg)
	engine.Use(ginutils.CORSMiddleware())

	engine.OPTIONS("/*path", ginutils.CORSMiddleware())

	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/", api.GetRoot)
	engine.GET("/status", func(c *gin.Context) {
		ginutils.RenderSuccess(c, map[string]interface{}{
			"status": true,
		})
	})

	api.MakeMetricsRoute(engine)
	api.LoadPlatforms(engine)

	if viper.GetBool("observer.enabled") {
		logger.Info("Loading observer API")
		observerAPI := engine.Group("/observer/v1")
		api.SetupObserverAPI(observerAPI, cache)
	}
	if viper.GetBool("market.enabled") {
		logger.Info("Loading market API")
		marketAPI := engine.Group("/v1/market")
		api.SetupMarketAPI(marketAPI, cache)
	}

	signalForExit := make(chan os.Signal, 1)

	signal.Notify(signalForExit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatal("Application failed", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	defer func() {
		if err := srv.Shutdown(ctx); err != nil {
			logger.Fatal("Server Shutdown: ", err)
		}
	}()

	logger.Info("Running application", logger.Params{"bind": port})

	stop := <-signalForExit
	logger.Info("Stop signal Received", stop)
	logger.Info("Waiting for all jobs to stop")
}
