package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"regexp"
)

// init registers the prometheus metrics
func init() {
	prometheus.MustRegister(clientReqCount, clientReqDuration, serverReqCount, serverReqDuration, serverReqSizeBytes, serverRespSizeBytes)
}

func removeSensitiveInfo(info string) string {
	reg := regexp.MustCompile(`([a-zA-Z0-9\s]{30,})|([0-9]{4,})|(=(.*?)[^(&|$)]+)|(--[^$]+)`)
	return reg.ReplaceAllString(info, "")
}
