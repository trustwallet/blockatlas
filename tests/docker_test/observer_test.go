// +build integration

package docker_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/observer/subscriber"
	"github.com/trustwallet/blockatlas/storage"
	"github.com/trustwallet/blockatlas/tests/docker_test/setup"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func TestSubscriberAddSubscription(t *testing.T) {
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

	assert.Nil(t, mq.Subscriptions.Declare())

	for _, event := range givenEvents {
		body, err := json.Marshal(event)
		assert.Nil(t, err)

		err = mq.Subscriptions.Publish(body)
		assert.Nil(t, err)

		go mq.Subscriptions.RunConsumer(subscriber.RunSubscriber, setup.Cache)
		time.Sleep(time.Second / 5)
	}

	result, err := setup.Cache.GetAllHM(storage.ATLAS_OBSERVER)
	assert.Nil(t, err)
	assert.NotNil(t, result)

	for _, wanted := range wantedEvents {
		result, err := setup.Cache.FindSubscriptions(wanted.Coin, []string{wanted.Address})
		assert.Nil(t, err)
		assert.Equal(t, result[0], wanted)
	}
}
