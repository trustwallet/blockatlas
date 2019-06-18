package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
)

var Cmd = cobra.Command{
	Use:     "api <bind>",
	Short:   "API server",
	Args:    cobra.MaximumNArgs(1),
	Run:     run,
}

func run(_ *cobra.Command, args []string) {
	var bind string
	if len(args) == 0 {
		bind = ":8420"
	} else {
		bind = args[0]
	}

	gin.SetMode(viper.GetString("gin.mode"))
	router := gin.Default()
	router.Use(util.CheckReverseProxy)
	router.GET("/", getRoot)
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": true,
		})
	})

	loadPlatforms(router)

	logrus.WithField("bind", bind).Info("Running application")
	if err := router.Run(bind); err != nil {
		logrus.WithError(err).Fatal("Application failed")
	}
}
