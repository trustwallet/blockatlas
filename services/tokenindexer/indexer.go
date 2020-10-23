package tokenindexer

import (
	"context"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/notifier"
	"go.elastic.co/apm"
)

func RunTokenIndexer(database *db.Instance, delivery amqp.Delivery) {
	tx := apm.DefaultTracer.StartTransaction("RunTokenIndexer", "app")
	defer tx.End()
	defer func() {
		if err := delivery.Ack(false); err != nil {
			logger.Error(err)
		}
	}()
	ctx := apm.ContextWithTransaction(context.Background(), tx)

	txs, err := notifier.GetTransactionsFromDelivery(delivery, ctx)
	if err != nil {
		logger.Error("failed to get transactions", err)
		if err := delivery.Ack(false); err != nil {
			logger.Error(err)
		}
		return
	}
	if len(txs) == 0 {
		return
	}

	assets := GetAssetsFromTransactions(txs)
	err = database.AddNewAssets(assets, ctx)
	if err != nil {
		logger.Error("failed to add assets", err)
		return
	}
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
