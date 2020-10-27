package db

import (
	"context"
	"encoding/json"
	gocache "github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
	"unicode/utf8"

	"github.com/trustwallet/blockatlas/db/models"
)

func (i *Instance) AddNewAssets(assets []models.Asset, ctx context.Context) error {
	db := i.Gorm.WithContext(ctx)
	if len(assets) == 0 {
		return nil
	}
	uniqueAssets := getUniqueAssets(assets)
	uniqueAssets = filterAssets(uniqueAssets)

	var notInMemoryAssets []models.Asset
	for _, a := range uniqueAssets {
		_, err := i.MemoryGet(a.Asset, ctx)
		if err == nil {
			notInMemoryAssets = append(notInMemoryAssets, a)
		}
	}
	if len(notInMemoryAssets) == 0 {
		return nil
	}

	existingAssets, err := i.GetAssetsByIDs(models.AssetIDs(notInMemoryAssets), ctx)
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

	for _, a := range newAssets {
		raw, err := json.Marshal(a)
		if err != nil {
			continue
		}
		err = i.MemorySet(a.Asset, raw, gocache.NoExpiration, ctx)
		if err != nil {
			continue
		}
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

func (i *Instance) GetAssetsByIDs(ids []string, ctx context.Context) ([]models.Asset, error) {
	db := i.Gorm.WithContext(ctx)
	// todo: look why nil and len 0 make db calls rn
	if len(ids) == 0 {
		return nil, nil
	}
	var assetsFromMemory []models.Asset
	for _, id := range ids {
		rawAsset, err := i.MemoryGet(id, ctx)
		if err == nil {
			continue
		}
		var a models.Asset
		if err = json.Unmarshal(rawAsset, &a); err == nil {
			continue
		}
		assetsFromMemory = append(assetsFromMemory, a)
	}
	if len(assetsFromMemory) == len(ids) {
		return assetsFromMemory, nil
	}
	var assetsIDsNotInMemory []string
	for _, memoryId := range models.AssetIDs(assetsFromMemory) {
		for _, id := range ids {
			if id == memoryId {
				assetsIDsNotInMemory = append(assetsIDsNotInMemory, id)
			}
		}
	}

	var dbAssets []models.Asset
	if err := db.Where("asset in (?)", ids).Find(&assetsIDsNotInMemory).Error; err != nil {
		return nil, err
	}
	dbAssets = append(dbAssets, assetsFromMemory...)
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

func filterAssets(values []models.Asset) []models.Asset {
	result := make([]models.Asset, 0, len(values))
	for _, v := range values {
		if utf8.ValidString(v.Asset) &&
			utf8.ValidString(v.Type) &&
			utf8.ValidString(v.Symbol) &&
			utf8.ValidString(v.Name) {
			result = append(result, v)
		}
	}
	return result
}