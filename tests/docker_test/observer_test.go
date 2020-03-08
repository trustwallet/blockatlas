// +build integration

package docker_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/subscription"
	"github.com/trustwallet/blockatlas/storage"
	"github.com/trustwallet/blockatlas/tests/docker_test/setup"
	"testing"
	"time"
)

func TestSubscriber(t *testing.T){
	assert.Nil(t, mq.Subscriptions.Declare())
	event := blockatlas.SubscriptionEvent{
		NewSubscriptions: blockatlas.Subscriptions{
			"60":[]string{"0x123"},
		},
		GUID:             "1",
		Operation:        subscription.AddSubscription,
	}

	body, err := json.Marshal(event)
	assert.Nil(t, err)

	err = mq.Subscriptions.Publish(body)
	assert.Nil(t, err)

	go mq.Subscriptions.RunConsumer(subscription.Consume, setup.Cache)
	time.Sleep(time.Second * 3)

	result, err := setup.Cache.GetAllHM(storage.ATLAS_OBSERVER)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}