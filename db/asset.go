package db

import (
	"encoding/json"
	"time"

	"gorm.io/gorm/clause"

	gocache "github.com/patrickmn/go-cache"

	"github.com/trustwallet/blockatlas/db/models"
)

func (i *Instance) GetAsset(assetId string) (models.Asset, error) {
	var asset models.Asset
	err := i.Gorm.First(&asset, "asset = ?", assetId).Error
	if err != nil {
		return asset, err
	}
	return asset, nil
}

func (i *Instance) GetAssetsByIDs(ids []string) ([]models.Asset, error) {
	//TODO: look why nil and len 0 make db calls rn
	if len(ids) == 0 {
		return nil, nil
	}

	var dbAssets []models.Asset
	if err := i.Gorm.
		Where("asset in (?)", ids).
		Find(&dbAssets).Error; err != nil {
		return nil, err
	}
	return dbAssets, nil
}

func (i *Instance) GetSubscriptionsByAddressIDs(ids []string, from time.Time) ([]models.SubscriptionsAssetAssociation, error) {
	var associations []models.SubscriptionsAssetAssociation
	if err := i.Gorm.
		Joins("join subscriptions on subscriptions.id = subscriptions_asset_associations.subscription_id", ids).
		Preload("Subscription").
		Preload("Asset").
		Where("subscriptions.address in (?)", ids).
		Where("subscriptions_asset_associations.updated_at > ?", from).
		Find(&associations).Error; err != nil {
		return nil, err
	}
	return associations, nil
}

func (i *Instance) AddNewAssets(assets []models.Asset) error {
	if len(assets) == 0 {
		return nil
	}

	uniqueAssets := getUniqueAssets(assets)

	var notInMemoryAssets []models.Asset
	for _, a := range uniqueAssets {
		_, err := i.MemoryGet(a.Asset)
		if err != nil {
			notInMemoryAssets = append(notInMemoryAssets, a)
		}
	}
	if len(notInMemoryAssets) == 0 {
		return nil
	}

	existingAssets, err := i.GetAssetsByIDs(models.AssetIDs(notInMemoryAssets))
	if err != nil {
		return err
	}
	if len(existingAssets) == 0 {
		i.addToMemory(notInMemoryAssets)
		return i.Gorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&notInMemoryAssets).Error
	}
	allAssetsMap := make(map[string]models.Asset)
	for _, ua := range notInMemoryAssets {
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
	i.addToMemory(newAssets)

	return i.Gorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&newAssets).Error
}

func (i *Instance) addToMemory(newAssets []models.Asset) {
	for _, a := range newAssets {
		raw, err := json.Marshal(a)
		if err != nil {
			continue
		}
		err = i.MemorySet(a.Asset, raw, gocache.NoExpiration)
		if err != nil {
			continue
		}
	}
}

func (i *Instance) GetAssetsFrom(from time.Time) ([]models.Asset, error) {
	var dbAssets []models.Asset
	if err := i.Gorm.
		Where("created_at > ?", from).
		Order("created_at desc").
		Limit(1000).
		Find(&dbAssets).Error; err != nil {
		return nil, err
	}
	return dbAssets, nil
}

func getUniqueAssets(values []models.Asset) []models.Asset {
	keys := make(map[string]bool)
	var list []models.Asset
	for _, entry := range values {
		if _, value := keys[entry.Asset]; !value {
			keys[entry.Asset] = true
			list = append(list, entry)
		}
	}
	return list
}
