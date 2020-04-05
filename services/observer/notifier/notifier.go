package notifier

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"sync"
	"time"
)

const DefaultPushNotificationsBatchLimit = 50

var MaxPushNotificationsBatchLimit uint = DefaultPushNotificationsBatchLimit

type TransactionNotification struct {
	Action blockatlas.TransactionType `json:"action"`
	Result *blockatlas.Tx             `json:"result"`
	Id     uint                       `json:"id"`
}

func RunNotifier(database *db.Instance, delivery amqp.Delivery) {
	defer func() {
		if err := delivery.Ack(false); err != nil {
			logger.Error(err)
		}
	}()
	var txs blockatlas.Txs
	if err := json.Unmarshal(delivery.Body, &txs); err != nil {
		logger.Error(err)
		return
	}
	if len(txs) == 0 {
		return
	}

	logger.Info("Consumed", logger.Params{"txs": len(txs), "coin": txs[0].Coin})

	blockTransactions := txs.GetTransactionsMap()
	if len(blockTransactions.Map) == 0 {
		return
	}

	addresses := blockTransactions.GetUniqueAddresses()
	subscriptionsDataList, err := database.GetSubscriptionData(txs[0].Coin, addresses)
	if err != nil || len(subscriptionsDataList) == 0 {
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(subscriptionsDataList))
	for _, data := range subscriptionsDataList {
		go buildAndPostMessage(
			blockTransactions,
			blockatlas.Subscription{Coin: data.Coin, Address: data.Address, Id: data.SubscriptionId},
			&wg)
	}
	wg.Wait()
}

func buildAndPostMessage(blockTransactions blockatlas.TxSetMap, sub blockatlas.Subscription, wg *sync.WaitGroup) {
	defer wg.Done()

	tx, ok := blockTransactions.Map[sub.Address]
	if !ok {
		return
	}
	notifications := make([]TransactionNotification, 0, len(tx.Txs()))
	for _, tx := range tx.Txs() {
		tx.Direction = tx.GetTransactionDirection(sub.Address)
		tx.InferUtxoValue(sub.Address, tx.Coin)
		notification := TransactionNotification{
			Action: tx.Type,
			Result: &tx,
			Id:     sub.Id,
		}

		logger.Info("Notification ready", logger.Params{"Id": sub.Id, "coin": sub.Coin, "txID": tx.ID})

		notifications = append(notifications, notification)
	}

	batches := getNotificationBatches(notifications, MaxPushNotificationsBatchLimit)

	for _, batch := range batches {
		publishNotificationBatch(batch)
	}
}

func publishNotificationBatch(batch []TransactionNotification) {
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

func GetInterval(value int, minInterval, maxInterval time.Duration) time.Duration {
	interval := time.Duration(value) * time.Millisecond
	pMin := numbers.Max(minInterval.Nanoseconds(), interval.Nanoseconds())
	pMax := numbers.Min(int(maxInterval.Nanoseconds()), int(pMin))
	return time.Duration(pMax)
}

func getNotificationBatches(notifications []TransactionNotification, sizeUint uint) [][]TransactionNotification {
	size := int(sizeUint)
	resultLength := (len(notifications) + size - 1) / size
	result := make([][]TransactionNotification, resultLength)
	lo, hi := 0, size
	for i := range result {
		if hi > len(notifications) {
			hi = len(notifications)
		}
		result[i] = notifications[lo:hi:hi]
		lo, hi = hi, hi+size
	}
	return result
}
