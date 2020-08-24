package db

import (
	"context"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

func (i *Instance) GetSubscriptionsForNotification(coin uint, addresses []string, ctx context.Context) ([]models.NotificationSubscription, error) {
	if len(addresses) == 0 {
		return nil, errors.E("Empty addresses")
	}
	//g := apmgorm.WithContext(ctx, i.Gorm)
	//var subscriptionsDataList []models.NotificationSubscription
	//err := g.
	//	Model(&models.NotificationSubscription{}).
	//	Where("address in (?) AND coin = ?", addresses, coin).
	//	Find(&subscriptionsDataList).Error
	//
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}

func (i *Instance) AddSubscriptions(subscriptions []models.NotificationSubscription, ctx context.Context) error {
	if len(subscriptions) == 0 {
		return errors.E("Empty subscriptions")
	}

	//subscriptionsBatch := toSubscriptionBatch(subscriptions, batchLimit, ctx)
	//g := apmgorm.WithContext(ctx, i.Gorm)
	//
	//for _, s := range subscriptionsBatch {
	//	if err := bulkCreate(g, s); err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (i *Instance) DeleteSubscriptions(subscriptions []models.NotificationSubscription, ctx context.Context) error {
	if len(subscriptions) == 0 {
		return errors.E("Empty subscriptions")
	}

	//g := apmgorm.WithContext(ctx, i.Gorm)
	//for _, s := range subscriptions {
	//	err := g.Where("coin = ? and address = ?", s.Coin, s.Address).Delete(&models.NotificationSubscription{}).Error
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}
