package observer

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Dispatcher struct {
	Client http.Client
}

func (d *Dispatcher) Run(events <-chan Event) {
	for event := range events {
		d.dispatch(event)
	}
}

func (d *Dispatcher) dispatch(event Event) {
	txJson, err := json.Marshal(&event.Tx)
	if err != nil {
		logrus.Panic(err)
	}

	webhook := event.Subscription.Webhook
	log := logrus.WithFields(logrus.Fields{
		"webhook": webhook,
		"coin": event.Subscription.Coin,
		"txID": event.Tx.ID,
	})

	_, err = d.Client.Post(webhook, "application/json", bytes.NewReader(txJson))
	if err != nil {
		log.WithError(err).Errorf("Failed to dispatch event %s: %s", webhook, err)
	}

	log.Debug("Dispatch")
}
