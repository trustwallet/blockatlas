package observer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
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
		logger.Panic(err)
	}

	webhooks := event.Subscription.Webhooks
	log := logrus.WithFields(logrus.Fields{
		"webhook": webhooks,
		"coin":    event.Subscription.Coin,
		"txID":    event.Tx.ID,
	})
	for _, hook := range webhooks {
		go d.postWebhook(hook, txJson, log)
	}
	log.Infoln(fmt.Sprintf("Dispatch: hook count = %d", len(webhooks)))
}

func (d *Dispatcher) postWebhook(hook string, data []byte, log *logrus.Entry) {
	_, err := d.Client.Post(hook, "application/json", bytes.NewReader(data))
	if err != nil {
		err = errors.E(err, errors.Params{"hook": hook})
		log.WithError(err).Errorf("Failed to dispatch event %s", hook)
	}
	log.Info("Dispatch: hook: ", hook)
}
