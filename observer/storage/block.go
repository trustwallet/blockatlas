package storage

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"sync"
)

type BlockMap struct {
	heights map[uint]Block
	lock    sync.RWMutex
}

func (s *Storage) GetHeights() map[uint]Block {
	s.blockHeights.lock.RLock()
	defer s.blockHeights.lock.RUnlock()
	return s.blockHeights.heights
}

func (s *Storage) GetBlock(coin uint) (b Block, ok bool) {
	s.blockHeights.lock.Lock()
	defer s.blockHeights.lock.Unlock()

	b, ok = s.blockHeights.heights[coin]
	if ok {
		return
	}

	b = Block{Coin: coin}
	s.blockHeights.heights[b.Coin] = b
	return
}

func (s *Storage) GetBlockNumber(coin uint) (int64, error) {
	b, ok := s.GetBlock(coin)
	if ok {
		return b.BlockHeight, nil
	}
	err := s.Get(&b)
	if err != nil {
		return 0, nil
	}
	return b.BlockHeight, nil
}

func (s *Storage) SetBlockNumber(coin uint, num int64) {
	b, _ := s.GetBlock(coin)
	s.blockHeights.lock.Lock()
	defer s.blockHeights.lock.Unlock()
	b.BlockHeight = num
	s.blockHeights.heights[b.Coin] = b
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

	values := make([]interface{}, 0)
	h := s.GetHeights()
	for _, v := range h {
		values = append(values, &v)
	}
	err := s.CreateOrUpdateMany(values...)
	if err != nil {
		return errors.E(err).PushToSentry()
	}
	return nil
}
