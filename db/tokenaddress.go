package db

import (
	"context"
	"fmt"
	"github.com/trustwallet/blockatlas/db/models"
	"go.elastic.co/apm/module/apmgorm"
)

func (i *Instance) AddAssociationToAddress(association models.AddressToTokenAssociation, ctx context.Context) error {
	db := apmgorm.WithContext(ctx, i.Gorm)
	return db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING").Create(&association).Error
}

func (i *Instance) GetAssociationsByAddresses(addresses []string, ctx context.Context) ([]models.AddressToTokenAssociation, error) {
	//db := apmgorm.WithContext(ctx, i.Gorm)
	//var result []models.AddressToTokenAssociation
	//err := db.Where("address in (?)", addresses).Find(&result).Error
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}

func (i *Instance) AddAssociations(associations map[string][]string, ctx context.Context) error {
	db := apmgorm.WithContext(ctx, i.Gorm)
	for address, assets := range associations {
		var addressFromDB models.Address
		err := db.Where("address = ?", address).First(&addressFromDB).Error
		if err != nil {
			// bulk create of assets
		}
		// add association for all assets with current address
		fmt.Println(assets)
	}
	return db.Error
}
