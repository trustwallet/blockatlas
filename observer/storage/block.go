package storage

import (
	"fmt"
	"sync"
)

const (
	ATLAS_BLOCK_NUMBER = "ATLAS_BLOCK_NUMBER_%d"
)

type BlockMap struct {
	heights map[uint]int64
	lock    sync.RWMutex
}

func (bm *BlockMap) SetBlock(coin uint, b int64) {
	bm.lock.Lock()
	defer bm.lock.Unlock()
	bm.heights[coin] = b
}

func (bm *BlockMap) GetBlock(coin uint) (int64, bool) {
	bm.lock.RLock()
	defer bm.lock.RUnlock()
	b, ok := bm.heights[coin]
	return b, ok
}

func (bm *BlockMap) GetHeights() map[uint]int64 {
	bm.lock.RLock()
	defer bm.lock.RUnlock()
	return bm.heights
}

func (s *Storage) GetBlockNumber(coin uint) (int64, error) {
	b, ok := s.blockHeights.GetBlock(coin)
	if ok {
		return b, nil
	}

	err := s.GetValue(getBlockKey(coin), &b)
	if err != nil {
		return 0, nil
	}
	return b, nil
}

func (s *Storage) SetBlockNumber(coin uint, num int64) error {
	s.blockHeights.SetBlock(coin, num)
	return s.Add(getBlockKey(coin), num)
}

func getBlockKey(coin uint) string {
	return fmt.Sprintf(ATLAS_BLOCK_NUMBER, coin)
}
