package db

import (
	"github.com/trustwallet/blockatlas/db/models"
	"gorm.io/gorm/clause"
)

func (i *Instance) CreateTransactions(txs []models.Transaction) error {
	return i.Gorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&txs).Error
}
