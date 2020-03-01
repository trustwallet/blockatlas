package observer

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
)

type Dispatcher struct {
	Client http.Client
}

type DispatchEvent struct {
	Action blockatlas.TransactionType `json:"action"`
	Result *blockatlas.Tx             `json:"result"`
	GUID   string                     `json:"guid"`
}

func (d *Dispatcher) Run(events <-chan Event) {
	for event := range events {
		d.dispatch(event)
	}
}

func (d *Dispatcher) dispatch(event Event) {
	guid := event.Subscription.GUID

	action := DispatchEvent{
		Action: event.Tx.Type,
		Result: event.Tx,
		GUID:   guid,
	}

	txJson, err := json.Marshal(action)
	if err != nil {
		logger.Panic(err)
	}

	logParams := logger.Params{
		"guid": guid,
		"coin": event.Subscription.Coin,
		"txID": event.Tx.ID,
	}

	go d.postMessageToQueue(guid, txJson, logParams)

	logger.Info("Dispatching messages...", logParams)
}

func (d *Dispatcher) postMessageToQueue(message string, rawMessage []byte, logParams logger.Params) {
	err := mq.Transactions.Publish(rawMessage)
	if err != nil {
		err = errors.E(err, "Failed to dispatch event", errors.Params{"message": message}, logParams)
		logger.Error(err, logger.Params{"message": message}, logParams)
	}
	logger.Info("Message dispatched", logger.Params{"message": message}, logParams)
}
