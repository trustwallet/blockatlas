package notifier

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
	"strconv"

	"go.elastic.co/apm"
)

const (
	DefaultPushNotificationsBatchLimit = 50
	Notifier                           = "Notifier"
)

var MaxPushNotificationsBatchLimit uint = DefaultPushNotificationsBatchLimit

type NotifierConsumer struct {
	Database *db.Instance
	MQClient *new_mq.Client
}

func (c *NotifierConsumer) GetQueue() string {
	return string(new_mq.RawTransactions)
}

func (c *NotifierConsumer) Callback(msg amqp.Delivery) error {
	tx := apm.DefaultTracer.StartTransaction("RunNotifier", "app")
	defer tx.End()
	ctx := apm.ContextWithTransaction(context.Background(), tx)
	txs, err := GetTransactionsFromDelivery(msg, Notifier, ctx)
	if err != nil {
		log.Error("failed to get transactions", err)
	}
	if len(txs) < 1 {
		return nil
	}
	allAddresses := make([]string, 0)
	for _, tx := range txs {
		allAddresses = append(allAddresses, tx.GetAddresses()...)
	}
	addresses := ToUniqueAddresses(allAddresses)
	for i := range addresses {
		addresses[i] = strconv.Itoa(int(txs[0].Coin)) + "_" + addresses[i]
	}
	subscriptionsDataList, err := c.Database.GetSubscriptionsForNotifications(addresses, ctx)
	if err != nil || len(subscriptionsDataList) == 0 {
		return nil
	}

	notifications := make([]TransactionNotification, 0)
	for _, sub := range subscriptionsDataList {
		ua, _, ok := address.UnprefixedAddress(sub.Address.Address)
		if !ok {
			continue
		}
		notificationsForAddress := buildNotificationsByAddress(ua, txs, ctx)
		notifications = append(notifications, notificationsForAddress...)
	}

	batches := getNotificationBatches(notifications, MaxPushNotificationsBatchLimit, ctx)

	for _, batch := range batches {
		publishNotificationBatch(c.MQClient, batch, ctx)
	}
	log.Info("------------------------------------------------------------")
	return nil
}

func GetTransactionsFromDelivery(delivery amqp.Delivery, service string, ctx context.Context) (blockatlas.Txs, error) {
	span, _ := apm.StartSpan(ctx, "GetTransactionsFromDelivery", "app")
	defer span.End()

	var txs blockatlas.Txs
	if err := json.Unmarshal(delivery.Body, &txs); err != nil {
		return nil, err
	}
	if len(txs) == 0 {
		return nil, errors.New("empty txs list")
	}
	log.WithFields(log.Fields{"service": service, "txs": len(txs), "coin": coin.Coins[txs[0].Coin].Handle}).Info("Consumed")
	return txs, nil
}

func publishNotificationBatch(mqClient *new_mq.Client, batch []TransactionNotification, ctx context.Context) {
	span, _ := apm.StartSpan(ctx, "getNotificationBatches", "app")
	defer span.End()

	raw, err := json.Marshal(batch)
	if err != nil {
		log.Fatal("publishNotificationBatch marshal: ", err)
	}

	err = mqClient.Push(new_mq.TxNotifications, raw)
	if err != nil {
		log.Fatal("publishNotificationBatch publish:", err)
	}
	log.WithFields(log.Fields{"service": Notifier, "txs": len(batch)}).Info("Txs batch dispatched")
}
