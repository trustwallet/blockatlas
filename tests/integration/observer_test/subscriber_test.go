// +build integration

package observer_test

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/subscriber"
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

		go mq.RunConsumerForChannelWithCancelAndDbConn(subscriber.RunTransactionsSubscriber, subscriptionChannel, database, ctx)
		time.Sleep(time.Second * 2)
		cancel()
	}

	for _, wanted := range wantedEvents {
		result, err := database.GetSubscriptionsForNotifications([]string{wanted.Address}, context.Background())
		assert.Nil(t, err)
		assert.Equal(t, result[0].Address.Address, wanted.Address)
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

	database.AddSubscriptionsForNotifications(
		[]string{"61_0x0000000000000000000000000000000000000000"}, context.Background())

	database.AddSubscriptionsForNotifications(
		[]string{"62_0x0000000000000000000000000000000000000000"}, context.Background())

	database.AddSubscriptionsForNotifications(
		[]string{"63_0x0000000000000000000000000000000000000000"}, context.Background())

	database.AddSubscriptionsForNotifications(
		[]string{"64_0x0000000000000000000000000000000000000000"}, context.Background())

	for _, event := range givenEvents {
		body, err := json.Marshal(event)
		assert.Nil(t, err)

		err = mq.Subscriptions.Publish(body)
		assert.Nil(t, err)

		ctx, cancel := context.WithCancel(context.Background())

		go mq.RunConsumerForChannelWithCancelAndDbConn(subscriber.RunTransactionsSubscriber, subscriptionChannel, database, ctx)
		time.Sleep(time.Second)
		cancel()
	}

	for _, wanted := range wantedEvents {
		result, err := database.GetSubscriptionsForNotifications([]string{wanted.Address}, context.Background())
		assert.Nil(t, err)
		assert.Len(t, result, 1)
	}

	abs61, err := database.GetSubscriptionsForNotifications([]string{"61_0x0000000000000000000000000000000000000000"}, context.Background())
	assert.Nil(t, err)
	assert.Len(t, abs61, 1)

	abs62, err := database.GetSubscriptionsForNotifications([]string{"62_0x0000000000000000000000000000000000000000"}, context.Background())
	assert.Nil(t, err)
	assert.Len(t, abs62, 1)

	abs63, err := database.GetSubscriptionsForNotifications([]string{"63_0x0000000000000000000000000000000000000000"}, context.Background())
	assert.Nil(t, err)
	assert.Len(t, abs63, 1)

	abs64, err := database.GetSubscriptionsForNotifications([]string{"64_0x0000000000000000000000000000000000000000"}, context.Background())
	assert.Nil(t, err)
	assert.Len(t, abs64, 1)

	abs65, err := database.GetSubscriptionsForNotifications([]string{"65_0x0000000000000000000000000000000000000000"}, context.Background())
	assert.Nil(t, err)
	assert.Len(t, abs65, 0)
}
