package db

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"go.elastic.co/apm/module/apmgorm"
)

func (i *Instance) GetSubscriptionsForNotifications(addresses []string, ctx context.Context) ([]models.NotificationSubscription, error) {
	if len(addresses) == 0 {
		return nil, errors.E("Empty addresses")
	}
	db := apmgorm.WithContext(ctx, i.Gorm)

	addressesSubQuery := db.
		Table("addresses").
		Select("id").
		Where("address in (?)", addresses).
		QueryExpr()

	var subscriptionsDataList []models.NotificationSubscription
	err := db.
		Preload("Address").
		Where("address_id in (?)", addressesSubQuery).
		Find(&subscriptionsDataList).
		Error
	if err != nil {
		return nil, err
	}
	return subscriptionsDataList, nil
}

func (i *Instance) AddSubscriptionsForNotifications(addresses []string, ctx context.Context) error {
	if len(addresses) == 0 {
		return errors.E("Empty subscriptions")
	}
	db := apmgorm.WithContext(ctx, i.Gorm)

	return db.Transaction(func(tx *gorm.DB) error {
		uniqueAddresses := getUniqueStrings(addresses)
		uniqueAddressesModel := make([]models.Address, 0, len(uniqueAddresses))
		for _, a := range uniqueAddresses {
			uniqueAddressesModel = append(uniqueAddressesModel, models.Address{
				Address: a,
			})
		}

		err := BulkInsert(db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING"), uniqueAddressesModel)
		if err != nil {
			return err
		}

		var addressesFromDB []models.Address
		err = db.Where("address in (?)", uniqueAddresses).Find(&addressesFromDB).Error
		if err != nil {
			return err
		}

		result := make([]models.NotificationSubscription, 0, len(addressesFromDB))
		for _, a := range addressesFromDB {
			result = append(result, models.NotificationSubscription{
				AddressID: a.ID,
			})
		}
		return BulkInsert(db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING"), result)
	})
}

func (i *Instance) DeleteSubscriptionsForNotifications(addresses []string, ctx context.Context) error {
	if len(addresses) == 0 {
		return errors.E("Empty subscriptions")
	}
	db := apmgorm.WithContext(ctx, i.Gorm)

	//var notificationsSubscriptions []models.NotificationSubscription
	//for _, a := range addresses {
	//	ma := models.Address{Address: a}
	//	notificationsSubscriptions = append(notificationsSubscriptions, models.NotificationSubscription{Address: ma})
	//}

	addressSubQuery := db.Table("addresses").
		Select("id").
		Where("address in (?)", addresses).
		QueryExpr()

	return db.
		Where("address_id in (?)", addressSubQuery).
		Delete(&models.NotificationSubscription{}).
		Error
}
