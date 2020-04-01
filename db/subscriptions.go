package db

import (
	"fmt"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"strings"
)

func (i *Instance) GetSubscriptionData(coin uint, addresses []string) ([]models.SubscriptionData, error) {
	if len(addresses) == 0 {
		return nil, errors.E("Empty addresses")
	}

	var subscriptionsDataList []models.SubscriptionData
	err := i.DB.
		Model(&models.SubscriptionData{}).
		Where("address in (?) AND coin = ?", addresses, coin).
		Find(&subscriptionsDataList).Error

	if err != nil {
		return nil, err
	}
	return subscriptionsDataList, nil
}

func (i *Instance) AddSubscriptions(id uint, subscriptions []models.SubscriptionData) error {
	txInstance := Instance{DB: i.DB.Begin()}
	defer func() {
		if r := recover(); r != nil {
			txInstance.DB.Rollback()
		}
	}()

	if err := txInstance.DB.Error; err != nil {
		return err
	}
	if len(subscriptions) == 0 {
		return errors.E("Empty subscriptions")
	}
	var (
		existingSub models.Subscription
		err         error
	)

	recordNotFound := txInstance.DB.
		Where(models.Subscription{SubscriptionId: id}).
		First(&existingSub).
		RecordNotFound()

	subscriptions = removeSubscriptionDuplicates(subscriptions)
	if recordNotFound {
		if err = txInstance.DB.Create(&models.Subscription{SubscriptionId: id}).Error; err != nil {
			txInstance.DB.Rollback()
			return err
		}
		err = txInstance.BulkCreate(subscriptions)
	} else {
		err = txInstance.AddToExistingSubscription(id, subscriptions)
	}

	if err != nil {
		txInstance.DB.Rollback()
		return err
	}
	return txInstance.DB.Commit().Error
}

func (i *Instance) AddToExistingSubscription(id uint, subscriptions []models.SubscriptionData) error {
	var (
		existingData []models.SubscriptionData
		association  = i.DB.Model(&models.Subscription{SubscriptionId: id}).Association("Data")
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
	return nil
}

func (i *Instance) DeleteAllSubscriptions(id uint) error {
	request := i.DB.Where("subscription_id = ?", id)
	if err := request.Error; err != nil {
		return err
	}
	return request.Delete(&models.SubscriptionData{}).Error
}

func (i *Instance) DeleteSubscriptions(subscriptions []models.SubscriptionData) error {
	var (
		errorsList = make([]error, 0)
		errDetails string
	)
	for _, sub := range subscriptions {
		if err := i.DB.Delete(&models.SubscriptionData{}, sub).Error; err != nil {
			errorsList = append(errorsList, err)
		}
	}
	if len(errorsList) != 0 {
		for _, err := range errorsList {
			errDetails += err.Error() + " "
		}
		return errors.E(errDetails)
	}
	return nil
}

func (i *Instance) BulkCreate(fs []models.SubscriptionData) error {
	var (
		valueStrings []string
		valueArgs    []interface{}
	)

	for _, f := range fs {
		valueStrings = append(valueStrings, "(?, ?, ?)")

		valueArgs = append(valueArgs, f.SubscriptionId)
		valueArgs = append(valueArgs, f.Coin)
		valueArgs = append(valueArgs, f.Address)
	}

	smt := `INSERT INTO subscription_data(subscription_id, coin, address) VALUES %s`

	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))

	if err := i.DB.Exec(smt, valueArgs...).Error; err != nil {
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
