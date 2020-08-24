package db

import (
	"context"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"go.elastic.co/apm/module/apmgorm"
)

func (i *Instance) GetSubscriptionsForNotification(addresses []string, ctx context.Context) ([]models.NotificationSubscription, error) {
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
		Find(&subscriptionsDataList).Error

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (i *Instance) AddSubscriptions(addresses []string, ctx context.Context) error {
	if len(addresses) == 0 {
		return errors.E("Empty subscriptions")
	}

	db := apmgorm.WithContext(ctx, i.Gorm)

	var notificationsSubscriptions []models.NotificationSubscription
	for _, a := range addresses {
		na := models.Address{Address: a}
		notificationsSubscriptions = append(notificationsSubscriptions, models.NotificationSubscription{Address: na})
	}
	return BulkInsert(db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING"), notificationsSubscriptions)
}

func (i *Instance) DeleteSubscriptions(addresses []string, ctx context.Context) error {
	if len(addresses) == 0 {
		return errors.E("Empty subscriptions")
	}

	db := apmgorm.WithContext(ctx, i.Gorm)

	var notificationsSubscriptions []models.NotificationSubscription
	for _, a := range addresses {
		na := models.Address{Address: a}
		notificationsSubscriptions = append(notificationsSubscriptions, models.NotificationSubscription{Address: na})
	}

	return db.
		Where("address in (?)", addresses).
		Delete(&models.NotificationSubscription{}).
		Error
}
