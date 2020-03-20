package healthcheck

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"sync"
	"time"
)

var observerStatus status

func init() {
	observerStatus = status{
		m: make(map[string]Info),
	}
}

type status struct {
	m map[string]Info
	sync.RWMutex
}

func (cr *status) update(coin string, info Info) {
	cr.Lock()
	cr.m[coin] = info
	cr.Unlock()
}

func (cr *status) get() map[string]Info {
	cr.RLock()
	defer cr.RUnlock()
	return cr.m
}

func Worker(storage storage.Tracker, api blockatlas.BlockAPI) {
	var (
		lastParsedBlock, currentBlock int64
		err                           error
	)

	lastParsedBlock, err = storage.GetLastParsedBlockNumber(api.Coin().ID)
	if err != nil {
		logger.Fatal(err)
	}

	for {
		currentBlock, err = storage.GetLastParsedBlockNumber(api.Coin().ID)
		if err != nil {
			logger.Error(err)
			continue
		}

		if currentBlock > lastParsedBlock {
			lastParsedBlock = currentBlock
			observerStatus.update(api.Coin().Handle, Info{
				LastParsedBlock: lastParsedBlock,
				Healthy:         true,
			})
		} else {
			observerStatus.update(api.Coin().Handle, Info{
				LastParsedBlock: lastParsedBlock,
				Healthy:         false,
			})
		}

		time.Sleep(time.Minute * 5)
	}
}

func GetStatus() map[string]Info {
	return observerStatus.get()
}
