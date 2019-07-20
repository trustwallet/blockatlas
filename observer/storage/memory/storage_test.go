package memory

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/observer"
	"reflect"
	"testing"
)

const ethCoin = coin.ETH
const addr1 = "0xde0B295669a9FD93d5F28D9Ec85E40f4cb697BAe"
const webhook1 = "http://apple.com/push"

const addr2 = "0xEA674fdDe714fd979de3EdF0F56AA9716B898ec8"
const webhook2 = "http://trustwallet.com/webhook"

func TestMemoryStorage_Add(t *testing.T) {
	var observerMap = make(map[string][]string)
	var storage observer.Storage = &Client{
		observers: observerMap,
	}

	_ = storage.Add([]observer.Subscription{
		{
			Coin:    ethCoin,
			Address: addr1,
			WebHook: webhook1,
		},
		{
			Coin:    ethCoin,
			Address: addr1,
			WebHook: webhook2,
		},
		{
			Coin:    ethCoin,
			Address: addr2,
			WebHook: webhook2,
		},
	})

	if len(observerMap) != 2 {
		t.Error("incorrect map")
	}
	if len(observerMap[encodeKey(ethCoin, addr1)]) != 2 {
		t.Error("incorrect hook list")
	}
	if len(observerMap[encodeKey(ethCoin, addr2)]) != 1 {
		t.Error("incorrect hook list")
	}
}

func TestMemoryStorage_List(t *testing.T) {
	var observerMap = make(map[string][]string)
	var storage = &Client{
		observers: observerMap,
	}

	observerMap[encodeKey(ethCoin, addr1)] = []string{webhook1, webhook2}
	observerMap[encodeKey(ethCoin, addr2)] = []string{webhook2}

	if len(storage.List()) != 3 {
		t.Error("observers not listed properly")
	}
}

func TestMemoryStorage_Remove(t *testing.T) {
	var observerMap = make(map[string][]string)
	var storage = &Client{
		observers: observerMap,
	}

	obs := observer.Subscription{
		Coin:    ethCoin,
		Address: addr1,
		WebHook: webhook1,
	}
	observerMap[encodeKey(ethCoin, addr1)] = []string{webhook1}

	_ = storage.Delete([]observer.Subscription{obs})

	if len(storage.List()) != 0 {
		t.Error("observer not removed")
	}
}

func TestMemoryStorage_Get(t *testing.T) {
	var observerMap = make(map[string][]string)
	var storage observer.Storage = &Client{
		observers: observerMap,
	}

	obs1 := observer.Subscription{
		Coin:    ethCoin,
		Address: addr1,
		WebHook: webhook1,
	}

	observerMap[encodeKey(ethCoin, addr1)] = []string{webhook1}
	observerMap[encodeKey(ethCoin, addr2)] = []string{webhook2}

	res, _ := storage.Lookup(ethCoin, addr1)

	if !reflect.DeepEqual(res, []observer.Subscription{obs1}) {
		t.Error("wrong observer")
	}
}

func TestMemoryStorage_Contains(t *testing.T) {
	var observerMap = make(map[string][]string)
	var storage = &Client{
		observers: observerMap,
	}

	observerMap[encodeKey(ethCoin, addr1)] = []string{webhook1}

	if yes, _ := storage.Contains(ethCoin, addr1); !yes {
		t.Errorf("observer should contain coint:%d address:%s", ethCoin, addr1)
	}

	if yes, _ := storage.Contains(ethCoin, addr2); yes {
		t.Errorf("observer should not contain coint:%d address:%s", ethCoin, addr2)
	}
}
