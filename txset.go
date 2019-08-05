package blockatlas

import "sync"

type TxSet struct {
	items map[*Tx]bool
	lock  sync.RWMutex
}

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *TxSet) Add(t *Tx) *TxSet {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.items == nil {
		s.items = make(map[*Tx]bool)
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return s
}

func (s *TxSet) Txs() []Tx {
	s.lock.RLock()
	defer s.lock.RUnlock()
	items := []Tx{}
	for i := range s.items {
		items = append(items, *i)
	}
	return items
}

func (s *TxSet) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}
