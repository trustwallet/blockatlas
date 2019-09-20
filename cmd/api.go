package cmd

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/api"
	observerStorage "github.com/trustwallet/blockatlas/observer/storage"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
)

var apiCmd = cobra.Command{
	Use:   "api <bind>",
	Short: "API server",
	Args:  cobra.MaximumNArgs(1),
	Run:   runApi,
}

func runApi(_ *cobra.Command, args []string) {
	var bind string
	if len(args) == 0 {
		bind = ":8420"
	} else {
		bind = args[0]
	}
	RunApi(bind, nil)
}

func RunApi(bind string, c chan *gin.Engine) {
	gin.SetMode(viper.GetString("gin.mode"))
	engine := gin.Default()

	sg := sentrygin.New(sentrygin.Options{})
	engine.Use(util.CheckReverseProxy, sg)

	api.MakeMetricsRoute(engine)

	engine.GET("/", api.GetRoot)
	engine.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": true,
		})
	})

	api.LoadPlatforms(engine)
	if observerStorage.App != nil {
		observerAPI := engine.Group("/observer/v1")
		api.SetupObserverAPI(observerAPI)
	}

	if c != nil {
		c <- engine
	}

	logger.Info("Running application", logger.Params{"bind": bind})
	if err := engine.Run(bind); err != nil {
		logger.Fatal("Application failed", err)
	}
}

func init() {
	rootCmd.AddCommand(&apiCmd)
}
