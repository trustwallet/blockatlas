package tokenindexer

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/notifier"
	"go.elastic.co/apm"
)

const TokenIndexer = "TokenIndexer"

type TokenIndexerConsumer struct {
	Database *db.Instance
}

func (c *TokenIndexerConsumer) GetQueue() string {
	return string(new_mq.RawTransactionsTokenIndexer)
}

func (c *TokenIndexerConsumer) Callback(msg amqp.Delivery) error {
	tx := apm.DefaultTracer.StartTransaction("RunTokenIndexer", "app")
	defer tx.End()
	ctx := apm.ContextWithTransaction(context.Background(), tx)

	txs, err := notifier.GetTransactionsFromDelivery(msg, TokenIndexer, ctx)
	if err != nil {
		log.WithFields(log.Fields{"service": TokenIndexer}).Error("failed to get transactions", err)
		return err
	}
	if len(txs) == 0 {
		return nil
	}

	assets := GetAssetsFromTransactions(txs)
	err = c.Database.AddNewAssets(assets, ctx)
	if err != nil {
		log.WithFields(log.Fields{"service": TokenIndexer}).Error("failed to add assets", err)
		return err
	}
	log.Info("------------------------------------------------------------")
	return nil
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
