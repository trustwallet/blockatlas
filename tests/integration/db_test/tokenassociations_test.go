// +build integration

package db_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"sort"
	"testing"
	"time"
)

func Test_GetAssetsMapByAddresses(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)

	assets := []models.Asset{{Asset: "aa"}, {Asset: "bbb"}, {Asset: "cccc"}}

	err := database.AddAssociationsForAddress("a", assets, context.Background())
	assert.Nil(t, err)

	err = database.AddAssociationsForAddress("b", nil, context.Background())
	assert.Nil(t, err)

	m, err := database.GetAssetsMapByAddresses([]string{"a", "b"}, context.Background())
	assert.Nil(t, err)
	wantedMap := make(map[string][]models.Asset)
	wantedMap["a"] = assets
	for i, a := range m {
		for ii, aa := range a {
			assert.Equal(t, wantedMap[i][ii].Asset, aa.Asset)
		}
	}

}

func Test_GetAssetsMapByAddressesFromTime(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)

	assets := []models.Asset{{Asset: "aa"}, {Asset: "bbb"}, {Asset: "cccc"}}

	err := database.AddAssociationsForAddress("a", assets, context.Background())
	assert.Nil(t, err)

	err = database.AddAssociationsForAddress("b", nil, context.Background())
	assert.Nil(t, err)
	tm := time.Now().Unix() - 100
	m, err := database.GetAssetsMapByAddressesFromTime([]string{"a", "b"}, time.Unix(tm, 0), context.Background())
	assert.Nil(t, err)
	wantedMap := make(map[string][]models.Asset)
	wantedMap["a"] = assets

	for i, a := range m {
		for ii, aa := range a {
			assert.Equal(t, wantedMap[i][ii].Asset, aa.Asset)
		}
	}

	m, err = database.GetAssetsMapByAddressesFromTime([]string{"a", "b"}, time.Unix(tm+101, 0), context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 0, len(m))
}

func Test_GetSubscribedAddressesForAssets(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)

	assets := []models.Asset{{Asset: "aa"}, {Asset: "bbb"}, {Asset: "cccc"}}

	err := database.AddAssociationsForAddress("a", assets, context.Background())
	assert.Nil(t, err)

	err = database.AddAssociationsForAddress("b", nil, context.Background())
	assert.Nil(t, err)

	m, err := database.GetSubscribedAddressesForAssets(context.Background(), []string{"a", "b"})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(m))
}

func Test_AddNewAssociationForAddress(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)
	assets := []models.Asset{{Asset: "aa"}, {Asset: "bbb"}, {Asset: "cccc"}}

	err := database.AddAssociationsForAddress("a", assets, context.Background())
	assert.Nil(t, err)

	associations, err := database.GetAssociationsByAddresses([]string{"a"}, context.Background())
	assert.Nil(t, err)

	var assetIDsFromDB []models.Asset
	for _, a := range associations {
		assetIDsFromDB = append(assetIDsFromDB, a.Asset)
	}

	sort.Slice(assets, func(i, j int) bool {
		return len(assets[i].Asset) > len(assets[j].Asset)
	})

	sort.Slice(assetIDsFromDB, func(i, j int) bool {
		return len(assetIDsFromDB[i].Asset) > len(assetIDsFromDB[j].Asset)
	})

	for i, a := range assets {
		assert.Equal(t, assetIDsFromDB[i].Asset, a.Asset)
	}

	err = database.AddAssociationsForAddress("b", nil, context.Background())
	assert.Nil(t, err)

	associations2, err := database.GetAssociationsByAddresses([]string{"b"}, context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, associations2)
}

func Test_UpdateAssociationsForExistingAddresses(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)
	assets := []models.Asset{{Asset: "f"}}

	err := database.AddAssociationsForAddress("A", assets, context.Background())
	assert.Nil(t, err)

	err = database.AddAssociationsForAddress("B", assets, context.Background())
	assert.Nil(t, err)

	assetsForA := []models.Asset{{Asset: "aa"}, {Asset: "bbb"}, {Asset: "cccc"}}
	assetsForB := []models.Asset{{Asset: "bbb"}, {Asset: "cccc"}}

	updateMap := make(map[string][]models.Asset)
	updateMap["A"] = assetsForA
	updateMap["B"] = assetsForB

	err = database.UpdateAssociationsForExistingAddresses(updateMap, context.Background())
	assert.Nil(t, err)

	associationsA, err := database.GetAssociationsByAddresses([]string{"A"}, context.Background())
	assert.Nil(t, err)

	var assetIDsFromDBA []models.Asset
	for _, a := range associationsA {
		assetIDsFromDBA = append(assetIDsFromDBA, a.Asset)
	}
	assetsA := []models.Asset{{Asset: "aa"}, {Asset: "bbb"}, {Asset: "cccc"}, {Asset: "f"}}

	sort.Slice(assetsA, func(i, j int) bool {
		return len(assetsA[i].Asset) > len(assetsA[j].Asset)
	})

	sort.Slice(assetIDsFromDBA, func(i, j int) bool {
		return len(assetIDsFromDBA[i].Asset) > len(assetIDsFromDBA[j].Asset)
	})

	for i, a := range assetsA {
		assert.Equal(t, assetIDsFromDBA[i].Asset, a.Asset)
	}

	associationsB, err := database.GetAssociationsByAddresses([]string{"B"}, context.Background())
	assert.Nil(t, err)

	var assetIDsFromDBB []models.Asset
	for _, a := range associationsB {
		assetIDsFromDBB = append(assetIDsFromDBB, a.Asset)
	}
	assetsB := []models.Asset{{Asset: "bbb"}, {Asset: "cccc"}, {Asset: "f"}}

	sort.Slice(assetsB, func(i, j int) bool {
		return len(assetsB[i].Asset) > len(assetsB[j].Asset)
	})

	sort.Slice(assetIDsFromDBB, func(i, j int) bool {
		return len(assetIDsFromDBB[i].Asset) > len(assetIDsFromDBB[j].Asset)
	})

	for i, a := range assetsB {
		assert.Equal(t, assetIDsFromDBB[i].Asset, a.Asset)
	}

	associationsAB, err := database.GetAssociationsByAddresses([]string{"A", "B"}, context.Background())
	assert.Nil(t, err)

	var assetIDsFromDBAB []models.Asset
	for _, a := range associationsAB {
		assetIDsFromDBAB = append(assetIDsFromDBAB, a.Asset)
	}
	assetsAB := []models.Asset{{Asset: "cccc"}, {Asset: "cccc"}, {Asset: "bbb"}, {Asset: "bbb"}, {Asset: "aa"}, {Asset: "f"}, {Asset: "f"}}

	sort.Slice(assetsAB, func(i, j int) bool {
		return len(assetsAB[i].Asset) > len(assetsAB[j].Asset)
	})

	sort.Slice(assetIDsFromDBAB, func(i, j int) bool {
		return len(assetIDsFromDBAB[i].Asset) > len(assetIDsFromDBAB[j].Asset)
	})

	for i, a := range assetsAB {
		assert.Equal(t, assetIDsFromDBAB[i].Asset, a.Asset)
	}
}
