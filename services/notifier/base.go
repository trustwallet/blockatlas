package notifier

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/address"
	"strconv"

	"go.elastic.co/apm"
)

const (
	DefaultPushNotificationsBatchLimit = 50

	Notifier = "Notifier"
)

var MaxPushNotificationsBatchLimit uint = DefaultPushNotificationsBatchLimit

func RunNotifier(database *db.Instance, delivery amqp.Delivery) {
	tx := apm.DefaultTracer.StartTransaction("RunNotifier", "app")
	defer tx.End()
	ctx := apm.ContextWithTransaction(context.Background(), tx)

	defer func() {
		if err := delivery.Ack(false); err != nil {
			log.Error(err)
		}
	}()

	txs, err := GetTransactionsFromDelivery(delivery, Notifier, ctx)
	if err != nil {
		log.Error("failed to get transactions", err)
	}

	allAddresses := make([]string, 0)
	for _, tx := range txs {
		allAddresses = append(allAddresses, tx.GetAddresses()...)
	}

	addresses := ToUniqueAddresses(allAddresses)
	for i := range addresses {
		addresses[i] = strconv.Itoa(int(txs[0].Coin)) + "_" + addresses[i]
	}

	if len(txs) < 1 {
		return
	}
	subscriptionsDataList, err := database.GetSubscriptionsForNotifications(addresses, ctx)
	if err != nil || len(subscriptionsDataList) == 0 {
		return
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
		publishNotificationBatch(batch, ctx)
	}
	log.Info("------------------------------------------------------------")
}
