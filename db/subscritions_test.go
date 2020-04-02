package db

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/db/models"
	"testing"
)

func TestGetSubscriptionsToDeleteAndUpdate(t *testing.T) {
	oldSubscriptions := []models.SubscriptionData{{
		SubscriptionId: 1,
		Coin:           60,
		Address:        "A",
	}, {
		SubscriptionId: 1,
		Coin:           60,
		Address:        "B",
	}, {
		SubscriptionId: 1,
		Coin:           60,
		Address:        "C",
	}}

	newSubscription := []models.SubscriptionData{{
		SubscriptionId: 1,
		Coin:           60,
		Address:        "B",
	}, {
		SubscriptionId: 1,
		Coin:           60,
		Address:        "C",
	}}

	update, delete := getSubscriptionsToDeleteAndUpdate(oldSubscriptions, newSubscription)
	assert.Len(t, update, 0)
	assert.Len(t, delete, 1)
	assert.Equal(t, "A", delete[0].Address)

	oldSubscriptions = []models.SubscriptionData{{
		SubscriptionId: 1,
		Coin:           60,
		Address:        "A",
	}, {
		SubscriptionId: 1,
		Coin:           60,
		Address:        "B",
	}, {
		SubscriptionId: 1,
		Coin:           60,
		Address:        "C",
	}, {
		SubscriptionId: 1,
		Coin:           60,
		Address:        "D",
	}}

	newSubscription = []models.SubscriptionData{{
		SubscriptionId: 1,
		Coin:           60,
		Address:        "E",
	}}

	update, delete = getSubscriptionsToDeleteAndUpdate(oldSubscriptions, newSubscription)
	assert.Len(t, update, 1)
	assert.Len(t, delete, 4)

	oldSubscriptions = []models.SubscriptionData{{
		SubscriptionId: 1,
		Coin:           60,
		Address:        "A",
	}, {
		SubscriptionId: 1,
		Coin:           60,
		Address:        "B",
	}}

	newSubscription = []models.SubscriptionData{{
		SubscriptionId: 1,
		Coin:           60,
		Address:        "A",
	}, {
		SubscriptionId: 1,
		Coin:           60,
		Address:        "B",
	}}

	update, delete = getSubscriptionsToDeleteAndUpdate(oldSubscriptions, newSubscription)
	assert.Len(t, update, 0)
	assert.Len(t, delete, 0)

}
