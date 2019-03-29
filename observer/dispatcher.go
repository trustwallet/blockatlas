package observer

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/observer/storage"
	"net/http"
)

type Dispatcher struct {
	Client *http.Client
}

func (d Dispatcher) NotifyObservers(txs []models.Tx) {
	s := storage.GetInstance()
	for _, tx := range txs {
		logrus.Debugf("Tx: %s, From: %s, To: %s", tx.Id, tx.From, tx.To)
		if s.Contains(tx.Coin, tx.To) {
			go d.notify(s.Get(tx.Coin, tx.To), tx)
		}
	}
}

func (d Dispatcher) notify(ob models.Observer, tx models.Tx) {
	txJson, jsonErr := json.Marshal(&tx)
	if jsonErr != nil {
		logrus.WithError(jsonErr).Errorf("Failed to marshal json: %s", jsonErr)
		return
	}

	logrus.Info("POST %s - JSON: %s", ob.Webhook, txJson)

	/*
	body := bytes.NewBuffer(txJson)
	_, postError := d.Client.Post(ob.Webhook, "application/json", body)
	if postError != nil {
		logrus.WithError(postError).Errorf("Failed to call webhook %s: %s", ob.Webhook, postError)
	}
	*/
}
