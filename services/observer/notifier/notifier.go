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

type DispatchEvent struct {
	Action blockatlas.TransactionType `json:"action"`
	Result *blockatlas.Tx             `json:"result"`
	Id     uint                       `json:"id"`
}

func RunNotifier(delivery amqp.Delivery) {
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
	subscriptionsDataList, err := db.GetSubscriptionData(txs[0].Coin, addresses)
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
	for _, tx := range tx.Txs() {
		tx.Direction = tx.GetTransactionDirection(sub.Address)
		tx.InferUtxoValue(sub.Address, tx.Coin)
		action := DispatchEvent{
			Action: tx.Type,
			Result: &tx,
			Id:     sub.Id,
		}
		txJson, err := json.Marshal(action)
		if err != nil {
			logger.Panic(err)
		}

		logParams := logger.Params{
			"Id":   sub.Id,
			"coin": sub.Coin,
			"txID": tx.ID,
		}

		publishTransaction(sub.Id, txJson, logParams)
	}
}

func publishTransaction(id uint, rawMessage []byte, logParams logger.Params) {
	err := mq.Transactions.Publish(rawMessage)
	if err != nil {
		err = errors.E(err, "Failed to dispatch event", errors.Params{"id": id}, logParams)
		logger.Fatal(err, logger.Params{"id": id}, logParams)
	}
	logger.Info("Message dispatched", logger.Params{"id": id}, logParams)
}

func GetInterval(value int, minInterval, maxInterval time.Duration) time.Duration {
	interval := time.Duration(value) * time.Millisecond
	pMin := numbers.Max(minInterval.Nanoseconds(), interval.Nanoseconds())
	pMax := numbers.Min(int(maxInterval.Nanoseconds()), int(pMin))
	return time.Duration(pMax)
}
