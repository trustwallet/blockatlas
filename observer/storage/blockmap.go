package storage

import (
	"sync"
)

type BlockMap struct {
	heights map[interface{}]*Block
	lock    sync.RWMutex
}

func (bm *BlockMap) SetBlock(b *Block) {
	bm.lock.Lock()
	defer bm.lock.Unlock()
	bm.heights[b.Coin] = b
}

func (bm *BlockMap) GetBlock(coin int) (*Block, bool) {
	bm.lock.RLock()
	defer bm.lock.RUnlock()
	b, ok := bm.heights[coin]
	return b, ok
}

func (bm *BlockMap) GetHeights() map[interface{}]*Block {
	bm.lock.RLock()
	defer bm.lock.RUnlock()
	return bm.heights
}
