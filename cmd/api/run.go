package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	observerStorage "github.com/trustwallet/blockatlas/observer/storage"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
)

var Cmd = cobra.Command{
	Use:   "api <bind>",
	Short: "API server",
	Args:  cobra.MaximumNArgs(1),
	Run:   run,
}

var engine *gin.Engine

func run(_ *cobra.Command, args []string) {
	var bind string
	if len(args) == 0 {
		bind = ":8420"
	} else {
		bind = args[0]
	}

	Run(bind, nil)
}

func Run(bind string, c chan *gin.Engine) {
	gin.SetMode(viper.GetString("gin.mode"))
	engine = gin.Default()
	engine.Use(util.CheckReverseProxy)
	engine.GET("/", getRoot)
	engine.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": true,
		})
	})

	loadPlatforms(engine)
	if observerStorage.App != nil {
		observerAPI := engine.Group("/observer/v1")
		setupObserverAPI(observerAPI)
	}

	if c != nil {
		c <- engine
	}

	logrus.WithField("bind", bind).Info("Running application")
	if err := engine.Run(bind); err != nil {
		logrus.WithError(err).Fatal("Application failed")
	}
}
