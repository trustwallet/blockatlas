package observer

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"sync/atomic"
	"time"
)

type Stream struct {
	BlockAPI     blockatlas.BlockAPI
	Tracker      Tracker
	PollInterval time.Duration
	BacklogCount int
	coin         uint
	log          *logrus.Entry

	// Concurrency
	blockNumber int64
}

func (s *Stream) Execute(ctx context.Context) <-chan *blockatlas.Block {
	cn := s.BlockAPI.Coin()
	s.coin = cn.ID
	s.log = logrus.WithField("platform", cn.Handle)
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
		s.log.WithError(err).Error("Polling failed: tracker didn't return last known block number")
		return
	}

	height, err := s.BlockAPI.CurrentBlockNumber()
	if err != nil {
		s.log.WithError(err).Error("Polling failed: source didn't return chain head number")
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
		block := s.loadBlock(i)
		if block != nil {
			c <- block
		}
		err = s.Tracker.SetBlockNumber(s.coin, i)
		if err != nil {
			s.log.WithError(err).Error("Polling failed: could not update block number at tracker")
		}
	}
}

func (s *Stream) loadBlock(num int64) *blockatlas.Block {
	block, err := s.BlockAPI.GetBlockByNumber(num)
	if err != nil {
		s.log.WithError(err).Errorf("Polling failed: could not get block %d", num)
		return nil
	}
	s.log.WithField("num", num).Info("Got new block")
	return block
}
