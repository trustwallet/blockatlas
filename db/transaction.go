package db

import (
	"github.com/trustwallet/blockatlas/db/models"
	"gorm.io/gorm/clause"
)

func (i *Instance) CreateTransactions(txs []models.Transaction) error {
	return i.Gorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&txs).Error
}

func (i *Instance) GetTransactionsByAccount(account string, coin uint, limit int) (txs []models.Transaction, err error) {
	err = i.Gorm.Where("coin = ? AND \"from\" = ?", coin, account).
		Or("coin = ? AND \"to\" = ?", coin, account).
		Order("date DESC").
		Limit(limit).
		Find(&txs).Error
	return
}
