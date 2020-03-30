package db

import (
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

func GetSubscriptionData(coin uint, addresses []string) ([]models.SubscriptionData, error) {
	if len(addresses) == 0 {
		return nil, errors.E("Empty addresses")
	}

	var subscriptionsDataList []models.SubscriptionData

	err := GormDb.
		Model(&models.SubscriptionData{}).
		Where("address in (?) AND coin = ?", addresses, coin).
		Find(&subscriptionsDataList).Error

	if err != nil {
		return nil, err
	}

	return subscriptionsDataList, nil
}

func AddSubscriptions(guid string, subscriptions []models.SubscriptionData) error {
	if len(subscriptions) == 0 {
		return errors.E("Empty subscriptions")
	}

	var (
		existingSub models.Subscription
		err         error
	)

	recordNotFound := GormDb.
		Where(models.Subscription{GUID: guid}).
		First(&existingSub).
		RecordNotFound()

	subscriptions = removeSubscriptionDuplicates(subscriptions)

	if recordNotFound {
		err = AddSubscription(guid, subscriptions)
	} else {
		err = AddToExistingSubscription(guid, subscriptions)
	}

	if err != nil {
		return err
	}

	return nil
}

func AddSubscription(guid string, data []models.SubscriptionData) error {
	return GormDb.Create(&models.Subscription{GUID: guid, Data: data}).Error
}

func AddToExistingSubscription(guid string, newSubscriptions []models.SubscriptionData) error {
	var (
		existingData []models.SubscriptionData
		sub          = &models.Subscription{GUID: guid}
	)

	if err := GormDb.Model(sub).Association("Data").Find(&existingData).Error; err != nil {
		return err
	}

	updateList, deleteList := getSubscriptionsToDeleteAndUpdate(existingData, newSubscriptions)

	if len(updateList) > 0 {
		if err := GormDb.Model(sub).Association("Data").Append(updateList).Error; err != nil {
			return err
		}
	}
	if len(deleteList) > 0 {
		if err := DeleteSubscriptions(deleteList); err != nil {
			return err
		}
	}
	return nil
}

func DeleteSubscriptions(subscriptions []models.SubscriptionData) error {
	var (
		errorsList = make([]error, 0)
		errDetails string
	)

	for _, sub := range subscriptions {
		if err := GormDb.Delete(&models.SubscriptionData{}, sub).Error; err != nil {
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
		if s.Address == sub.Address && sub.Coin == s.Coin && sub.SubscriptionId == sub.SubscriptionId {
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
