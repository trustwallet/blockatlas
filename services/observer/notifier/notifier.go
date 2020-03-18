package notifier

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"github.com/trustwallet/blockatlas/services/observer/parser"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

type DispatchEvent struct {
	Action blockatlas.TransactionType `json:"action"`
	Result *blockatlas.Tx             `json:"result"`
	GUID   string                     `json:"guid"`
}

func RunNotifier(delivery amqp.Delivery, s storage.Addresses) {
	var blockData parser.BlockData
	if err := json.Unmarshal(delivery.Body, &blockData); err != nil {
		logger.Error(err)
		return
	}

	txMap := blockData.Block.GetTransactionsMap()
	if len(txMap) == 0 {
		return
	}

	// Build list of unique addresses
	var addresses []string
	for address := range txMap {
		if len(address) == 0 {
			continue
		}
		addresses = append(addresses, address)
	}

	// Lookup subscriptions
	subs, err := s.Lookup(blockData.Coin, addresses)
	if err != nil || len(subs) == 0 {
		return
	}
	for _, sub := range subs {
		tx, ok := txMap[sub.Address]
		if !ok {
			continue
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

			go publishTransaction(sub.GUID, txJson, logParams)
		}
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