package db

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgorm"
	"strings"
	"time"
)

func (i *Instance) GetSubscriptions(coin uint, addresses []string, ctx context.Context) ([]models.Subscription, error) {
	if len(addresses) == 0 {
		return nil, errors.E("Empty addresses")
	}
	g := apmgorm.WithContext(ctx, i.Gorm)
	var subscriptionsDataList []models.Subscription
	err := g.
		Model(&models.Subscription{}).
		Where("address in (?) AND coin = ?", addresses, coin).
		Find(&subscriptionsDataList).Error

	if err != nil {
		return nil, err
	}
	return subscriptionsDataList, nil
}

func (i *Instance) AddSubscriptions(subscriptions []models.Subscription, ctx context.Context) error {
	if len(subscriptions) == 0 {
		return errors.E("Empty subscriptions")
	}
	subscriptionsBatch := convertToBatch(subscriptions, 3000, ctx)
	g := apmgorm.WithContext(ctx, i.Gorm)

	for _, s := range subscriptionsBatch {
		if err := bulkCreate(g, s); err != nil {
			return err
		}
	}

	return nil
}

func (i *Instance) DeleteSubscriptions(subscriptions []models.Subscription, ctx context.Context) error {
	g := apmgorm.WithContext(ctx, i.Gorm)
	for _, s := range subscriptions {
		err := g.Where("coin = ? and address = ?", s.Coin, s.Address).Delete(&models.Subscription{}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

const rawBulkInsert = `INSERT INTO subscriptions(created_at,coin,address) VALUES %s ON CONFLICT DO NOTHING`

func bulkCreate(db *gorm.DB, dataList []models.Subscription) error {
	var (
		valueStrings []string
		valueArgs    []interface{}
	)

	for _, d := range dataList {
		valueStrings = append(valueStrings, "(?, ?, ?)")

		valueArgs = append(valueArgs, time.Now())
		valueArgs = append(valueArgs, d.Coin)
		valueArgs = append(valueArgs, d.Address)
	}

	smt := fmt.Sprintf(rawBulkInsert, strings.Join(valueStrings, ","))

	if err := db.Exec(smt, valueArgs...).Error; err != nil {
		return err
	}

	return nil
}

func convertToBatch(txs []models.Subscription, sizeUint uint, ctx context.Context) [][]models.Subscription {
	span, _ := apm.StartSpan(ctx, "convertToBatch", "app")
	defer span.End()
	size := int(sizeUint)
	resultLength := (len(txs) + size - 1) / size
	result := make([][]models.Subscription, resultLength)
	lo, hi := 0, size
	for i := range result {
		if hi > len(txs) {
			hi = len(txs)
		}
		result[i] = txs[lo:hi:hi]
		lo, hi = hi, hi+size
	}
	return result
}
