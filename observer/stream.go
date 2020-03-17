package observer

import (
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"sync"
	"sync/atomic"
	"time"
)

type Stream struct {
	BlockAPI     blockatlas.BlockAPI
	Tracker      storage.Tracker
	PollInterval time.Duration
	BacklogCount int
	coin         uint
	logParams    logger.Params
	blockNumber  int64
	mu           sync.Mutex
}

func (s *Stream) setLatestParsedBlock(num int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	newNum := atomic.AddInt64(&s.blockNumber, 1)
	err := s.Tracker.SetBlockNumber(s.coin, newNum)
	if err != nil {
		return errors.E(err, "SetBlockNumber failed", s.logParams, logger.Params{"block": num, "coin": s.coin})
	}
	return nil
}

func (s *Stream) Execute() {
	coin := s.BlockAPI.Coin()
	s.coin = coin.ID
	s.logParams = logger.Params{"platform": coin.Handle}
	conns := viper.GetInt("observer.stream_conns")
	if conns == 0 {
		logger.Fatal("observer.stream_conns is 0")
	}

	for {
		lastHeight, height, err := s.getLastBlockParams()
		if err != nil {
			logger.Error(err)

		}

		s.fetchAndPublishBlocks(lastHeight, height)

		time.Sleep(s.PollInterval)
	}
}

func (s *Stream) getLastBlockParams() (int64, int64, error) {
	lastHeight, err := s.Tracker.GetBlockNumber(s.coin)
	if err != nil {
		return 0, 0, errors.E(err, "Polling failed: tracker didn't return last known block number", s.logParams)
	}

	height, err := s.BlockAPI.CurrentBlockNumber()
	height -= s.BlockAPI.Coin().MinConfirmations
	if err != nil {
		return 0, 0, errors.E(err, "Polling failed: source didn't return chain head number", s.logParams)
	}

	if height-lastHeight > int64(s.BacklogCount) {
		lastHeight = height - int64(s.BacklogCount)
	}

	backLogMax := viper.GetInt64("observer.backlog_max_blocks")
	if height-lastHeight > backLogMax {
		lastHeight = height - backLogMax
	}

	return lastHeight, height, nil
}

func (s *Stream) fetchAndPublishBlocks(lastHeight, height int64) {
	atomic.StoreInt64(&s.blockNumber, lastHeight)
	var wg sync.WaitGroup
	for i := lastHeight + 1; i <= height; i++ {
		wg.Add(1)
		go s.publishBlock(i, &wg)
	}
	wg.Wait()
}

func (s *Stream) publishBlock(num int64, wg *sync.WaitGroup) {
	defer wg.Done()

	block, err := retry(5, time.Second*5, s.BlockAPI.GetBlockByNumber, num)
	if err != nil {
		logger.Error(err, "Polling failed: could not get block", s.logParams, logger.Params{"block": num})
		return
	}

	logger.Info(err, "Got new block", s.logParams, logger.Params{"block": num, "txs": len(block.Txs)})

	// Ignore block if it's not marked as parsed (set to redis) to prevent double notifications

	// TODO: add retry
	if err := s.setLatestParsedBlock(num); err != nil {
		logger.Error(err)
		return
	}

	if err := s.publish(*block); err != nil {
		logger.Error(err)
	}
}

func (s *Stream) publish(block blockatlas.Block) error {
	body, err := json.Marshal(block)
	if err != nil {
		return err
	}
	return mq.ConfirmedBlocks.Publish(body)
}
