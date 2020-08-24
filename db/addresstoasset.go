package db

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/trustwallet/blockatlas/db/models"
	"go.elastic.co/apm/module/apmgorm"
)

func (i Instance) GetSubscribedAddresses(ctx context.Context, addresses []string) ([]models.Address, error) {
	db := apmgorm.WithContext(ctx, i.Gorm)

	addressesSubQuery := db.
		Table("addresses").
		Select("id").
		Where("address in (?)", addresses).
		QueryExpr()

	var assetSubs []models.AssetSubscription
	err := db.
		Preload("Address").
		Where("address_id in (?)", addressesSubQuery).
		Find(&assetSubs).
		Error
	if err != nil {
		return nil, err
	}

	var result []models.Address
	for _, a := range assetSubs {
		result = append(result, a.Address)
	}
	return result, nil
}

func (i Instance) GetAssetsMapByAddresses(addresses []string, ctx context.Context) (map[string][]string, error) {
	db := apmgorm.WithContext(ctx, i.Gorm)

	var addressesFromDB []models.Address
	err := db.Where("address in (?)", addresses).Find(&addressesFromDB).Error
	if err != nil {
		return nil, err
	}

	addressesIDs := make([]uint, 0, len(addressesFromDB))
	for _, a := range addressesFromDB {
		addressesIDs = append(addressesIDs, a.ID)
	}

	var associations []models.AddressToAssetAssociation
	err = db.
		Preload("Address").
		Preload("Asset").
		Where("address_id in (?)", addressesIDs).
		Find(&associations).
		Error
	if err != nil {
		return nil, err
	}

	result := make(map[string][]string)
	for _, a := range associations {
		assets := result[a.Address.Address]
		result[a.Address.Address] = append(assets, a.Asset.AssetID)
	}
	return result, nil
}

func (i *Instance) GetAssociationsByAddresses(addresses []string, ctx context.Context) ([]models.AddressToAssetAssociation, error) {
	db := apmgorm.WithContext(ctx, i.Gorm)

	addressesSubQuery := db.Table("addresses").
		Select("id").
		Where("address in (?)", addresses).
		QueryExpr()

	var result []models.AddressToAssetAssociation
	err := db.
		Preload("Address").
		Preload("Asset").
		Where("address_id in (?)", addressesSubQuery).
		Find(&result).
		Error
	return result, err
}

func (i *Instance) AddAssociationsForAddress(address string, assets []string, ctx context.Context) error {
	db := apmgorm.WithContext(ctx, i.Gorm)
	return db.Transaction(func(tx *gorm.DB) error {
		uniqueAssets := getUniqueStrings(assets)
		uniqueAssetsModel := make([]models.Asset, 0, len(uniqueAssets))
		for _, l := range uniqueAssets {
			uniqueAssetsModel = append(uniqueAssetsModel, models.Asset{
				AssetID: l,
			})
		}

		err := BulkInsert(db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING"), uniqueAssetsModel)
		if err != nil {
			return err
		}

		var dbAssets []models.Asset
		err = db.Where("asset_id in (?)", uniqueAssets).Find(&dbAssets).Error
		if err != nil {
			return err
		}

		dbAddress := models.Address{Address: address}
		err = db.Where("address = ?", address).FirstOrCreate(&dbAddress).Error
		if err != nil {
			return err
		}

		result := make([]models.AddressToAssetAssociation, 0, len(dbAssets))
		for _, asset := range dbAssets {
			result = append(result, models.AddressToAssetAssociation{
				AddressID: dbAddress.ID,
				AssetID:   asset.ID,
			})
		}
		return BulkInsert(db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING"), result)
	})
}

func (i *Instance) UpdateAssociationsForExistingAddresses(associations map[string][]string, ctx context.Context) error {
	db := apmgorm.WithContext(ctx, i.Gorm)
	return db.Transaction(func(tx *gorm.DB) error {
		var assets []string

		for _, v := range associations {
			assets = append(assets, v...)
		}

		uniqueAssets := getUniqueStrings(assets)
		uniqueAssetsModel := make([]models.Asset, 0, len(uniqueAssets))
		for _, l := range uniqueAssets {
			uniqueAssetsModel = append(uniqueAssetsModel, models.Asset{
				AssetID: l,
			})
		}

		err := BulkInsert(db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING"), uniqueAssetsModel)
		if err != nil {
			return err
		}

		var dbAssets []models.Asset
		err = db.Where("asset_id in (?)", uniqueAssets).Find(&dbAssets).Error
		if err != nil {
			return err
		}

		assetsMap := makeMapAssets(dbAssets)

		var addresses []string
		for k := range associations {
			addresses = append(addresses, k)
		}

		var dbAddresses []models.Address
		if err := db.Where("address in (?)", addresses).Find(&dbAddresses).Error; err != nil {
			return err
		}

		addressesMap := makeMapAddress(dbAddresses)

		var result []models.AddressToAssetAssociation

		for address, assets := range associations {
			for _, asset := range assets {
				r := models.AddressToAssetAssociation{
					AddressID: addressesMap[address],
					AssetID:   assetsMap[asset],
				}
				result = append(result, r)
			}
		}
		return BulkInsert(db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING"), result)
	})
}

func makeMapAssets(addresses []models.Asset) map[string]uint {
	result := make(map[string]uint)
	for _, a := range addresses {
		result[a.AssetID] = a.ID
	}
	return result
}

func makeMapAddress(addresses []models.Address) map[string]uint {
	result := make(map[string]uint)
	for _, a := range addresses {
		result[a.Address] = a.ID
	}
	return result
}

func getUniqueStrings(values []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range values {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
