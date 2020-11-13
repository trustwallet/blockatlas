package notifier

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"go.elastic.co/apm"
)

func GetTransactionsFromDelivery(delivery amqp.Delivery, service string, ctx context.Context) (blockatlas.Txs, error) {
	var txs blockatlas.Txs

	span, _ := apm.StartSpan(ctx, "GetTransactionsFromDelivery", "app")
	defer span.End()

	if err := json.Unmarshal(delivery.Body, &txs); err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{"service": service, "txs": len(txs), "coin": txs[0].Coin}).Info("Consumed")

	if len(txs) == 0 {
		return nil, errors.New("empty txs list")
	}
	return txs, nil
}

func publishNotificationBatch(batch []TransactionNotification, ctx context.Context) {
	span, _ := apm.StartSpan(ctx, "getNotificationBatches", "app")
	defer span.End()

	raw, err := json.Marshal(batch)
	if err != nil {
		log.Fatal(err)
	}
	err = mq.TxNotifications.Publish(raw)
	if err != nil {
		log.Fatal(err)
	}

	log.WithFields(log.Fields{"service": Notifier, "txs": len(batch)}).Info("Txs batch dispatched")
}
