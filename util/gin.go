package util

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func RequireConfig(keys ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		for _, key := range keys {
			if !viper.IsSet(key) {
				logrus.Errorf("Config key %s not set!", key)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}
	}
}
