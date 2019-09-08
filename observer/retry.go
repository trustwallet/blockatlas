package observer

import (
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"math/rand"
	"time"
)

type GetBlockByNumber func(num int64) (*blockatlas.Block, error)

type stop struct {
	error
}

func retry(attempts int, sleep time.Duration, f GetBlockByNumber, n int64, log *logrus.Entry) (*blockatlas.Block, error) {
	r, err := f(n)
	if err != nil {
		if s, ok := err.(stop); ok {
			return nil, s.error
		}
		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			logger.Info("GetBlockByNumber retry",
				logger.Params{
					"number":   n,
					"attempts": attempts,
					"sleep":    sleep.String(),
				},
			)

			time.Sleep(sleep)
			return retry(attempts, sleep*2, f, n, log)
		}
	}
	return r, err
}
