package storage

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/storage/redis"
	"time"
)

type Storage struct {
	redis.Redis
	blockHeights BlockMap
}

func New() *Storage {
	s := new(Storage)
	s.blockHeights.heights = make(map[uint]int64)
	return s
}

func RestoreConnectionWorker(storage *Storage, uri string, timeout time.Duration) {
	logger.Info("Run Redis RestoreConnectionWorker")
	for {
		if !storage.IsReady() {
			for {
				logger.Warn("Redis is not available now")
				logger.Warn("Trying to connect to MQ...")
				if err := storage.Init(uri); err != nil {
					logger.Warn("Redis is still unavailable")
					time.Sleep(timeout)
					continue
				} else {
					logger.Info("Redis connection restored")
					break
				}
			}
		}
		time.Sleep(timeout)
	}
}

type Tracker interface {
	GetLastParsedBlockNumber(coin uint) (int64, error)
	SetLastParsedBlockNumber(coin uint, num int64) error
}

type Addresses interface {
	FindSubscriptions(coin uint, addresses []string) ([]blockatlas.Subscription, error)
	AddSubscriptions(subscriptions []blockatlas.Subscription) error
	DeleteSubscriptions(subscriptions []blockatlas.Subscription) error
}
