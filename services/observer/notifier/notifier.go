package notifier

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"

	"github.com/trustwallet/blockatlas/pkg/logger"

	"go.elastic.co/apm"
)

const DefaultPushNotificationsBatchLimit = 50

var MaxPushNotificationsBatchLimit uint = DefaultPushNotificationsBatchLimit

type TransactionNotification struct {
	Action blockatlas.TransactionType `json:"action"`
	Result *blockatlas.Tx             `json:"result"`
}

func RunNotifier(database *db.Instance, delivery amqp.Delivery) {
	tx := apm.DefaultTracer.StartTransaction("RunNotifier", "app")
	defer tx.End()
	ctx := apm.ContextWithTransaction(context.Background(), tx)

	defer func() {
		if err := delivery.Ack(false); err != nil {
			logger.Error(err)
		}
	}()

	txs, err := getTransactionsFromDelivery(delivery, ctx)
	if err != nil {
		logger.Error("failed to get transactions", err)
	}

	addresses := uniqueAddresses(getAddressesFromTransactions(txs))

	subscriptionsDataList, err := database.GetSubscriptions(txs[0].Coin, addresses, ctx)
	if err != nil || len(subscriptionsDataList) == 0 {
		return
	}

	notifications := make([]TransactionNotification, 0)
	for _, sub := range subscriptionsDataList {
		notificationsForAddress := buildNotificationsByAddress(sub.Address, txs, ctx)
		notifications = append(notifications, notificationsForAddress...)
	}

	batches := getNotificationBatches(notifications, MaxPushNotificationsBatchLimit, ctx)

	for _, batch := range batches {
		publishNotificationBatch(batch, ctx)
	}
}

func getTransactionsFromDelivery(delivery amqp.Delivery, ctx context.Context) (blockatlas.Txs, error) {
	var txs blockatlas.Txs

	span, ctx := apm.StartSpan(ctx, "getTransactionsFromDelivery", "app")
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

func getAddressesFromTransactions(txs blockatlas.Txs) []string {
	result := make([]string, 0)
	for _, tx := range txs {
		result = append(result, tx.GetAddresses()...)
	}
	return result
}

func uniqueAddresses(addresses []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range addresses {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func buildNotificationsByAddress(address string, txs blockatlas.Txs, ctx context.Context) []TransactionNotification {
	span, ctx := apm.StartSpan(ctx, "buildNotification", "app")
	defer span.End()

	transactionsByAddress := uniqueTransactions(findTransactionsByAddress(txs, address))

	result := make([]TransactionNotification, 0)
	for _, tx := range transactionsByAddress {
		tx.Direction = tx.GetTransactionDirection(address)
		tx.InferUtxoValue(address, tx.Coin)
		result = append(result, TransactionNotification{Action: tx.Type, Result: &tx})
	}

	return result
}

func uniqueTransactions(txs []blockatlas.Tx) []blockatlas.Tx {
	keys := make(map[string]bool)
	var list []blockatlas.Tx
	for _, entry := range txs {
		key := entry.ID + string(entry.Direction)
		if _, value := keys[key]; !value {
			keys[key] = true
			list = append(list, entry)
		}
	}
	return list
}

func findTransactionsByAddress(txs blockatlas.Txs, address string) []blockatlas.Tx {
	result := make([]blockatlas.Tx, 0)
	for _, tx := range txs {
		if containsAddress(tx, address) {
			result = append(result, tx)
		}
	}
	return result
}

func containsAddress(tx blockatlas.Tx, address string) bool {
	txAddresses := uniqueAddresses(tx.GetAddresses())
	for _, a := range txAddresses {
		if a == address {
			return true
		}
	}
	return false
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

func getNotificationBatches(notifications []TransactionNotification, sizeUint uint, ctx context.Context) [][]TransactionNotification {
	span, _ := apm.StartSpan(ctx, "getNotificationBatches", "app")
	defer span.End()
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
