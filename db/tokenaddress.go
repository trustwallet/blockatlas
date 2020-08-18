package db

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/trustwallet/blockatlas/db/models"
	"go.elastic.co/apm/module/apmgorm"
)

func (i *Instance) AddAssociationToAddress(association models.AddressToTokenAssociation, ctx context.Context) error {
	db := apmgorm.WithContext(ctx, i.Gorm)
	return db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING").Create(&association).Error
}

func (i *Instance) GetAssociationsByAddresses(addresses []string, ctx context.Context) ([]models.AddressToTokenAssociation, error) {
	db := apmgorm.WithContext(ctx, i.Gorm)

	addressesSubQuery := db.Table("addresses").
		Select("id").
		Where("address in (?)", addresses).
		QueryExpr()

	var result []models.AddressToTokenAssociation
	err := db.
		Preload("Address").
		Preload("Asset").
		Where("address_id in (?)", addressesSubQuery).
		Find(&result).Error

	return result, err
}

func (i *Instance) AddAssociationsForAddress(address string, assets []string, ctx context.Context) error {
	db := apmgorm.WithContext(ctx, i.Gorm)
	return db.Transaction(func(tx *gorm.DB) error {
		uniqueAssets := getUniqueAssets(assets)

		err := BulkInsert(db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING"), uniqueAssets)
		if err != nil {
			return err
		}

		var dbAssets []models.Asset
		if err := db.Find(&dbAssets).Error; err != nil {
			return err
		}

		var dbAddress models.Address

		if err := db.
			Where("address = ?", address).
			FirstOrCreate(&dbAddress).Error; err != nil {
			return err
		}

		var result []models.AddressToTokenAssociation
		for _, asset := range dbAssets {
			result = append(result, models.AddressToTokenAssociation{
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

		uniqueAssets := getUniqueAssets(assets)

		err := BulkInsert(db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING"), uniqueAssets)
		if err != nil {
			return err
		}

		var dbAssets []models.Asset
		if err := db.Find(&dbAssets).Error; err != nil {
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

		var result []models.AddressToTokenAssociation

		for address, assets := range associations {
			for _, asset := range assets {
				r := models.AddressToTokenAssociation{
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

func getUniqueAssets(assets []string) []models.Asset {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range assets {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	result := make([]models.Asset, 0, len(list))
	for _, l := range list {
		result = append(result, models.Asset{
			AssetID: l,
		})
	}

	return result
}
