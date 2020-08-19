// +build integration

package db_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"sort"
	"testing"
)

func Test_GetAssetsByAddressesMap(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)

	assets := []string{"aa", "bbb", "cccc"}

	err := database.AddAssociationsForAddress("a", assets, context.Background())
	assert.Nil(t, err)

	err = database.AddAssociationsForAddress("b", nil, context.Background())
	assert.Nil(t, err)

	m, err := database.GetTokensByAddressesMap([]string{"a", "b"}, context.Background())
	assert.Nil(t, err)
	wantedMap := make(map[string][]string)
	wantedMap["a"] = assets
	wantedMap["b"] = []string{""}
	assert.Equal(t, wantedMap, m)
}

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

	err = database.AddAssociationsForAddress("b", nil, context.Background())
	assert.Nil(t, err)

	associations2, err := database.GetAssociationsByAddresses([]string{"b"}, context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, associations2)
}

func Test_UpdateAssociationsForExistingAddresses(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)
	assets := []string{"f"}

	err := database.AddAssociationsForAddress("A", assets, context.Background())
	assert.Nil(t, err)

	err = database.AddAssociationsForAddress("B", assets, context.Background())
	assert.Nil(t, err)

	assetsForA := []string{"aa", "bbb", "cccc"}
	assetsForB := []string{"as", "bbb", "cccc"}

	updateMap := make(map[string][]string)
	updateMap["A"] = assetsForA
	updateMap["B"] = assetsForB

	err = database.UpdateAssociationsForExistingAddresses(updateMap, context.Background())
	assert.Nil(t, err)

	associationsA, err := database.GetAssociationsByAddresses([]string{"A"}, context.Background())
	assert.Nil(t, err)

	var assetIDsFromDBA []string
	for _, a := range associationsA {
		assetIDsFromDBA = append(assetIDsFromDBA, a.Asset.AssetID)
	}
	assetsA := []string{"aa", "bbb", "cccc", "f"}

	sort.Slice(assetsA, func(i, j int) bool {
		return len(assetsA[i]) > len(assetsA[j])
	})

	sort.Slice(assetIDsFromDBA, func(i, j int) bool {
		return len(assetIDsFromDBA[i]) > len(assetIDsFromDBA[j])
	})

	assert.Equal(t, assetIDsFromDBA, assetsA)

	associationsB, err := database.GetAssociationsByAddresses([]string{"B"}, context.Background())
	assert.Nil(t, err)

	var assetIDsFromDBB []string
	for _, a := range associationsB {
		assetIDsFromDBB = append(assetIDsFromDBB, a.Asset.AssetID)
	}
	assetsB := []string{"as", "bbb", "cccc", "f"}

	sort.Slice(assetsB, func(i, j int) bool {
		return len(assetsB[i]) > len(assetsB[j])
	})

	sort.Slice(assetIDsFromDBB, func(i, j int) bool {
		return len(assetIDsFromDBB[i]) > len(assetIDsFromDBB[j])
	})

	assert.Equal(t, assetIDsFromDBB, assetsB)

	associationsAB, err := database.GetAssociationsByAddresses([]string{"A", "B"}, context.Background())
	assert.Nil(t, err)

	var assetIDsFromDBAB []string
	for _, a := range associationsAB {
		assetIDsFromDBAB = append(assetIDsFromDBAB, a.Asset.AssetID)
	}
	assetsAB := []string{"cccc", "cccc", "bbb", "bbb", "aa", "as", "f", "f"}

	sort.Slice(assetsAB, func(i, j int) bool {
		return len(assetsAB[i]) > len(assetsAB[j])
	})

	sort.Slice(assetIDsFromDBAB, func(i, j int) bool {
		return len(assetIDsFromDBAB[i]) > len(assetIDsFromDBAB[j])
	})

	assert.Equal(t, assetIDsFromDBAB, assetsAB)
}
