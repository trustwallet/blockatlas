package db

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/tests/integration/setup/testdata"
	"testing"
)

func TestGetSubscriptionsToDeleteAndUpdate(t *testing.T) {
	oldSubscriptions := []models.SubscriptionData{{
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "A",
	}, {
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "B",
	}, {
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "C",
	}}

	newSubscription := []models.SubscriptionData{{
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "B",
	}, {
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "C",
	}}

	update, delete := getSubscriptionsToDeleteAndUpdate(oldSubscriptions, newSubscription)
	assert.Len(t, update, 0)
	assert.Len(t, delete, 1)
	assert.Equal(t, "A", delete[0].Address)

	oldSubscriptions = []models.SubscriptionData{{
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "A",
	}, {
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "B",
	}, {
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "C",
	}, {
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "D",
	}}

	newSubscription = []models.SubscriptionData{{
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "E",
	}}

	update, delete = getSubscriptionsToDeleteAndUpdate(oldSubscriptions, newSubscription)
	assert.Len(t, update, 1)
	assert.Len(t, delete, 4)

	oldSubscriptions = []models.SubscriptionData{{
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "A",
	}, {
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "B",
	}}

	newSubscription = []models.SubscriptionData{{
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "A",
	}, {
		SubscriptionId: 1,
		Coin:           &testdata.EthCoin.ID,
		Address:        "B",
	}}

	update, delete = getSubscriptionsToDeleteAndUpdate(oldSubscriptions, newSubscription)
	assert.Len(t, update, 0)
	assert.Len(t, delete, 0)

}
