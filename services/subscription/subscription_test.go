package subscription

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"testing"
)

func TestGetInterval(t *testing.T) {
	cache := storage.New()
	err := cache.Init("redis://localhost:6379")
	if err != nil {
		logger.Fatal(err)
	}

	var subscriptions = blockatlas.Subscriptions{
		"2": {"COINADDRESS 0"},
		"0": {"ADDRES 1", "aDDRESS 2"},
	}

	var mockedData = blockatlas.SubscriptionEvent{
		Subscriptions: subscriptions,
		GUID:          "EXAMPLE",
		Operation:     addSubscription,
	}

	b, _ := json.Marshal(mockedData)

	success := Consume(b, cache)
	if !success {
		t.Fatal()
	}
}
