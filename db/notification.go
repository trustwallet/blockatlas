package db

import (
	"context"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (i *Instance) GetSubscriptionsForNotifications(addresses []string, ctx context.Context) ([]models.NotificationSubscription, error) {
	if len(addresses) == 0 {
		return nil, errors.E("Empty addresses")
	}
	db := i.Gorm.WithContext(ctx)

	var subscriptionsDataList []models.NotificationSubscription
	err := db.Joins("Address").Find(&subscriptionsDataList, "address in (?)", addresses).Distinct().Error
	if err != nil {
		return nil, err
	}
	return subscriptionsDataList, nil
}

func (i *Instance) AddSubscriptionsForNotifications(addresses []string, ctx context.Context) error {
	if len(addresses) == 0 {
		return errors.E("Empty subscriptions")
	}
	db := i.Gorm.WithContext(ctx)
	return db.Transaction(func(tx *gorm.DB) error {
		uniqueAddresses := getUniqueStrings(addresses)
		uniqueAddressesModel := make([]models.Address, 0, len(uniqueAddresses))
		for _, a := range uniqueAddresses {
			uniqueAddressesModel = append(uniqueAddressesModel, models.Address{
				Address: a,
			})
		}

		err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&uniqueAddressesModel).Error
		if err != nil {
			return err
		}

		var dbAddresses []models.Address
		err = db.Where("address in (?)", uniqueAddresses).Find(&dbAddresses).Distinct().Error
		if err != nil {
			return err
		}

		result := make([]models.NotificationSubscription, 0, len(dbAddresses))
		for _, a := range dbAddresses {
			result = append(result, models.NotificationSubscription{AddressID: a.ID})
		}
		return db.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{
					Name: "address_id",
				},
			},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"deleted_at": nil,
			}),
		}).Create(&result).Error
	})
}

func (i *Instance) DeleteSubscriptionsForNotifications(addresses []string, ctx context.Context) error {
	if len(addresses) == 0 {
		return errors.E("Empty subscriptions")
	}
	q := `DELETE FROM notification_subscriptions ns USING addresses a where ns.address_id = a.id AND a.address IN (?);`
	return i.Gorm.WithContext(ctx).Exec(q, addresses).Error
}
