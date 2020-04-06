package db

import (
	"fmt"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"strconv"
	"strings"
	"time"
)

const rawBulkInsert = `INSERT INTO subscription_data(subscription_id, coin, address) VALUES %s ON CONFLICT DO NOTHING`

func (i *Instance) GetSubscriptionData(coin uint, addresses []string) ([]models.SubscriptionData, error) {
	if len(addresses) == 0 {
		return nil, errors.E("Empty addresses")
	}

	var subscriptionsDataList []models.SubscriptionData
	err := i.Gorm.
		Model(&models.SubscriptionData{}).
		Where("address in (?) AND coin = ?", addresses, coin).
		Find(&subscriptionsDataList).Error

	if err != nil {
		return nil, err
	}
	return subscriptionsDataList, nil
}

func (i *Instance) AddSubscriptions(id uint, subscriptions []models.SubscriptionData) error {
	if len(subscriptions) == 0 {
		return errors.E("Empty subscriptions")
	}

	txInstance := Instance{Gorm: i.Gorm.Begin()}
	defer func() {
		if r := recover(); r != nil {
			txInstance.Gorm.Rollback()
		}
	}()

	if err := txInstance.Gorm.Error; err != nil {
		return err
	}

	var (
		existingSub models.Subscription
		err         error
	)

	recordNotFound := txInstance.Gorm.
		Where(models.Subscription{SubscriptionId: id}).
		First(&existingSub).
		RecordNotFound()

	subscriptions = removeSubscriptionDuplicates(subscriptions)
	if recordNotFound {
		err = txInstance.Gorm.Set("gorm:insert_option",
			"ON CONFLICT (subscription_id) DO UPDATE SET subscription_id = excluded.subscription_id").
			Create(&models.Subscription{SubscriptionId: id, UpdatedAt: time.Now()}).Error
		if err != nil {
			txInstance.Gorm.Rollback()
			return err
		}
		err = txInstance.BulkCreate(subscriptions)
	} else {
		err = txInstance.AddToExistingSubscription(id, subscriptions)
	}

	if err != nil {
		txInstance.Gorm.Rollback()
		return err
	}
	return txInstance.Gorm.Commit().Error
}

func (i *Instance) AddToExistingSubscription(id uint, subscriptions []models.SubscriptionData) error {
	var (
		existingData []models.SubscriptionData
		association  = i.Gorm.Model(&models.Subscription{SubscriptionId: id}).Association("Data")
	)
	if err := association.Error; err != nil {
		return err
	}
	if err := association.Find(&existingData).Error; err != nil {
		return err
	}

	updateList, deleteList := getSubscriptionsToDeleteAndUpdate(existingData, subscriptions)
	if len(updateList) > 0 {
		if err := i.BulkCreate(updateList); err != nil {
			return err
		}
	}
	if len(deleteList) > 0 {
		if err := i.DeleteSubscriptions(deleteList); err != nil {
			return err
		}
	}

	if err := i.Gorm.Model(&models.Subscription{SubscriptionId: id}).Update("updated_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func (i *Instance) DeleteAllSubscriptions(id uint) error {
	request := i.Gorm.Where("subscription_id = ?", id)
	if err := request.Error; err != nil {
		return err
	}
	return request.Delete(&models.SubscriptionData{}).Error
}

func (i *Instance) DeleteSubscriptions(subscriptions []models.SubscriptionData) error {
	var idList = make([]string, 0, len(subscriptions))

	for _, sub := range subscriptions {
		idList = append(idList, strconv.Itoa(int(sub.ID)))
	}

	request := i.Gorm.Where("id in (?)", idList)

	if err := request.Error; err != nil {
		return err
	}
	if err := request.Delete(&models.SubscriptionData{}).Error; err != nil {
		return err
	}

	return nil
}

func (i *Instance) BulkCreate(dataList []models.SubscriptionData) error {
	var (
		valueStrings []string
		valueArgs    []interface{}
	)

	for _, d := range dataList {
		valueStrings = append(valueStrings, "(?, ?, ?)")

		valueArgs = append(valueArgs, d.SubscriptionId)
		valueArgs = append(valueArgs, d.Coin)
		valueArgs = append(valueArgs, d.Address)
	}

	smt := fmt.Sprintf(rawBulkInsert, strings.Join(valueStrings, ","))

	if err := i.Gorm.Exec(smt, valueArgs...).Error; err != nil {
		return err
	}

	return nil
}

func getSubscriptionsToDeleteAndUpdate(existing, new []models.SubscriptionData) (subToUpdate, subToDelete []models.SubscriptionData) {
	for _, n := range new {
		if !containSubscription(n, existing) {
			subToUpdate = append(subToUpdate, n)
		}
	}
	for _, e := range existing {
		if !containSubscription(e, new) {
			subToDelete = append(subToDelete, e)
		}
	}
	return subToUpdate, subToDelete
}

func containSubscription(sub models.SubscriptionData, list []models.SubscriptionData) bool {
	for _, s := range list {
		if s.Address == sub.Address && sub.Coin == s.Coin && s.SubscriptionId == sub.SubscriptionId {
			return true
		}
	}
	return false
}

func removeSubscriptionDuplicates(sub []models.SubscriptionData) []models.SubscriptionData {
	keys := make(map[models.SubscriptionData]bool)
	result := make([]models.SubscriptionData, 0)
	for _, entry := range sub {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}
