package db

import (
	"github.com/trustwallet/blockatlas/db/models"
	"gorm.io/gorm/clause"
)

func (i *Instance) CreateTransactions(txs []models.Transaction) error {
	return i.Gorm.Clauses(clause.OnConflict{DoNothing: true}).Create(&txs).Error
}

func (i *Instance) GetTransactionsByAccount(account string, coin uint) (txs []models.Transaction, err error) {
	err = i.Gorm.Where("coin = ? AND (\"from\" = ? OR \"to\" = ?)", coin, account, account).
		Order("date DESC").
		Limit(25).
		Find(&txs).Error
	return
}
