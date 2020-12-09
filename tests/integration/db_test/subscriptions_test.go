// +build integration

package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/subscriber"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
)

func TestDb_AddSubscriptionsBulk(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)
	var subscriptions []string

	for i := 0; i < 100; i++ {
		subscriptions = append(subscriptions, "testAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddr")
	}

	assert.Nil(t, database.AddSubscriptionsForNotifications(subscriptions))
	for i := 0; i < 100; i++ {
		s, err := database.GetSubscriptionsForNotifications([]string{"testAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddr"})
		assert.Nil(t, err)
		assert.NotNil(t, s)
	}
}

func TestDb_AddSubscriptions(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)

	assert.Nil(t, database.AddSubscriptionsForNotifications([]string{"60_testAddr", "60_testAddr2", "60_testAddr3"}))

	subs, err := database.GetSubscriptionsForNotifications([]string{"60_testAddr"})
	assert.Nil(t, err)
	assert.NotNil(t, subs)
	assert.Equal(t, 1, len(subs))
	assert.Equal(t, "60_testAddr", subs[0].Address.Address)

	subs, err = database.GetSubscriptionsForNotifications([]string{"60_testAddr2"})
	assert.Nil(t, err)
	assert.NotNil(t, subs)
	assert.Equal(t, 1, len(subs))
	assert.Equal(t, "60_testAddr2", subs[0].Address.Address)

	subs, err = database.GetSubscriptionsForNotifications([]string{"60_testAddr3"})
	assert.Nil(t, err)
	assert.NotNil(t, subs)
	assert.Equal(t, 1, len(subs))
	assert.Equal(t, "60_testAddr3", subs[0].Address.Address)
}

func TestDb_FindSubscriptions(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)

	var subscriptionsA []blockatlas.Subscription

	subscriptionsA = append(subscriptionsA, blockatlas.Subscription{
		Coin:    60,
		Address: "etherAddress",
	})

	subscriptionsA = append(subscriptionsA, blockatlas.Subscription{
		Coin:    714,
		Address: "binanceAddress",
	})

	subscriptionsA = append(subscriptionsA, blockatlas.Subscription{
		Coin:    148,
		Address: "AtomAddress",
	})

	subscriptionsA = append(subscriptionsA, blockatlas.Subscription{
		Coin:    144,
		Address: "XLMAddress",
	})

	subscriptionsA = append(subscriptionsA, blockatlas.Subscription{
		Coin:    61,
		Address: "ETCAddress",
	})

	assert.Nil(t, database.AddSubscriptionsForNotifications(subscriber.ToSubscriptionData(subscriptionsA)))

	var subscriptionsB []blockatlas.Subscription

	for _, sub := range subscriptionsA {
		subscriptionsB = append(subscriptionsB, sub)
	}
	assert.Nil(t, database.AddSubscriptionsForNotifications(subscriber.ToSubscriptionData(subscriptionsB)))

	returnedSubs, err := database.GetSubscriptionsForNotifications([]string{"60_etherAddress"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(returnedSubs))

	returnedSubs, err = database.GetSubscriptionsForNotifications([]string{"714_binanceAddress"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(returnedSubs))

	returnedSubs, err = database.GetSubscriptionsForNotifications([]string{"144_XLMAddress"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(returnedSubs))

	returnedSubs, err = database.GetSubscriptionsForNotifications([]string{"148_AtomAddress"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(returnedSubs))

	returnedSubs, err = database.GetSubscriptionsForNotifications([]string{"61_ETCAddress"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(returnedSubs))
}

func TestDb_DeleteSubscriptions(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)
	subscriptions := []string{
		"60_testAddr",
		"714_testAddr2",
		"144_testAddr3",
	}

	assert.Nil(t, database.AddSubscriptionsForNotifications(subscriptions))

	subs60, err := database.GetSubscriptionsForNotifications([]string{"60_testAddr"})
	assert.Nil(t, err)
	assert.NotNil(t, subs60)
	assert.Equal(t, 1, len(subs60))

	subs714, err := database.GetSubscriptionsForNotifications([]string{"714_testAddr2"})
	assert.Nil(t, err)
	assert.NotNil(t, subs714)
	assert.Equal(t, 1, len(subs714))

	subs144, err := database.GetSubscriptionsForNotifications([]string{"144_testAddr3"})
	assert.Nil(t, err)
	assert.NotNil(t, subs144)
	assert.Equal(t, 1, len(subs144))

	assert.Nil(t, database.DeleteSubscriptionsForNotifications([]string{subscriptions[0]}))

	subs714N2, err := database.GetSubscriptionsForNotifications([]string{"714_testAddr2"})
	assert.Nil(t, err)
	assert.NotNil(t, subs714N2)
	assert.Equal(t, 1, len(subs714N2))

	subs144N2, err := database.GetSubscriptionsForNotifications([]string{"144_testAddr3"})
	assert.Nil(t, err)
	assert.NotNil(t, subs144N2)
	assert.Equal(t, 1, len(subs144N2))

	subs60N2, err := database.GetSubscriptionsForNotifications([]string{"60_testAddr"})
	assert.Nil(t, err)
	assert.Len(t, subs60N2, 0)
}

func TestDb_DuplicateEntries(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)
	var subscriptions []blockatlas.Subscription

	for i := 0; i < 10; i++ {
		subscriptions = append(subscriptions, blockatlas.Subscription{
			Coin:    60,
			Address: "testAddr",
		})
	}

	assert.Nil(t, database.AddSubscriptionsForNotifications(subscriber.ToSubscriptionData(subscriptions)))

	subs, err := database.GetSubscriptionsForNotifications([]string{"60_testAddr"})
	assert.Nil(t, err)
	assert.NotNil(t, subs)
	assert.Equal(t, 1, len(subs))
}

func TestDb_CreateDeleteCreate(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)
	var subscriptions []blockatlas.Subscription
	subscriptions = append(subscriptions, blockatlas.Subscription{
		Coin:    60,
		Address: "testAddr",
	})

	assert.Nil(t, database.AddSubscriptionsForNotifications(subscriber.ToSubscriptionData(subscriptions)))
	subs, err := database.GetSubscriptionsForNotifications([]string{"60_testAddr"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(subs))

	assert.Nil(t, database.DeleteSubscriptionsForNotifications([]string{"60_testAddr"}))

	subs2, err := database.GetSubscriptionsForNotifications([]string{"60_testAddr"})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(subs2))

	assert.Nil(t, database.AddSubscriptionsForNotifications(subscriber.ToSubscriptionData(subscriptions)))

	subs3, err := database.GetSubscriptionsForNotifications([]string{"60_testAddr"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(subs3))
}
