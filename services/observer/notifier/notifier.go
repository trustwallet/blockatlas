package notifier

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"github.com/trustwallet/blockatlas/storage"
	"sync"
	"time"
)

type DispatchEvent struct {
	Action blockatlas.TransactionType `json:"action"`
	Result *blockatlas.Tx             `json:"result"`
	GUID   string                     `json:"guid"`
}

func RunNotifier(delivery amqp.Delivery, s storage.Addresses) {
	defer func() {
		if err := delivery.Ack(false); err != nil {
			logger.Error(err)
		}
	}()
	var block blockatlas.Block
	if err := json.Unmarshal(delivery.Body, &block); err != nil {
		logger.Error(err)
		return
	}
	if len(block.Txs) == 0 {
		return
	}

	logger.Info("Consumed", logger.Params{"num": block.Number, "txs": len(block.Txs), "coin": block.Txs[0].Coin})

	blockTransactions := block.GetTransactionsMap()
	if len(blockTransactions.Map) == 0 {
		return
	}

	addresses := blockTransactions.GetUniqueAddresses()
	subs, err := s.FindSubscriptions(block.Txs[0].Coin, addresses)
	if err != nil || len(subs) == 0 {
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(subs))
	for _, sub := range subs {
		go buildAndPostMessage(blockTransactions, sub, &wg)
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
			GUID:   sub.GUID,
		}
		txJson, err := json.Marshal(action)
		if err != nil {
			logger.Panic(err)
		}

		logParams := logger.Params{
			"guid": sub.GUID,
			"coin": sub.Coin,
			"txID": tx.ID,
		}

		publishTransaction(sub.GUID, txJson, logParams)
	}
}

func publishTransaction(message string, rawMessage []byte, logParams logger.Params) {
	err := mq.Transactions.Publish(rawMessage)
	if err != nil {
		err = errors.E(err, "Failed to dispatch event", errors.Params{"message": message}, logParams)
		logger.Fatal(err, logger.Params{"message": message}, logParams)
	}
	logger.Info("Message dispatched", logger.Params{"message": message}, logParams)
}

func GetInterval(value int, minInterval, maxInterval time.Duration) time.Duration {
	interval := time.Duration(value) * time.Millisecond
	pMin := numbers.Max(minInterval.Nanoseconds(), interval.Nanoseconds())
	pMax := numbers.Min(int(maxInterval.Nanoseconds()), int(pMin))
	return time.Duration(pMax)
}
