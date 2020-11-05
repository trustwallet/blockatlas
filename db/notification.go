package db

import (
	"context"
	"errors"
	"github.com/trustwallet/blockatlas/db/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (i *Instance) GetSubscriptionsForNotifications(addresses []string, ctx context.Context) ([]models.NotificationSubscription, error) {
	if len(addresses) == 0 {
		return nil, errors.New("Empty addresses")
	}

	db := i.Gorm.WithContext(ctx)
	var subscriptionsDataList []models.NotificationSubscription
	err := db.Joins("Address").Limit(len(addresses)).Find(&subscriptionsDataList, "address in (?)", addresses).Error
	if err != nil {
		return nil, err
	}
	return subscriptionsDataList, nil
}

func (i *Instance) AddSubscriptionsForNotifications(addresses []string, ctx context.Context) error {
	if len(addresses) == 0 {
		return errors.New("Empty subscriptions")
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

		err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&uniqueAddressesModel).Error
		if err != nil {
			return err
		}

		var dbAddresses []models.Address
		if err = tx.Find(&dbAddresses, "address in (?)", uniqueAddresses).Error; err != nil {
			return err
		}

		result := make([]models.NotificationSubscription, 0, len(dbAddresses))
		for _, a := range dbAddresses {
			result = append(result, models.NotificationSubscription{AddressID: a.ID})
		}
		return tx.Clauses(clause.OnConflict{
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
		return errors.New("Empty subscriptions")
	}
	q := `DELETE FROM notification_subscriptions ns USING addresses a where ns.address_id = a.id AND a.address IN (?);`
	return i.Gorm.WithContext(ctx).Exec(q, addresses).Error
}
