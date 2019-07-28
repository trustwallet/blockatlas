package util

import (
	"github.com/trustwallet/blockatlas"
	"sync"
)

type TxSet struct {
	items map[blockatlas.Tx]bool
	lock  sync.RWMutex
}

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *TxSet) Add(t blockatlas.Tx) *TxSet {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.items == nil {
		s.items = make(map[blockatlas.Tx]bool)
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return s
}

// Clear removes all elements from the Set
func (s *TxSet) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = make(map[blockatlas.Tx]bool)
}

func (s *TxSet) Delete(item blockatlas.Tx) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}
	return ok
}

func (s *TxSet) Has(item blockatlas.Tx) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.items[item]
	return ok
}

func (s *TxSet) Txs() []blockatlas.Tx {
	s.lock.RLock()
	defer s.lock.RUnlock()
	items := []blockatlas.Tx{}
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

// Size returns the size of the set
func (s *TxSet) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}
