package db

import (
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/golibs/types"
	"gorm.io/gorm/clause"
)

func (i *Instance) CreateSubscriptions(addresses []types.Subscription) error {
	if len(addresses) == 0 {
		return nil
	}
	// remove duplicates
	addressIds := make(map[string]bool)
	for _, address := range addresses {
		addressIds[address.AddressID()] = true
	}
	result := make([]models.Subscription, 0)
	for addressId := range addressIds {
		result = append(result, models.Subscription{Address: addressId})
	}

	return i.Gorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&result).Error
}

func (i *Instance) GetSubscriptions(addresses []string) ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	err := i.Gorm.Find(&subscriptions, "address in ?", addresses).Error
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (i *Instance) DeleteSubscriptions(addresses []string) error {
	subscriptions, err := i.GetSubscriptions(addresses)
	if err != nil {
		return err
	}
	if len(subscriptions) == 0 {
		return nil
	}
	subscriptionsIds := make([]uint, 0)
	for _, subscription := range subscriptions {
		subscriptionsIds = append(subscriptionsIds, subscription.ID)
	}
	if err = i.Gorm.
		Where("subscription_id in (?)", subscriptionsIds).
		Delete(&models.SubscriptionsAssetAssociation{}).Error; err != nil {
		return err
	}
	return i.Gorm.
		Where("id in (?)", subscriptionsIds).
		Delete(&models.Subscription{}).Error
}

func (i *Instance) CreateSubscriptionsAssets(associations []models.SubscriptionsAssetAssociation) error {
	if len(associations) == 0 {
		return nil
	}
	return i.Gorm.Clauses(
		clause.OnConflict{
			OnConstraint: "subscriptions_asset_associations_pkey",
			DoUpdates:    clause.AssignmentColumns([]string{"updated_at"})},
	).Create(&associations).Error
}
