package memory

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/observer"
	"reflect"
	"testing"
)

const ethCoin = coin.ETH
const addr1 = "0xde0B295669a9FD93d5F28D9Ec85E40f4cb697BAe"

var webhook1 = []string{"http://apple.com/push"}

const addr2 = "0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"

var webhook2 = []string{"http://trustwallet.com/webhook"}

const xpub = "zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC"

var xpubAddresses = []string{
	"bc1qrfr44n2j4czd5c9txwlnw0yj2h82x9566fglqj",
	"bc1qhn03cww757mnnlpkdvvfkaydxqygm86nvkm92h",
	"bc1q9n3l67ndzjjq4rat9rews35p3my4qa0h0pdlnv",
	"bc1q2fpry7zwqh575huc9urwfdvjtuvz508wez56ff",
	"bc1qk3yj6h79qw7tnsg4durc9sd5fpd3qt0p0m8u5p",
	"bc1qc7ekqf2t0elfsmtgr2mgd7da2up4vgq8uqk2nh",
	"bc1q6e8sdxlgc7ekqkqyevtrx8wshfv7sg66z3z6ce",
	"bc1q7nn4txus4g6fc5v7d2tha35ely8mfpd8qvv6eg",
	"bc1qv454wacvnenr3hzzldjqn8cgfltdlxwe96h737",
	"bc1qkl3rl62ekkp8s2m63hlw0rhcx8rj4g0jz8m6cg",
}

func TestMemoryStorage_Add(t *testing.T) {
	var observerMap = make(map[string]observer.Subscription)
	var storage observer.Storage = &Storage{
		observers: observerMap,
	}

	_ = storage.Add([]observer.Subscription{{
		Coin:     ethCoin,
		Address:  addr1,
		Webhooks: webhook1,
	}})

	if len(observerMap) != 1 {
		t.Error("observer not added")
	}
}

func TestMemoryStorage_List(t *testing.T) {
	var observerMap = make(map[string]observer.Subscription)
	var storage = &Storage{
		observers: observerMap,
	}

	obs1 := observer.Subscription{
		Coin:     ethCoin,
		Address:  addr1,
		Webhooks: webhook1,
	}
	obs2 := observer.Subscription{
		Coin:     ethCoin,
		Address:  addr2,
		Webhooks: webhook2,
	}

	observerMap[key(ethCoin, addr1)] = obs1
	observerMap[key(ethCoin, addr2)] = obs2

	if len(storage.List()) != 2 {
		t.Error("observers not listed properly")
	}
}

func TestMemoryStorage_Remove(t *testing.T) {
	var observerMap = make(map[string]observer.Subscription)
	var storage = &Storage{
		observers: observerMap,
	}

	obs := observer.Subscription{
		Coin:     ethCoin,
		Address:  addr1,
		Webhooks: webhook1,
	}
	observerMap[key(ethCoin, addr1)] = obs

	_ = storage.Delete([]observer.Subscription{obs})

	if len(storage.List()) != 0 {
		t.Error("observer not removed")
	}
}

func TestMemoryStorage_Get(t *testing.T) {
	var observerMap = make(map[string]observer.Subscription)
	var storage observer.Storage = &Storage{
		observers: observerMap,
	}

	obs1 := observer.Subscription{
		Coin:     ethCoin,
		Address:  addr1,
		Webhooks: webhook1,
	}
	obs2 := observer.Subscription{
		Coin:     ethCoin,
		Address:  addr2,
		Webhooks: webhook2,
	}

	observerMap[key(ethCoin, addr1)] = obs1
	observerMap[key(ethCoin, addr2)] = obs2

	res, _ := storage.Lookup(ethCoin, addr1)

	if !reflect.DeepEqual(res, []observer.Subscription{obs1}) {
		t.Error("wrong observer")
	}
}

func TestMemoryStorage_Contains(t *testing.T) {
	var observerMap = make(map[string]observer.Subscription)
	var storage = &Storage{
		observers: observerMap,
	}
	obs1 := observer.Subscription{
		Coin:     ethCoin,
		Address:  addr1,
		Webhooks: webhook1,
	}

	observerMap[key(ethCoin, addr1)] = obs1

	if yes, _ := storage.Contains(ethCoin, addr1); !yes {
		t.Errorf("observer should contain coint:%d address:%s", ethCoin, addr1)
	}

	if yes, _ := storage.Contains(ethCoin, addr2); yes {
		t.Errorf("observer should not contain coint:%d address:%s", ethCoin, addr2)
	}
}

func TestMemoryStorage_SaveAddresses(t *testing.T) {
	var addresses = make(map[string]string)
	var xAddresses = make(map[string][]string)
	var storage observer.Storage = &Storage{
		addresses:     addresses,
		xpubAddresses: xAddresses,
	}

	err := storage.SaveXpubAddresses(coin.BTC, xpubAddresses, xpub)
	if err != nil {
		t.Error(err)
	}
	if len(addresses) == 0 {
		t.Error("addresses not added")
	}
	if len(xAddresses) == 0 {
		t.Error("xpub not added")
	}
}

func TestMemoryStorage_GetAddresses(t *testing.T) {
	var addresses = make(map[string]string)
	var xAddresses = make(map[string][]string)
	var storage observer.Storage = &Storage{
		addresses:     addresses,
		xpubAddresses: xAddresses,
	}

	err := storage.SaveXpubAddresses(coin.BTC, xpubAddresses, xpub)
	if err != nil {
		t.Error(err)
	}
	if len(addresses) == 0 {
		t.Error("addresses not added")
	}
	res, err := storage.GetXpubFromAddress(coin.BTC, xpubAddresses[0])
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, xpub) {
		t.Error("wrong addresses")
	}
	a, err := storage.GetAddressFromXpub(coin.BTC, xpub)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(xpubAddresses, a) {
		t.Error("wrong xpub addresses")
	}
}
