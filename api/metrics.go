package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"regexp"
	"time"
)

const (
	namespace = "server"
)

var (
	labels = []string{"status", "endpoint", "method"}

	reqCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "",
			Name:      "http_request_count_total",
			Help:      "Total number of HTTP requests made.",
		}, labels,
	)

	reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request latencies in seconds.",
		}, labels,
	)

	reqSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: namespace,
			Name:      "http_request_size_bytes",
			Help:      "HTTP request sizes in bytes.",
		}, labels,
	)

	respSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: namespace,
			Name:      "http_response_size_bytes",
			Help:      "HTTP request sizes in bytes.",
		}, labels,
	)
)

// init registers the prometheus metrics
func init() {
	prometheus.MustRegister(reqCount, reqDuration, reqSizeBytes, respSizeBytes)
}

func PromMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		status := fmt.Sprintf("%d", c.Writer.Status())
		url := c.Request.URL.Path
		method := c.Request.Method

		reg := regexp.MustCompile(`([a-zA-Z0-9]{30,})`)
		endpoint := reg.ReplaceAllString(url, "")

		lvs := []string{status, endpoint, method}

		reqCount.WithLabelValues(lvs...).Inc()
		reqDuration.WithLabelValues(lvs...).Observe(time.Since(start).Seconds())
		reqSizeBytes.WithLabelValues(lvs...).Observe(calcRequestSize(c.Request))
		respSizeBytes.WithLabelValues(lvs...).Observe(float64(c.Writer.Size()))
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
