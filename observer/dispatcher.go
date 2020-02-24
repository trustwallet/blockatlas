package observer

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
)

type DispatchProtocol string

const (
	HTTP DispatchProtocol = "http"
	AMQP DispatchProtocol = "amqp"
)

type Dispatcher struct {
	Client http.Client
	DispatchProtocol DispatchProtocol
}

type DispatchEvent struct {
	Action  blockatlas.TransactionType `json:"action"`
	Result  *blockatlas.Tx             `json:"result"`
	Webhook string                     `json:"webhook"`
}

func (d *Dispatcher) Run(events <-chan Event) {
	for event := range events {
		d.dispatch(event)
	}
}

func (d *Dispatcher) dispatch(event Event) {
	webhook := event.Subscription.Webhook

	action := DispatchEvent{
		Action:  event.Tx.Type,
		Result:  event.Tx,
		Webhook: webhook,
	}

	txJson, err := json.Marshal(action)
	if err != nil {
		logger.Panic(err)
	}

	logParams := logger.Params{
		"webhook": webhook,
		"coin":    event.Subscription.Coin,
		"txID":    event.Tx.ID,
	}
	switch d.DispatchProtocol {
	case HTTP:
		go d.postWebhook(webhook, txJson, logParams)
	case AMQP:
		go d.postMessageToQueue(webhook, txJson, logParams)
	default:
		logger.Fatal("DispatchProtocol is incorrect", logger.Params{"protocol": d.DispatchProtocol})
	}

	logger.Info("Dispatching messages...", logParams)
}

// use postWebhook if you want to transfer event by http protocol
func (d *Dispatcher) postWebhook(hook string, data []byte, logParams logger.Params) {
	_, err := d.Client.Post(hook, "application/json", bytes.NewReader(data))
	if err != nil {
		err = errors.E(err, "Failed to dispatch event", errors.Params{"webhook": hook}, logParams)
		logger.Error(err, logger.Params{"webhook": hook}, logParams)
	}
	logger.Info("Webhook dispatched", logger.Params{"webhook": hook}, logParams)
}

func (d *Dispatcher) postMessageToQueue(message string, rawMessage []byte, logParams logger.Params) {
	err := mq.Publish(rawMessage)
	if err != nil {
		err = errors.E(err, "Failed to dispatch event", errors.Params{"message": message}, logParams)
		logger.Error(err, logger.Params{"message": message}, logParams)
	}
	logger.Info("Message dispatched", logger.Params{"message": message}, logParams)
}
