package notifier

import (
	"encoding/json"
	"errors"

	"github.com/trustwallet/golibs/coin"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func GetTransactionsFromDelivery(delivery amqp.Delivery, service string) (blockatlas.Txs, error) {
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

func publishNotificationBatch(batch []TransactionNotification) {
	raw, err := json.Marshal(batch)
	if err != nil {
		log.Fatal("publishNotificationBatch marshal: ", err)
	}
	err = mq.TxNotifications.Publish(raw)
	if err != nil {
		log.Fatal("publishNotificationBatch publish:", err)
	}

	log.WithFields(log.Fields{"service": Notifier, "txs": len(batch)}).Info("Txs batch dispatched")
}
