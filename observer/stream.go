package observer

import (
	"context"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/util"
	"sync"
	"sync/atomic"
	"time"
)

type Stream struct {
	BlockAPI     blockatlas.BlockAPI
	Tracker      Tracker
	PollInterval time.Duration
	BacklogCount int
	coin         uint
	logParams    logger.Params

	// Concurrency
	blockNumber int64
	semaphore   *util.Semaphore
	wg          sync.WaitGroup
}

func (s *Stream) Execute(ctx context.Context) <-chan *blockatlas.Block {
	cn := s.BlockAPI.Coin()
	s.coin = cn.ID
	s.logParams = logger.Params{"platform": cn.Handle}
	conns := viper.GetInt("observer.stream_conns")
	if conns == 0 {
		logger.Fatal("observer.stream_conns is 0")
	}
	s.semaphore = util.NewSemaphore(conns)
	c := make(chan *blockatlas.Block)
	go s.run(ctx, c)
	return c
}

func (s *Stream) run(ctx context.Context, c chan<- *blockatlas.Block) {
	ticker := time.NewTicker(s.PollInterval)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			close(c)
			return
		case <-ticker.C:
			s.load(c)
		}
	}
}

func (s *Stream) load(c chan<- *blockatlas.Block) {
	lastHeight, err := s.Tracker.GetBlockNumber(s.coin)
	if err != nil {
		logger.Error(err, "Polling failed: tracker didn't return last known block number", s.logParams)
		return
	}

	height, err := s.BlockAPI.CurrentBlockNumber()
	height -= s.BlockAPI.Coin().MinConfirmations
	if err != nil {
		logger.Error(err, "Polling failed: source didn't return chain head number", s.logParams)
		return
	}

	if height-lastHeight > int64(s.BacklogCount) {
		lastHeight = height - int64(s.BacklogCount)
	}
	backLogMax := viper.GetInt64("observer.backlog_max_blocks")
	if height-lastHeight > backLogMax {
		lastHeight = height - backLogMax
	}

	atomic.StoreInt64(&s.blockNumber, lastHeight)
	for i := lastHeight + 1; i <= height; i++ {
		s.wg.Add(1)
		go s.loadBlock(c, i)
	}
	s.wg.Wait()
}

func (s *Stream) loadBlock(c chan<- *blockatlas.Block, num int64) {
	defer s.wg.Done()
	s.semaphore.Acquire()
	defer s.semaphore.Release()

	block, err := retry(5, time.Second*5, s.BlockAPI.GetBlockByNumber, num)
	if err != nil {
		logger.Error(err, "Polling failed: could not get block", s.logParams, logger.Params{"block": num})
		return
	}
	c <- block
	logger.Info(err, "Got new block", s.logParams, logger.Params{"block": num, "txs": len(block.Txs)})

	// Not strictly correct nor avoids race conditions
	// But good enough
	newNum := atomic.AddInt64(&s.blockNumber, 1)
	err = s.Tracker.SetBlockNumber(s.coin, newNum)
	if err != nil {
		logger.Error(err, "Polling failed: could not update block number at tracker", s.logParams)
		return
	}
}
