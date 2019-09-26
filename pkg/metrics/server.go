package metrics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
)

const (
	serverNamespace = "server"
)

var (
	serverLabels   = []string{"status", "endpoint", "method"}
	serverReqCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: serverNamespace,
			Name:      "http_request_count_total",
			Help:      "Total number of HTTP requests made.",
		}, serverLabels,
	)
	serverReqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: serverNamespace,
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request latencies in seconds.",
		}, serverLabels,
	)

	serverReqSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: serverNamespace,
			Name:      "http_request_size_bytes",
			Help:      "HTTP request sizes in bytes.",
		}, serverLabels,
	)
	serverRespSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: serverNamespace,
			Name:      "http_response_size_bytes",
			Help:      "HTTP request sizes in bytes.",
		}, serverLabels,
	)
)

func PromMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		status := fmt.Sprintf("%d", c.Writer.Status())
		url := c.Request.URL.Path
		method := c.Request.Method

		endpoint := removeSensitiveInfo(url)
		lvs := []string{status, endpoint, method}

		serverReqCount.WithLabelValues(lvs...).Inc()
		serverReqDuration.WithLabelValues(lvs...).Observe(time.Since(start).Seconds())
		serverReqSizeBytes.WithLabelValues(lvs...).Observe(calcRequestSize(c.Request))
		serverRespSizeBytes.WithLabelValues(lvs...).Observe(float64(c.Writer.Size()))
	}
}

func calcRequestSize(r *http.Request) float64 {
	size := 0
	if r.URL != nil {
		size = len(r.URL.String())
	}

	size += len(r.Method)
	size += len(r.Proto)

	for name, values := range r.Header {
		size += len(name)
		for _, value := range values {
			size += len(value)
		}
	}
	size += len(r.Host)

	// r.Form and r.MultipartForm are assumed to be included in r.URL.
	if r.ContentLength != -1 {
		size += int(r.ContentLength)
	}
	return float64(size)
}
