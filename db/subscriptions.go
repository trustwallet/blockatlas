package db

import (
	"github.com/jinzhu/gorm"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

func GetSubscriptionData(db *gorm.DB, coin uint, addresses []string) ([]models.SubscriptionData, error) {
	if len(addresses) == 0 {
		return nil, errors.E("Empty addresses")
	}

	var subscriptionsDataList []models.SubscriptionData

	err := db.
		Model(&models.SubscriptionData{}).
		Where("address in (?) AND coin = ?", addresses, coin).
		Find(&subscriptionsDataList).Error

	if err != nil {
		return nil, err
	}

	return subscriptionsDataList, nil
}

func AddSubscriptions(db *gorm.DB, id uint, subscriptions []models.SubscriptionData) error {
	if len(subscriptions) == 0 {
		return errors.E("Empty subscriptions")
	}

	var (
		existingSub models.Subscription
		err         error
	)

	recordNotFound := db.
		Where(models.Subscription{SubscriptionId: id}).
		First(&existingSub).
		RecordNotFound()

	subscriptions = removeSubscriptionDuplicates(subscriptions)

	if recordNotFound {
		err = AddSubscription(db, id, subscriptions)
	} else {
		err = AddToExistingSubscription(db, id, subscriptions)
	}

	if err != nil {
		return err
	}

	return nil
}

func AddSubscription(db *gorm.DB, id uint, data []models.SubscriptionData) error {
	return db.Create(&models.Subscription{SubscriptionId: id, Data: data}).Error
}

func AddToExistingSubscription(db *gorm.DB, id uint, subscriptions []models.SubscriptionData) error {
	var (
		existingData []models.SubscriptionData
		sub          = &models.Subscription{SubscriptionId: id}
	)

	association := db.Model(sub).Association("Data")

	if err := association.Error; err != nil {
		return err
	}

	if err := association.Find(&existingData).Error; err != nil {
		return err
	}

	updateList, deleteList := getSubscriptionsToDeleteAndUpdate(existingData, subscriptions)

	if len(updateList) > 0 {
		if err := association.Append(updateList).Error; err != nil {
			return err
		}
	}
	if len(deleteList) > 0 {
		if err := DeleteSubscriptions(db, deleteList); err != nil {
			return err
		}
	}
	return nil
}

func DeleteAllSubscriptions(db *gorm.DB, id uint) error {
	return db.Where("subscription_id = ?", id).Delete(&models.SubscriptionData{}).Error
}

func DeleteSubscriptions(db *gorm.DB, subscriptions []models.SubscriptionData) error {
	var (
		errorsList = make([]error, 0)
		errDetails string
	)

	for _, sub := range subscriptions {
		if err := db.Delete(&models.SubscriptionData{}, sub).Error; err != nil {
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
