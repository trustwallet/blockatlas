package storage

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"sync"
)

type BlockMap struct {
	heights map[uint]*Block
	lock    sync.RWMutex
}

func (s *Storage) GetBlock(coin uint) (b *Block, ok bool) {
	s.blockHeights.lock.RLock()
	b, ok = s.blockHeights.heights[coin]
	s.blockHeights.lock.RUnlock()
	if ok {
		return
	}

	s.blockHeights.lock.Lock()
	b = &Block{Coin: coin}
	s.blockHeights.heights[coin] = b
	s.blockHeights.lock.Unlock()

	return
}

func (s *Storage) GetBlockNumber(coin uint) (int64, error) {
	b, ok := s.GetBlock(coin)
	if ok {
		return b.BlockHeight, nil
	}
	err := s.Get(b)
	if err != nil {
		return 0, nil
	}
	return b.BlockHeight, nil
}

func (s *Storage) SetBlockNumber(coin uint, num int64) error {
	b, _ := s.GetBlock(coin)
	s.blockHeights.lock.Lock()
	b.BlockHeight = num
	s.blockHeights.lock.Unlock()
	return s.SaveBlock(coin, num)
}

func (s *Storage) SaveBlock(coin uint, num int64) error {
	b := Block{Coin: coin, BlockHeight: num}
	err := s.CreateOrUpdate(&b)
	if err != nil {
		return errors.E(err, errors.Params{"block": num, "coin": coin}).PushToSentry()
	}
	return nil
}

func (s *Storage) SaveAllBlocks() error {
	logger.Info("Saving cache blocks in database")

	values := make([]*Block, 0)
	for _, v := range s.blockHeights.heights {
		values = append(values, v)
	}
	err := s.CreateOrUpdateMany(values)
	if err != nil {
		return errors.E(err).PushToSentry()
	}
	return nil
}
