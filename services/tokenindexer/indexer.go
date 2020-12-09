package tokenindexer

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/notifier"
)

const TokenIndexer = "TokenIndexer"

func RunTokenIndexer(database *db.Instance, delivery amqp.Delivery) {
	defer func() {
		if err := delivery.Ack(false); err != nil {
			log.WithFields(log.Fields{"service": TokenIndexer}).Error(err)
		}
	}()

	txs, err := notifier.GetTransactionsFromDelivery(delivery, TokenIndexer)
	if err != nil {
		log.WithFields(log.Fields{"service": TokenIndexer}).Error("failed to get transactions", err)
		if err := delivery.Ack(false); err != nil {
			log.WithFields(log.Fields{"service": TokenIndexer}).Error(err)
		}
		return
	}
	if len(txs) == 0 {
		return
	}

	assets := GetAssetsFromTransactions(txs)
	err = database.AddNewAssets(assets)
	if err != nil {
		log.WithFields(log.Fields{"service": TokenIndexer}).Error("failed to add assets", err)
		return
	}
	log.Info("------------------------------------------------------------")
}

func GetAssetsFromTransactions(txs []blockatlas.Tx) []models.Asset {
	var result []models.Asset
	for _, tx := range txs {
		a, ok := tx.AssetModel()
		if !ok {
			continue
		}
		if a.Asset == "" {
			continue
		}
		result = append(result, a)
	}
	return result
}
