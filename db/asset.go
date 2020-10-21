package db

import (
	"context"
	"gorm.io/gorm/clause"
	"time"

	"github.com/trustwallet/blockatlas/db/models"
)

func (i *Instance) AddNewAssets(assets []models.Asset, ctx context.Context) error {
	if len(assets) == 0 {
		return nil
	}
	uniqueAssets := getUniqueAssets(assets)
	existingAssets, err := i.GetAssetsByIDs(models.AssetIDs(uniqueAssets), ctx)
	if err != nil {
		return err
	}
	allAssetsMap := make(map[string]models.Asset)
	for _, ua := range uniqueAssets {
		allAssetsMap[ua.Asset] = ua
	}
	existingAssetsMap := make(map[string]models.Asset)
	for _, ea := range existingAssets {
		existingAssetsMap[ea.Asset] = ea
	}
	var newAssets []models.Asset
	for k := range allAssetsMap {
		a, ok := existingAssetsMap[k]
		if !ok {
			newAssets = append(newAssets, a)
		}
	}
	if len(newAssets) == 0 {
		return nil
	}

	return i.Gorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&newAssets).Error
}

func (i *Instance) GetAssetsByIDs(ids []string, ctx context.Context) ([]models.Asset, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var dbAssets []models.Asset
	if err := i.Gorm.Where("asset in (?)", ids).Find(&dbAssets).Error; err != nil {
		return nil, err
	}
	return dbAssets, nil
}

func (i *Instance) GetAssetsByIDsFrom(ids []string, from time.Time, ctx context.Context) ([]models.Asset, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var dbAssets []models.Asset
	if err := i.Gorm.Where("asset in (?)", ids).Find(&dbAssets, "created_at > ?", from).Error; err != nil {
		return nil, err
	}
	return dbAssets, nil
}
