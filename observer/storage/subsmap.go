package storage

import (
	"sync"
)

type SubsMap struct {
	subs     map[string][]Subscription
	xpubSubs map[string][]Subscription
	lock     sync.RWMutex
}

func (sm *SubsMap) AddSubscriptions(subs []interface{}) {
	for _, sub := range subs {
		s, ok := sub.(Subscription)
		if ok {
			sm.AddSubscription(s)
		}
	}
}

func (sm *SubsMap) DeleteSubscriptions(subs []interface{}) {
	for _, sub := range subs {
		s, ok := sub.(Subscription)
		if ok {
			sm.DeleteSubscription(s)
		}
	}
}

func (sm *SubsMap) DeleteSubscription(sub Subscription) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	subs, ok := sm.subs[sub.Address]
	if !ok || subs == nil {
		return
	}

	newSubs := make([]Subscription, 0)
	for _, s := range subs {
		if s.Equal(sub) {
			continue
		}
		newSubs = append(newSubs, s)
	}
	sm.subs[sub.Address] = newSubs
}

func (sm *SubsMap) AddSubscription(sub Subscription) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	s, ok := sm.subs[sub.Address]
	if !ok {
		s = make([]Subscription, 0)
	}
	sm.subs[sub.Address] = append(s, sub)

	xs, ok := sm.xpubSubs[sub.Xpub.String]
	if !ok {
		xs = make([]Subscription, 0)
	}
	sm.subs[sub.Xpub.String] = append(xs, sub)
}

func (sm *SubsMap) GetSubscription(address string) ([]Subscription, bool) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	s, ok := sm.subs[address]
	if !ok || s == nil || len(s) == 0 {
		return s, ok
	}
	xs, ok := sm.xpubSubs[s[0].Xpub.String]
	if ok {
		s = append(s, xs...)
	}
	return s, true
}

func (sm *SubsMap) GetSubscriptions() map[string][]Subscription {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	return sm.subs
}
