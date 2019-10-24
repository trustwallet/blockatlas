package storage

import (
	"sync"
)

type BlockMap struct {
	heights map[interface{}]*Block
	lock    sync.RWMutex
}

func (s *BlockMap) SetBlock(b *Block) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.heights[b.Coin] = b
}

func (s *BlockMap) GetBlock(coin int) (*Block, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	b, ok := s.heights[coin]
	return b, ok
}

func (s *BlockMap) GetHeights() map[interface{}]*Block {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.heights
}
