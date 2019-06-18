package observer

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"time"
)

type Stream struct {
	BlockAPI     blockatlas.BlockAPI
	Tracker      Tracker
	PollInterval time.Duration
	BacklogCount int
	coin         uint
	log          *logrus.Entry
}

func (s *Stream) Execute(ctx context.Context) <-chan *blockatlas.Block {
	s.coin = s.BlockAPI.Coin().Index
	s.log = logrus.WithField("platform", s.BlockAPI.Handle())
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

	if height - lastHeight > int64(s.BacklogCount) {
		lastHeight = height - int64(s.BacklogCount)
	}

	for i := lastHeight + 1; i <= height; i++ {
		block, err := s.BlockAPI.GetBlockByNumber(i)
		if err != nil {
			s.log.WithError(err).Errorf("Polling failed: could not get block %d", i)
			return
		}
		c <- block

		err = s.Tracker.SetBlockNumber(s.coin, i)
		if err != nil {
			s.log.WithError(err).Error("Polling failed: could not update block number at tracker")
			return
		}
	}
}
