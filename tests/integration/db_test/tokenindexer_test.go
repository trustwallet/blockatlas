// +build integration

package db_test

import (
	"context"
	gocache "github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	assert "github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"sort"
	"testing"
	"time"
)

func Test_AddNewAssets_Simple(t *testing.T) {
	a := []models.Asset{
		{
			Asset:    "c714_a",
			Decimals: 18,
			Name:     "A",
			Symbol:   "ABC",
			Type:     "BEP20",
		},
		{
			Asset:    "c714_b",
			Decimals: 18,
			Name:     "B",
			Symbol:   "BCD",
			Type:     "BEP20",
		},
	}
	err := database.AddNewAssets(a, context.Background())
	assert.Nil(t, err)
	assets, err := database.GetAssetsByIDs([]string{"c714_b", "c714_a"}, context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, assets)
	a = append(a, models.Asset{
		Asset:    "c714_d",
		Decimals: 18,
		Name:     "D",
		Symbol:   "DTS",
		Type:     "BEP20",
	})
	err = database.AddNewAssets(a, context.Background())
	assert.Nil(t, err)
	err = database.AddNewAssets([]models.Asset{{
		Asset:    "c714_p",
		Decimals: 0,
		Name:     "D",
		Symbol:   "DTS",
		Type:     "BEP20",
	}}, context.Background())
	assert.Nil(t, err)
	assets, err = database.GetAssetsByIDs([]string{"c714_p"}, context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 0, len(assets))
}

func Test_GetAssetsFrom_Simple(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)
	database.MemoryCache = gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	a := []models.Asset{
		{
			Asset:    "c714_a",
			Coin:     714,
			Decimals: 18,
			Name:     "A",
			Symbol:   "ABC",
			Type:     "BEP20",
		},
		{
			Asset:    "c714_b",
			Decimals: 18,
			Coin:     60,
			Name:     "B",
			Symbol:   "BCD",
			Type:     "BEP20",
		},
	}
	err := database.AddNewAssets(a, context.Background())
	assert.Nil(t, err)
	assets, err := database.GetAssetsFrom(time.Unix(0, 0), -1, context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, assets)
	assets, err = database.GetAssetsFrom(time.Unix(0, 0), 60, context.Background())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(assets))
}

func Test_AddNewAssets(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)
	database.MemoryCache = gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	type testsStruct struct {
		Name         string
		Assets       []models.Asset
		AssetsIDs    []string
		WantedErr    error
		WantedAssets []models.Asset
	}
	tests := []testsStruct{
		{
			Name: "Normal case",
			Assets: []models.Asset{
				{
					Asset:    "c714_a",
					Decimals: 15,
					Name:     "A",
					Symbol:   "ABC",
					Type:     "BEP20",
				},
				{
					Asset:    "c714_b",
					Decimals: 16,
					Name:     "BB",
					Symbol:   "BCD",
					Type:     "BEP20",
				},
			},
			AssetsIDs: []string{"c714_a", "c714_b"},
			WantedErr: nil,
			WantedAssets: []models.Asset{
				{
					Asset:    "c714_a",
					Decimals: 15,
					Name:     "A",
					Symbol:   "ABC",
					Type:     "BEP20",
				},
				{
					Asset:    "c714_b",
					Decimals: 16,
					Name:     "BB",
					Symbol:   "BCD",
					Type:     "BEP20",
				},
			},
		},
		{
			Name: "Case with new tokens and old tokens",
			Assets: []models.Asset{
				{
					Asset:    "c714_c",
					Decimals: 17,
					Name:     "CCC",
					Symbol:   "FFF",
					Type:     "ERC20",
				},
				{
					Asset:    "c714_d",
					Decimals: 18,
					Name:     "DDDD",
					Symbol:   "RRR",
					Type:     "TRC20",
				},
			},
			AssetsIDs: []string{"c714_a", "c714_b", "c714_c", "c714_d"},
			WantedErr: nil,
			WantedAssets: []models.Asset{
				{
					Asset:    "c714_a",
					Decimals: 15,
					Name:     "A",
					Symbol:   "ABC",
					Type:     "BEP20",
				},
				{
					Asset:    "c714_b",
					Decimals: 16,
					Name:     "BB",
					Symbol:   "BCD",
					Type:     "BEP20",
				},
				{
					Asset:    "c714_c",
					Decimals: 17,
					Name:     "CCC",
					Symbol:   "FFF",
					Type:     "ERC20",
				},
				{
					Asset:    "c714_d",
					Decimals: 18,
					Name:     "DDDD",
					Symbol:   "RRR",
					Type:     "TRC20",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := database.AddNewAssets(tt.Assets, context.Background())
			assert.Equal(t, tt.WantedErr, err)
			assets, err := database.GetAssetsByIDs(tt.AssetsIDs, context.Background())
			assert.Nil(t, err)
			sort.Slice(tt.WantedAssets, func(i, j int) bool {
				return len(tt.WantedAssets[i].Name) > len(tt.WantedAssets[j].Name)
			})
			sort.Slice(assets, func(i, j int) bool {
				return len(assets[i].Name) > len(assets[j].Name)
			})
			log.Info(tt.WantedAssets)
			log.Info("----------------------")
			log.Info(assets)
			for i, a := range assets {
				assert.Equal(t, tt.WantedAssets[i].Asset, a.Asset)
				assert.Equal(t, tt.WantedAssets[i].Name, a.Name)
				assert.Equal(t, tt.WantedAssets[i].Symbol, a.Symbol)
				assert.Equal(t, tt.WantedAssets[i].Type, a.Type)
				assert.Equal(t, tt.WantedAssets[i].Decimals, a.Decimals)
			}
		})
	}
}
