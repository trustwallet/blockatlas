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

	var duration time.Duration
	t := api.Coin().BlockTime / 1000

	if t > 0 && t < 4 {
		duration = time.Duration(int64(time.Second) * int64(t))
	} else {
		duration = time.Minute * 11
	}
	logger.Info("Setting update duration", logger.Params{"duration": duration})

	for {
		currentBlock, err = storage.GetLastParsedBlockNumber(api.Coin().ID)
		if err != nil {
			logger.Error(err)
			continue
		}

		latestBlock, err := api.CurrentBlockNumber()
		if err != nil {
			logger.Error(err)
			continue
		}

		logger.Info("fetched", logger.Params{"currentBlock": currentBlock, "lastParsedBlock": lastParsedBlock, "latestBlock": latestBlock, "handle": api.Coin().Handle})

		if currentBlock > lastParsedBlock || latestBlock == currentBlock {
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

		time.Sleep(duration)
	}
}

func GetStatus() map[string]Info {
	return observerStatus.get()
}
