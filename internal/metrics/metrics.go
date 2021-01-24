package metrics

import (
	"time"

	"github.com/trustwallet/blockatlas/db"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	workerBlockParsing = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "worker",
			Name:      "block_parsing",
			Help:      "Last parsed block",
		},
		[]string{"coin", "priority"},
	)
)

func setupUpdateTrackerMetrics(db *db.Instance) {
	go func() {
		for {
			trackers, err := db.GetLastParsedBlockNumbers()
			if err != nil {
				continue
			}
			for _, tracker := range trackers {
				workerBlockParsing.With(prometheus.Labels{"coin": tracker.Coin, "priority": tracker.Priority}).Add(float64(tracker.Height))
			}
			time.Sleep(1 * time.Second)
		}
	}()
}

func Setup(db *db.Instance) {
	setupUpdateTrackerMetrics(db)

	prometheus.DefaultRegisterer.Unregister(prometheus.NewGoCollector())
	prometheus.DefaultRegisterer.Unregister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
}
