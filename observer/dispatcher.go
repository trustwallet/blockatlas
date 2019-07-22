package observer

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"net/http"
)

type Dispatcher struct {
	Client http.Client
}

type DispatchEvent struct {
	Action string         `json:"action"`
	Result *blockatlas.Tx `json:"result"`
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
	txJson, err := json.Marshal(action)
	if err != nil {
		logrus.Panic(err)
	}

	webhooks := event.Subscription.Webhooks
	log := logrus.WithFields(logrus.Fields{
		"webhook": webhooks,
		"coin":    event.Subscription.Coin,
		"txID":    event.Tx.ID,
	})

	for _, hook := range webhooks {
		_, err = d.Client.Post(hook, "application/json", bytes.NewReader(txJson))
		if err != nil {
			log.WithError(err).Errorf("Failed to dispatch event %s: %s", hook, err)
		}

		log.Debug("Dispatch")
	}
}
