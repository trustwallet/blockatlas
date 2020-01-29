package observer

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

type Dispatcher struct {
}

type DispatchEvent struct {
	Action blockatlas.TransactionType `json:"action"`
	Result *blockatlas.Tx             `json:"result"`
}

func (d *Dispatcher) Run(events <-chan Event) {
	for event := range events {
		d.dispatch(event)
	}
}

func (d *Dispatcher) dispatch(event Event) {
	action := DispatchEvent{
		Action: event.Tx.Type,
		Result: event.Tx,
	}
	webhook := event.Subscription.Webhook
	logParams := logger.Params{
		"webhook": webhook,
		"coin":    event.Subscription.Coin,
		"txID":    event.Tx.ID,
	}
	go d.postWebhook(webhook, action, logParams)
	logger.Info("Dispatching webhooks...", logger.Params{"webhook": webhook}, logParams)
}

func (d *Dispatcher) postWebhook(hook string, data interface{}, logParams logger.Params) {
	client := blockatlas.InitClient(hook)
	client.Headers["Content-Type"] = "application/json"
	err := client.Post(nil, "", data)
	if err != nil {
		err = errors.E(err, errors.Params{"hook": hook}).PushToSentry()
		logger.Error(err, "Failed to dispatch event", logger.Params{"webhook": hook}, logParams)
	}
	logger.Info("Webhook dispatched", logger.Params{"webhook": hook}, logParams)
}
