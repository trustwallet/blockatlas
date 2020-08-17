package db

import (
	"context"
	"github.com/trustwallet/blockatlas/db/models"
	"go.elastic.co/apm/module/apmgorm"
)

func (i *Instance) AddToken(token models.Token, ctx context.Context) error {
	db := apmgorm.WithContext(ctx, i.Gorm)
	return db.Set("gorm:insert_option", "ON CONFLICT DO NOTHING").Create(&token).Error
}

func (i *Instance) GetTokenByTokenID(tokenID string, ctx context.Context) (models.Token, error) {
	db := apmgorm.WithContext(ctx, i.Gorm)
	var result models.Token
	err := db.Where("token_id = ?", tokenID).First(&result).Error
	return result, err
}
