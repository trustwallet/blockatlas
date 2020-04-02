// +build integration

package observer_test

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/observer/subscriber"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func TestSubscriberAddSubscription(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)

	_, goFile, _, _ := runtime.Caller(0)
	testFilePathGiven := filepath.Join(filepath.Dir(goFile), "data", "given_subscriptions_added.json")
	testFileGiven, err := ioutil.ReadFile(testFilePathGiven)
	if err != nil {
		t.Fatal(err)
	}
	var givenEvents []blockatlas.SubscriptionEvent
	if err := json.Unmarshal(testFileGiven, &givenEvents); err != nil {
		t.Fatal(err)
	}

	testFilePathWanted := filepath.Join(filepath.Dir(goFile), "data", "wanted_subscriptions_added.json")
	testFileWanted, err := ioutil.ReadFile(testFilePathWanted)
	if err != nil {
		t.Fatal(err)
	}
	var wantedEvents []blockatlas.Subscription
	if err := json.Unmarshal(testFileWanted, &wantedEvents); err != nil {
		t.Fatal(err)
	}

	for _, event := range givenEvents {
		body, err := json.Marshal(event)
		assert.Nil(t, err)

		err = mq.Subscriptions.Publish(body)
		assert.Nil(t, err)

		ctx, cancel := context.WithCancel(context.Background())

		go mq.RunConsumerForChannelWithCancelAndDbConn(subscriber.RunSubscriber, subscriptionChannel, database, ctx)
		time.Sleep(time.Second * 2)
		cancel()
	}

	for _, wanted := range wantedEvents {
		result, err := database.GetSubscriptionData(wanted.Coin, []string{wanted.Address})
		assert.Nil(t, err)
		assert.Equal(t, result[0].SubscriptionId, wanted.Id)
		assert.Equal(t, result[0].Coin, wanted.Coin)
		assert.Equal(t, result[0].Address, wanted.Address)
	}
}

func TestSubscriber_UpdateSubscription(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)
	_, goFile, _, _ := runtime.Caller(0)
	testFilePathGiven := filepath.Join(filepath.Dir(goFile), "data", "given_subscriptions_deleted.json")
	testFileGiven, err := ioutil.ReadFile(testFilePathGiven)
	if err != nil {
		t.Fatal(err)
	}
	var givenEvents []blockatlas.SubscriptionEvent
	if err := json.Unmarshal(testFileGiven, &givenEvents); err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(testFileGiven, &givenEvents); err != nil {
		t.Fatal(err)
	}

	testFilePathWanted := filepath.Join(filepath.Dir(goFile), "data", "wanted_subscriptions_added.json")
	testFileWanted, err := ioutil.ReadFile(testFilePathWanted)
	if err != nil {
		t.Fatal(err)
	}
	var wantedEvents []blockatlas.Subscription
	if err := json.Unmarshal(testFileWanted, &wantedEvents); err != nil {
		t.Fatal(err)
	}

	database.AddSubscriptions(10, []models.SubscriptionData{
		{Coin: 61, Address: "0x0000000000000000000000000000000000000000", SubscriptionId: 10},
	})
	database.AddSubscriptions(1, []models.SubscriptionData{
		{Coin: 62, Address: "0x0000000000000000000000000000000000000000", SubscriptionId: 1},
	})
	database.AddSubscriptions(2, []models.SubscriptionData{
		{Coin: 63, Address: "0x0000000000000000000000000000000000000000", SubscriptionId: 2},
	})
	database.AddSubscriptions(3, []models.SubscriptionData{
		{Coin: 64, Address: "0x0000000000000000000000000000000000000000", SubscriptionId: 3},
	})

	for _, event := range givenEvents {
		body, err := json.Marshal(event)
		assert.Nil(t, err)

		err = mq.Subscriptions.Publish(body)
		assert.Nil(t, err)

		ctx, cancel := context.WithCancel(context.Background())

		go mq.RunConsumerForChannelWithCancelAndDbConn(subscriber.RunSubscriber, subscriptionChannel, database, ctx)
		time.Sleep(time.Second)
		cancel()
	}

	for _, wanted := range wantedEvents {
		result, err := database.GetSubscriptionData(wanted.Coin, []string{wanted.Address})
		assert.Nil(t, err)
		assert.Equal(t, result[0].SubscriptionId, wanted.Id)
		assert.Equal(t, result[0].Coin, wanted.Coin)
		assert.Equal(t, result[0].Address, wanted.Address)

	}

	abs61, err := database.GetSubscriptionData(61, []string{"0x0000000000000000000000000000000000000000"})
	assert.Nil(t, err)
	assert.Len(t, abs61, 0)

	abs62, err := database.GetSubscriptionData(62, []string{"0x0000000000000000000000000000000000000000"})
	assert.Nil(t, err)
	assert.Len(t, abs62, 0)

	abs63, err := database.GetSubscriptionData(63, []string{"0x0000000000000000000000000000000000000000"})
	assert.Nil(t, err)
	assert.Len(t, abs63, 0)

	abs64, err := database.GetSubscriptionData(64, []string{"0x0000000000000000000000000000000000000000"})
	assert.Nil(t, err)
	assert.Len(t, abs64, 0)
}
