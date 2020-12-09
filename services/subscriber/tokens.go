package subscriber

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
)

const Tokens Subscriber = "tokens"

func RunTokensSubscriber(database *db.Instance, delivery amqp.Delivery) {
	event := make(map[string][]models.Asset)
	if err := json.Unmarshal(delivery.Body, &event); err != nil {
		if err := delivery.Ack(false); err != nil {
			log.WithFields(log.Fields{"service": Tokens}).Error(err)
		}
	}

	for address, assets := range event {
		if err := database.AddAssociationsForAddress(address, assets); err != nil {
			log.Error("Failed to AddAssociationsForAddress: " + err.Error())
		}
	}
	log.WithFields(log.Fields{"service": Tokens, "count": len(event)}).Info("Subscribed")
	if err := delivery.Ack(false); err != nil {
		log.WithFields(log.Fields{"service": Tokens}).Error(err)
	}
	log.Info("------------------------------------------------------------")
}
