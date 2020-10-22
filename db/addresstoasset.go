package db

import (
	"context"
	"github.com/trustwallet/blockatlas/db/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func (i Instance) GetSubscribedAddressesForAssets(ctx context.Context, addresses []string) ([]models.Address, error) {
	db := i.Gorm.WithContext(ctx)
	var result []models.Address
	err := db.Model(&models.AssetSubscription{}).
		Select("id", "address").
		Joins("LEFT JOIN addresses a ON a.id = address_id").
		Where("address in (?)", addresses).
		Scan(&result).
		Limit(len(addresses)).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i Instance) GetAssetsMapByAddresses(addresses []string, ctx context.Context) (map[string][]models.Asset, error) {
	db := i.Gorm.WithContext(ctx)
	var associations []models.AddressToAssetAssociation
	if err := db.Joins("Address").Joins("Asset").Find(&associations, "address in (?)", addresses).Error; err != nil {
		return nil, err
	}

	result := make(map[string][]models.Asset)
	for _, a := range associations {
		assets := result[a.Address.Address]
		result[a.Address.Address] = append(assets, a.Asset)
	}
	return result, nil
}

func (i Instance) GetAssetsMapByAddressesFromTime(addresses []string, from time.Time, ctx context.Context) (map[string][]models.Asset, error) {
	if len(addresses) == 0 {
		return map[string][]models.Asset{}, nil
	}
	db := i.Gorm.WithContext(ctx)
	var associations []models.AddressToAssetAssociation
	err := db.Joins("Address").Where("address in (?)", addresses).Joins("Asset").Find(&associations, "address_to_asset_associations.created_at > ?", from).Error
	if err != nil {
		return nil, err
	}

	result := make(map[string][]models.Asset)
	for _, a := range associations {
		assets := result[a.Address.Address]
		result[a.Address.Address] = append(assets, a.Asset)
	}
	return result, nil
}

func (i *Instance) GetAssociationsByAddresses(addresses []string, ctx context.Context) ([]models.AddressToAssetAssociation, error) {
	db := i.Gorm.WithContext(ctx)
	var result []models.AddressToAssetAssociation
	if err := db.Joins("Address").Joins("Asset").Find(&result, "address in (?)", addresses).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i *Instance) GetAssociationsByAddressesFromTime(addresses []string, from time.Time, ctx context.Context) ([]models.AddressToAssetAssociation, error) {
	db := i.Gorm.WithContext(ctx)
	var result []models.AddressToAssetAssociation
	err := db.Joins("Address").Where("address in (?)", addresses).Joins("Asset").Find(&result, "created_at > ?", from).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i *Instance) AddAssociationsForAddress(address string, assets []models.Asset, ctx context.Context) error {
	db := i.Gorm.WithContext(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		uniqueAssets := getUniqueAssets(assets)

		var err error
		dbAddress := models.Address{Address: address}
		err = tx.Clauses(clause.OnConflict{DoNothing: true}).FirstOrCreate(&dbAddress, "address = ?", address).Error
		if err != nil {
			return err
		}

		if len(uniqueAssets) > 0 {
			if err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&uniqueAssets).Error; err != nil {
				return err
			}
		}

		var dbAssets []models.Asset
		if len(uniqueAssets) > 0 {
			err = tx.
				Where("asset in (?)", models.AssetIDs(uniqueAssets)).
				Find(&dbAssets).Error
			if err != nil {
				return err
			}
		}

		assetsSub := models.AssetSubscription{AddressID: dbAddress.ID}
		err = tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{
					Name: "address_id",
				},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"deleted_at": nil,
			}),
		}).Create(&assetsSub).Error
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
		if len(result) > 0 {
			return tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&result).Error
		}
		return nil
	})
}

func (i *Instance) UpdateAssociationsForExistingAddresses(associations map[string][]models.Asset, ctx context.Context) error {
	db := i.Gorm.WithContext(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		assets := make([]models.Asset, 0, len(associations))
		for _, v := range associations {
			assets = append(assets, v...)
		}

		if len(assets) == 0 {
			return nil
		}

		uniqueAssets := getUniqueAssets(assets)
		uniqueAssetsModel := make([]models.Asset, 0, len(uniqueAssets))
		for _, l := range uniqueAssets {
			uniqueAssetsModel = append(uniqueAssetsModel, models.Asset{Asset: l.Asset})
		}
		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&uniqueAssetsModel).Error; err != nil {
			return err
		}

		var dbAssets []models.Asset
		err := tx.
			Where("asset in (?)", models.AssetIDs(uniqueAssets)).
			Find(&dbAssets).
			Limit(len(uniqueAssets)).
			Error
		if err != nil {
			return err
		}

		assetsMap := makeMapAssets(dbAssets)

		addresses := make([]string, 0, len(associations))
		for k := range associations {
			addresses = append(addresses, k)
		}

		var dbAddresses []models.Address
		if err := tx.Find(&dbAddresses, "address in (?)", addresses).Limit(len(addresses)).Error; err != nil {
			return err
		}

		var addressSubs []models.AssetSubscription
		for _, a := range dbAddresses {
			sub := models.AssetSubscription{AddressID: a.ID}
			addressSubs = append(addressSubs, sub)
		}

		err = tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{
					Name: "address_id",
				},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"deleted_at": nil,
			}),
		}).Create(&addressSubs).Error
		if err != nil {
			return err
		}

		addressesMap := makeMapAddress(dbAddresses)

		var result []models.AddressToAssetAssociation
		for address, assets := range associations {
			for _, asset := range assets {
				addressID, ok := addressesMap[address]
				if !ok || addressID == 0 {
					continue
				}
				assetID, ok := assetsMap[asset.Asset]
				if !ok || assetID == 0 {
					continue
				}
				r := models.AddressToAssetAssociation{
					AddressID: addressID,
					AssetID:   assetID,
				}
				result = append(result, r)
			}
		}
		return tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&result).Error
	})
}

func makeMapAssets(addresses []models.Asset) map[string]uint {
	result := make(map[string]uint)
	for _, a := range addresses {
		result[a.Asset] = a.ID
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
	keys := make(map[string]struct{})
	var list []string
	for _, entry := range values {
		if _, value := keys[entry]; !value {
			keys[entry] = struct{}{}
			list = append(list, entry)
		}
	}
	return list
}

func getUniqueAssets(values []models.Asset) []models.Asset {
	keys := make(map[string]struct{})
	var list []models.Asset
	for _, entry := range values {
		if _, value := keys[entry.Asset]; !value {
			keys[entry.Asset] = struct{}{}
			list = append(list, entry)
		}
	}
	return list
}
