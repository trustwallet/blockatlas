// build integration

package db_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"testing"
)

func TestDb_AddSubscriptionsBulk(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)
	var subscriptions []string

	for i := 0; i < 100; i++ {
		subscriptions = append(subscriptions, "testAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddr")
	}

	assert.Nil(t, database.AddSubscriptionsForNotifications(subscriptions, context.Background()))
	for i := 0; i < 100; i++ {
		s, err := database.GetSubscriptionsForNotifications([]string{"testAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddr"}, context.Background())
		assert.Nil(t, err)
		assert.NotNil(t, s)
	}
}

func TestDb_AddSubscriptions(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)

	assert.Nil(t, database.AddSubscriptionsForNotifications([]string{"60_testAddr", "60_testAddr2", "60_testAddr3"}, context.Background()))

	subs, err := database.GetSubscriptionsForNotifications([]string{"60_testAddr"}, context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, subs)
	assert.Equal(t, 1, len(subs))
	assert.Equal(t, "60_testAddr", subs[0].Address.Address)

	subs, err = database.GetSubscriptionsForNotifications([]string{"60_testAddr2"}, context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, subs)
	assert.Equal(t, 1, len(subs))
	assert.Equal(t, "60_testAddr2", subs[0].Address.Address)

	subs, err = database.GetSubscriptionsForNotifications([]string{"60_testAddr3"}, context.Background())
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

	assert.Nil(t, database.AddSubscriptions(subscriber.ToSubscriptionData(subscriptionsA), context.Background()))

	var subscriptionsB []blockatlas.Subscription

	for _, sub := range subscriptionsA {
		subscriptionsB = append(subscriptionsB, sub)
	}
	assert.Nil(t, database.AddSubscriptions(subscriber.ToSubscriptionData(subscriptionsB), context.Background()))

	returnedSubs, err := database.GetSubscriptions(60, []string{"etherAddress"}, context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(returnedSubs))

	returnedSubs, err = database.GetSubscriptions(714, []string{"binanceAddress"}, context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(returnedSubs))

	returnedSubs, err = database.GetSubscriptions(144, []string{"XLMAddress"}, context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(returnedSubs))

	returnedSubs, err = database.GetSubscriptions(148, []string{"AtomAddress"}, context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(returnedSubs))

	returnedSubs, err = database.GetSubscriptions(61, []string{"ETCAddress"}, context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(returnedSubs))
}

//func TestDb_DeleteSubscriptions(t *testing.T) {
//	setup.CleanupPgContainer(database.Gorm)
//	var subscriptions []models.NotificationSubscription
//
//	subscriptions = append(subscriptions, models.NotificationSubscription{
//		Coin:    60,
//		Address: "testAddr",
//	})
//
//	subscriptions = append(subscriptions, models.NotificationSubscription{
//		Coin:    714,
//		Address: "testAddr2",
//	})
//
//	subscriptions = append(subscriptions, models.NotificationSubscription{
//		Coin:    144,
//		Address: "testAddr3",
//	})
//
//	assert.Nil(t, database.AddSubscriptions(subscriptions, context.Background()))
//
//	subs60, err := database.GetSubscriptions(60, []string{"testAddr"}, context.Background())
//	assert.Nil(t, err)
//	assert.NotNil(t, subs60)
//	assert.Equal(t, 1, len(subs60))
//	assert.Equal(t, subscriptions[0].Coin, subs60[0].Coin)
//	assert.Equal(t, subscriptions[0].Address, subs60[0].Address)
//
//	subs714, err := database.GetSubscriptions(714, []string{"testAddr2"}, context.Background())
//	assert.Nil(t, err)
//	assert.NotNil(t, subs714)
//	assert.Equal(t, 1, len(subs714))
//	assert.Equal(t, subscriptions[1].Coin, subs714[0].Coin)
//	assert.Equal(t, subscriptions[1].Address, subs714[0].Address)
//
//	subs144, err := database.GetSubscriptions(144, []string{"testAddr3"}, context.Background())
//	assert.Nil(t, err)
//	assert.NotNil(t, subs144)
//	assert.Equal(t, 1, len(subs144))
//	assert.Equal(t, subscriptions[2].Coin, subs144[0].Coin)
//	assert.Equal(t, subscriptions[2].Address, subs144[0].Address)
//
//	subsToDel := []models.NotificationSubscription{subscriptions[0]}
//
//	assert.Nil(t, database.DeleteSubscriptions(subsToDel, context.Background()))
//
//	subs714N2, err := database.GetSubscriptions(714, []string{"testAddr2"}, context.Background())
//	assert.Nil(t, err)
//	assert.NotNil(t, subs714N2)
//	assert.Equal(t, 1, len(subs714N2))
//	assert.Equal(t, subscriptions[1].Coin, subs714N2[0].Coin)
//	assert.Equal(t, subscriptions[1].Address, subs714N2[0].Address)
//
//	subs144N2, err := database.GetSubscriptions(144, []string{"testAddr3"}, context.Background())
//	assert.Nil(t, err)
//	assert.NotNil(t, subs144N2)
//	assert.Equal(t, 1, len(subs144N2))
//	assert.Equal(t, subscriptions[2].Coin, subs144N2[0].Coin)
//	assert.Equal(t, subscriptions[2].Address, subs144N2[0].Address)
//
//	subs60N2, err := database.GetSubscriptions(60, []string{"testAddr"}, context.Background())
//	assert.Nil(t, err)
//	assert.Len(t, subs60N2, 0)
//}
//
//func TestDeleteAll(t *testing.T) {
//	setup.CleanupPgContainer(database.Gorm)
//	var subscriptions []models.NotificationSubscription
//
//	subscriptions = append(subscriptions, models.NotificationSubscription{
//		Coin:    60,
//		Address: "testAddr",
//	})
//
//	subscriptions = append(subscriptions, models.NotificationSubscription{
//		Coin:    714,
//		Address: "testAddr2",
//	})
//
//	subscriptions = append(subscriptions, models.NotificationSubscription{
//		Coin:    144,
//		Address: "testAddr3",
//	})
//	assert.Nil(t, database.AddSubscriptions(subscriptions, context.Background()))
//
//	subs60, err := database.GetSubscriptions(60, []string{"testAddr"}, context.Background())
//	assert.Nil(t, err)
//	assert.Len(t, subs60, 1)
//
//	subs714, err := database.GetSubscriptions(714, []string{"testAddr2"}, context.Background())
//	assert.Nil(t, err)
//	assert.Len(t, subs714, 1)
//
//	subs144, err := database.GetSubscriptions(144, []string{"testAddr3"}, context.Background())
//	assert.Nil(t, err)
//	assert.Len(t, subs144, 1)
//}
//
//func TestDb_DuplicateEntries(t *testing.T) {
//	setup.CleanupPgContainer(database.Gorm)
//	var subscriptions []models.NotificationSubscription
//
//	for i := 0; i < 10; i++ {
//		subscriptions = append(subscriptions, models.NotificationSubscription{
//			Coin:    60,
//			Address: "testAddr",
//		})
//	}
//
//	assert.Nil(t, database.AddSubscriptions(subscriptions, context.Background()))
//
//	subs, err := database.GetSubscriptions(60, []string{"testAddr"}, context.Background())
//	assert.Nil(t, err)
//	assert.NotNil(t, subs)
//	assert.Equal(t, 1, len(subs))
//	assert.True(t, containSub(subscriptions[0], subs))
//}
//
//func TestDb_CreateDeleteCreate(t *testing.T) {
//	setup.CleanupPgContainer(database.Gorm)
//	var subscriptions []models.NotificationSubscription
//	subscriptions = append(subscriptions, models.NotificationSubscription{
//		Coin:    60,
//		Address: "testAddr",
//	})
//
//	assert.Nil(t, database.AddSubscriptions(subscriptions, context.Background()))
//	subs, err := database.GetSubscriptions(60, []string{"testAddr"}, context.Background())
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(subs))
//
//	time.Sleep(time.Second)
//
//	assert.Nil(t, database.DeleteSubscriptions(subs, context.Background()))
//
//	subs2, err := database.GetSubscriptions(60, []string{"testAddr"}, context.Background())
//	assert.Nil(t, err)
//	assert.Equal(t, 0, len(subs2))
//
//	assert.Nil(t, database.AddSubscriptions(subscriptions, context.Background()))
//
//	subs3, err := database.GetSubscriptions(60, []string{"testAddr"}, context.Background())
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(subs3))
//}
//
//func TestDb_UpdatedAt(t *testing.T) {
//	setup.CleanupPgContainer(database.Gorm)
//	var subscriptions []models.NotificationSubscription
//	subscriptions = append(subscriptions, models.NotificationSubscription{
//		Coin:    60,
//		Address: "testAddr",
//	})
//
//	assert.Nil(t, database.AddSubscriptions(subscriptions, context.Background()))
//	subs, err := database.GetSubscriptions(60, []string{"testAddr"}, context.Background())
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(subs))
//
//	time.Sleep(time.Second)
//
//	var existingSub models.NotificationSubscription
//	assert.False(t, database.Gorm.Where(models.NotificationSubscription{Address: "testAddr"}).First(&existingSub).RecordNotFound())
//	assert.Greater(t, time.Now().Unix(), existingSub.CreatedAt.Unix())
//	assert.Greater(t, existingSub.CreatedAt.Unix(), time.Now().Unix()-120)
//
//	subscriptions = append(subscriptions, models.NotificationSubscription{
//		Coin:    714,
//		Address: "newtestAddr",
//	})
//
//	assert.Nil(t, database.AddSubscriptions(subscriptions, context.Background()))
//
//	time.Sleep(time.Second)
//
//	var existingSub2 models.NotificationSubscription
//	assert.False(t, database.Gorm.Where(models.NotificationSubscription{Address: "testAddr"}).First(&existingSub2).RecordNotFound())
//
//	assert.Greater(t, time.Now().Unix(), existingSub2.CreatedAt.Unix())
//	assert.Greater(t, existingSub2.CreatedAt.Unix(), time.Now().Unix()-120)
//	assert.GreaterOrEqual(t, existingSub2.CreatedAt.Unix(), existingSub.CreatedAt.Unix())
//}
//
//func containSub(sub models.NotificationSubscription, list []models.NotificationSubscription) bool {
//	for _, s := range list {
//		if sub.Address == s.Address && sub.Coin == s.Coin {
//			return true
//		}
//	}
//	return false
//}
