package db

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"

	"github.com/trustwallet/blockatlas/db/models"
)

func (i *Instance) AddNewAssets(assets []models.Asset, ctx context.Context) error {
	db := i.Gorm.WithContext(ctx)
	if len(assets) == 0 {
		return nil
	}
	uniqueAssets := getUniqueAssets(assets)
	existingAssets, err := i.GetAssetsByIDs(models.AssetIDs(uniqueAssets), ctx)
	if err != nil {
		return err
	}
	if len(existingAssets) == 0 {
		return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&uniqueAssets).Error
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
	for k, v := range allAssetsMap {
		_, ok := existingAssetsMap[k]
		if !ok && v.Asset != "" {
			newAssets = append(newAssets, v)
		}
	}
	if len(newAssets) == 0 {
		return nil
	}

	assetsBatch := assetsBatch(newAssets, batchCount)

	return db.Transaction(func(tx *gorm.DB) error {
		for _, na := range assetsBatch {
			err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&na).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func assetsBatch(values []models.Asset, sizeUint uint) [][]models.Asset {
	size := int(sizeUint)
	resultLength := (len(values) + size - 1) / size
	result := make([][]models.Asset, resultLength)
	lo, hi := 0, size
	for i := range result {
		if hi > len(values) {
			hi = len(values)
		}
		result[i] = values[lo:hi:hi]
		lo, hi = hi, hi+size
	}
	return result
}

func (i *Instance) GetAssetsByIDs(ids []string, ctx context.Context) ([]models.Asset, error) {
	db := i.Gorm.WithContext(ctx)
	if len(ids) == 0 {
		return nil, nil
	}
	var dbAssets []models.Asset
	if err := db.Where("asset in (?)", ids).Find(&dbAssets).Error; err != nil {
		return nil, err
	}
	return dbAssets, nil
}

func (i *Instance) GetAssetsFrom(from time.Time, ctx context.Context) ([]models.Asset, error) {
	db := i.Gorm.WithContext(ctx)
	var dbAssets []models.Asset
	if err := db.Find(&dbAssets, "created_at > ?", from).Error; err != nil {
		return nil, err
	}
	return dbAssets, nil
}
