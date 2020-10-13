package db

import (
	"context"
	"github.com/trustwallet/blockatlas/db/models"
	"gorm.io/gorm/clause"
)

func (i Instance) CreateTokenTypes(ctx context.Context, tokenTypes []string) error {
	db := i.Gorm.WithContext(ctx)
	tokenTypesModel := make([]models.TokenType, 0, len(tokenTypes))
	for _, v := range tokenTypes {
		tokenTypesModel = append(tokenTypesModel, models.TokenType{
			Type: v,
		})
	}
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&tokenTypesModel).Error
}

func (i Instance) GetTokenTypes(ctx context.Context, filter []string) ([]models.TokenType, error) {
	db := i.Gorm.WithContext(ctx)
	var tokenTypes []models.TokenType
	if len(filter) > 0 {
		if err := db.Find(&tokenTypes, "type in (?)", filter).Error; err != nil {
			return nil, err
		}
	}

	if err := db.Find(&tokenTypes).Error; err != nil {
		return nil, err
	}
	return tokenTypes, nil
}
