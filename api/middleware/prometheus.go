package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"regexp"
)

var labels = []string{"status", "endpoint", "method"}

func Prometheus() gin.HandlerFunc {
	serverReqCount := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "atlas",
			Name:      "http_request_count_total",
			Help:      "Total number of HTTP requests made.",
		}, labels,
	)
	prometheus.MustRegister(serverReqCount)

	return func(c *gin.Context) {
		c.Next()

		status := fmt.Sprintf("%d", c.Writer.Status())
		url := c.Request.URL.Path
		method := c.Request.Method

		lvs := []string{status, removeAddress(url), method}

		serverReqCount.WithLabelValues(lvs...).Inc()
	}
}

func removeAddress(info string) string {
	reg := regexp.MustCompile(`([a-zA-Z0-9\s]{30,})|([0-9]{4,})|(=(.*?)[^(&|$)]+)|(--[^$]+)|(&asset_contract_addresses)`)
	return reg.ReplaceAllString(info, "")
}
