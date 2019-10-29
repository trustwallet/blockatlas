package storage

import (
	"sync"
)

type XpubMap struct {
	xpub map[string][]string
	lock sync.RWMutex
}

func (xm *XpubMap) GetXpubAddresses(xpub string) ([]string, bool) {
	xm.lock.RLock()
	defer xm.lock.RUnlock()
	s, ok := xm.xpub[xpub]
	if !ok || s == nil || len(s) == 0 {
		return s, ok
	}
	return s, true
}

func (xm *XpubMap) GetXpub(xpub string) ([]string, bool) {
	xm.lock.RLock()
	defer xm.lock.RUnlock()
	a, ok := xm.xpub[xpub]
	return a, ok
}

func (xm *XpubMap) GetXpubFromAddress(address string) (string, bool) {
	xm.lock.RLock()
	defer xm.lock.RUnlock()
	for xpub, addresses := range xm.xpub {
		for _, addr := range addresses {
			if addr == address {
				return xpub, true
			}
		}
	}
	return "", true
}

func (xm *XpubMap) SetXpub(xpub string, addresses []string) {
	xm.lock.Lock()
	defer xm.lock.Unlock()
	xm.xpub[xpub] = addresses
}

func (xm *XpubMap) SetXpubs(xpubs map[string][]string) {
	xm.lock.Lock()
	defer xm.lock.Unlock()
	for xpub, addresses := range xpubs {
		xm.xpub[xpub] = addresses
	}
}

func (xm *XpubMap) GetXpubs() map[string][]string {
	xm.lock.RLock()
	defer xm.lock.RUnlock()
	return xm.xpub
}
