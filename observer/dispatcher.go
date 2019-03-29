package observer

import (
	"bytes"
	"github.com/gin-gonic/gin/json"
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
		if s.Contains(tx.Coin, tx.To) {
			go d.notify(s.Get(tx.Coin, tx.To), tx)
		}
	}
}

func (d Dispatcher) notify(ob models.Observer, tx models.Tx) {
	if ob.Address != tx.To {
		return
	}

	txJson, jsonErr := json.Marshal(&tx)
	if jsonErr != nil {
		logrus.WithError(jsonErr).Errorf("Failed to marschal json: %s", jsonErr)
		return
	}

	body := bytes.NewBuffer(txJson)
	_, postError := d.Client.Post(ob.Webhook, "application/json", body)
	if postError != nil {
		logrus.WithError(postError).Errorf("Failed to call webhook %s: %s", ob.Webhook, postError)
	}
}
