package db

import (
	"context"
	"github.com/trustwallet/blockatlas/db/models"
	"gorm.io/gorm/clause"
)

func (i Instance) CreateTokenType(ctx context.Context, tokenType []models.TokenType) error {
	db := i.Gorm.WithContext(ctx)
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&tokenType).Error
}
