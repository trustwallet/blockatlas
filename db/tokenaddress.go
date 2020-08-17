package db

import (
	"context"
	"errors"
	"github.com/trustwallet/blockatlas/db/models"
	"go.elastic.co/apm/module/apmgorm"
)

func (i *Instance) AddTokenToAddress(association models.AddressToTokenAssociation, ctx context.Context) error {
	db := apmgorm.WithContext(ctx, i.Gorm)
	var result models.AddressToTokenAssociation
	err := db.Where("address = ?", association.Address).First(&result).Error
	if err != nil {
		return err
	}
	if association.LastUpdatedBlock < result.LastUpdatedBlock {
		return errors.New("block at db > new block ")
	}
	return db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING").Create(&association).Error
}

func (i *Instance) GetTokensByAddresses(addresses []string, ctx context.Context) ([]models.AddressToTokenAssociation, error) {
	db := apmgorm.WithContext(ctx, i.Gorm)
	var result []models.AddressToTokenAssociation
	err := db.Where("address in (?)", addresses).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i *Instance) AddAssociations(associations []models.AddressToTokenAssociation, ctx context.Context) error {
	return nil
}
