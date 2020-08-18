// +build integration

package db_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"sort"
	"testing"
)

func Test_AddNewAssociationForAddress(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)
	assets := []string{"aa", "bbb", "cccc"}

	err := database.AddAssociationsForAddress("a", assets, context.Background())
	assert.Nil(t, err)

	associations, err := database.GetAssociationsByAddresses([]string{"a"}, context.Background())
	assert.Nil(t, err)

	var assetIDsFromDB []string
	for _, a := range associations {
		assetIDsFromDB = append(assetIDsFromDB, a.Asset.AssetID)
	}

	sort.Slice(assets, func(i, j int) bool {
		return len(assets[i]) > len(assets[j])
	})

	sort.Slice(assetIDsFromDB, func(i, j int) bool {
		return len(assetIDsFromDB[i]) > len(assetIDsFromDB[j])
	})

	assert.Equal(t, assetIDsFromDB, assets)

	//var subscriptions []models.Subscription
	//
	//for i := 0; i < 100; i++ {
	//	subscriptions = append(subscriptions, models.Subscription{
	//		Coin:    uint(i),
	//		Address: "testAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddr",
	//	})
	//}
	//
	//assert.Nil(t, database.AddSubscriptions(subscriptions, context.Background()))
	//for i := 0; i < 100; i++ {
	//	s, err := database.GetSubscriptions(uint(i), []string{"testAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddrtestAddr"}, context.Background())
	//	assert.Nil(t, err)
	//	assert.NotNil(t, s)
	//}

}
