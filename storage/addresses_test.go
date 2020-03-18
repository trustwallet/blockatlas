package storage

import (
	"fmt"
	"github.com/alicebob/miniredis"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"testing"
)

func TestStorage_Lookup(t *testing.T) {
	s := initStorage(t)

	type fields struct {
		coin      int
		addresses []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []blockatlas.Subscription
	}{
		{"test fee transfer",
			fields{coin: 60, addresses: []string{"1", "2", "3"}},
			[]blockatlas.Subscription{
				{Coin: 60, Address: "1", GUID: "1"},
				{Coin: 60, Address: "2", GUID: "2"},
				{Coin: 60, Address: "3", GUID: "3"},
			},
		},
	}

	for _, tt := range tests {
		for i, a := range tt.fields.addresses {
			key := getSubscriptionKey(uint(tt.fields.coin), a)
			err := s.AddHM(ATLAS_OBSERVER, key, []string{tt.want[i].GUID})
			if err != nil {
				t.Fatal(err)
			}
		}

		t.Run(tt.name, func(t *testing.T) {
			if got, err := s.FindSubscriptions(uint(tt.fields.coin), tt.fields.addresses); !isEqual(got, tt.want) || err != nil {
				t.Fatal(got, tt.want)
			}
		})
	}
}

func isEqual(given, want []blockatlas.Subscription) bool {
	if len(given) != len(want) {
		return false
	}
	var givenCounter int
	for _, g := range given {
		var wantCounter int
		for _, w := range want {
			if w == g {
				wantCounter++
			}
		}
		if wantCounter > 0 {
			givenCounter++
		}
	}

	if givenCounter == len(want) {
		return true
	}
	return false
}

func initStorage(t *testing.T) Storage {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}

	storage := New()
	err = storage.Init(fmt.Sprintf("redis://%s", s.Addr()))
	if err != nil {
		logger.Fatal(err)
	}
	return *storage
}
