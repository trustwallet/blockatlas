package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

const (
	clientNamespace = "client"
)

var (
	clientLabels   = []string{"status", "endpoint", "method"}
	clientReqCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: clientNamespace,
			Name:      "http_request_count_total",
			Help:      "Total number of HTTP requests made.",
		}, clientLabels,
	)
	clientReqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: clientNamespace,
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request latencies in seconds.",
		}, clientLabels,
	)
)

func GetMetrics(status, url, method string, start time.Time) {
	endpoint := removeSensitiveInfo(url)
	lvs := []string{status, endpoint, method}
	clientReqCount.WithLabelValues(lvs...).Inc()
	clientReqDuration.WithLabelValues(lvs...).Observe(time.Since(start).Seconds())
}
