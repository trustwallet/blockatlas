package notifier

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/address"
)

const (
	Notifier = "Notifier"
)

func RunNotifier(database *db.Instance, delivery amqp.Delivery) {
	defer func() {
		if err := delivery.Ack(false); err != nil {
			log.WithFields(log.Fields{"service": Notifier}).Error(err)
		}
	}()

	txs, err := GetTransactionsFromDelivery(delivery, Notifier)
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

	if len(txs) == 0 {
		return
	}
	subscriptionsDataList, err := database.GetSubscriptionsForNotifications(addresses)
	if err != nil || len(subscriptionsDataList) == 0 {
		return
	}

	notifications := make([]TransactionNotification, 0)
	for _, sub := range subscriptionsDataList {
		ua, _, ok := address.UnprefixedAddress(sub.Address.Address)
		if !ok {
			continue
		}
		notificationsForAddress := buildNotificationsByAddress(ua, txs)
		notifications = append(notifications, notificationsForAddress...)
	}
	err = publishNotifications(notifications)
	if err != nil {
		log.WithFields(log.Fields{"service": Notifier}).Error(err)
	}

	log.Info("------------------------------------------------------------")
}
