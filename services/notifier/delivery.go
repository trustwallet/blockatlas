package notifier

import (
	"encoding/json"

	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/golibs/types"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func GetTransactionsFromDelivery(delivery amqp.Delivery, service string) (types.Txs, error) {
	var transactions types.Txs

	if err := json.Unmarshal(delivery.Body, &transactions); err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{"service": service, "notifications": len(transactions)}).Info("Consumed")

	return transactions, nil
}

func publishNotifications(notifications []types.TransactionNotification) error {
	raw, err := json.Marshal(notifications)
	if err != nil {
		return err
	}
	err = internal.TxNotifications.Publish(raw)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{"service": Notifier, "notifications": len(notifications)}).Info("Notifications send")

	return nil
}
