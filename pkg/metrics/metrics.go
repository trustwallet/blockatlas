package metrics

import "github.com/prometheus/client_golang/prometheus"

// init registers the prometheus metrics
func init() {
	prometheus.MustRegister(clientReqCount, clientReqDuration, serverReqCount, serverReqDuration, serverReqSizeBytes, serverRespSizeBytes)
}
