package notifier

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func GetTransactionsFromDelivery(delivery amqp.Delivery, service string) (blockatlas.Txs, error) {
	var transactions blockatlas.Txs

	if err := json.Unmarshal(delivery.Body, &transactions); err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{"service": service, "notifications": len(transactions)}).Info("Consumed")

	return transactions, nil
}

func publishNotifications(notifications []TransactionNotification) {
	raw, err := json.Marshal(notifications)
	if err != nil {
		log.Fatal("publishNotifications marshal: ", err)
	}
	err = mq.TxNotifications.Publish(raw)
	if err != nil {
		log.Fatal("publishNotifications publish:", err)
	}

	log.WithFields(log.Fields{"service": Notifier, "notifications": len(notifications)}).Info("Notifications sent")
}
