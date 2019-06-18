package observer

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/observer/storage"
	"net/http"
)

type Dispatcher struct {
	Client *http.Client
}

func (d Dispatcher) DispatchTransactions(txs []models.Tx) {
	s := storage.GetInstance()
	for _, tx := range txs {
		logrus.Debugf("Coin: %d, Tx: %s, From: %s, To: %s", tx.Coin, tx.Id, tx.From, tx.To)
		if s.Contains(tx.Coin, tx.To) {
			go d.dispatch(s.Get(tx.Coin, tx.To), tx)
		}
	}
}

func (d Dispatcher) dispatch(ob models.Observer, tx models.Tx) {
	txJson, jsonErr := json.Marshal(&tx)
	if jsonErr != nil {
		logrus.WithError(jsonErr).Errorf("Failed to convert Tx to JSON: %s", jsonErr)
		return
	}

	body := bytes.NewBuffer(txJson)
	_, postError := d.Client.Post(ob.Webhook, "application/json", body)
	if postError != nil {
		logrus.WithError(postError).Errorf("Failed to call webhook %s: %s", ob.Webhook, postError)
	}
}
