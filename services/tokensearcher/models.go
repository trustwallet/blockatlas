package tokensearcher

import "sync"

type assetsByAddresses struct {
	Result map[string][]string
	sync.Mutex
}
