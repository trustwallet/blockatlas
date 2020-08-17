package notifier

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"go.elastic.co/apm"
)

func GetTransactionsFromDelivery(delivery amqp.Delivery, ctx context.Context) (blockatlas.Txs, error) {
	var txs blockatlas.Txs

	span, _ := apm.StartSpan(ctx, "GetTransactionsFromDelivery", "app")
	defer span.End()

	if err := json.Unmarshal(delivery.Body, &txs); err != nil {
		return nil, err
	}

	logger.Info("Consumed", logger.Params{"txs": len(txs), "coin": txs[0].Coin})

	if len(txs) == 0 {
		return nil, errors.E("empty txs list")
	}
	return txs, nil
}

func publishNotificationBatch(batch []TransactionNotification, ctx context.Context) {
	span, _ := apm.StartSpan(ctx, "getNotificationBatches", "app")
	defer span.End()

	raw, err := json.Marshal(batch)
	if err != nil {
		err = errors.E(err, " failed to dispatch event")
		logger.Fatal(err)
	}
	err = mq.TxNotifications.Publish(raw)
	if err != nil {
		err = errors.E(err, " failed to dispatch event")
		logger.Fatal(err)
	}

	logger.Info("Txs batch dispatched", logger.Params{"txs": len(batch)})
}
