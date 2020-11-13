package tokenindexer

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/notifier"
	"go.elastic.co/apm"
)

const TokenIndexer = "TokenIndexer"

func RunTokenIndexer(database *db.Instance, delivery amqp.Delivery) {
	tx := apm.DefaultTracer.StartTransaction("RunTokenIndexer", "app")
	defer tx.End()
	defer func() {
		if err := delivery.Ack(false); err != nil {
			log.WithFields(log.Fields{"service": TokenIndexer}).Error(err)
		}
	}()
	ctx := apm.ContextWithTransaction(context.Background(), tx)

	txs, err := notifier.GetTransactionsFromDelivery(delivery, TokenIndexer, ctx)
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
	err = database.AddNewAssets(assets, ctx)
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
