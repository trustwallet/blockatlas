package subscriber

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"go.elastic.co/apm"
	"strconv"
)

const Tokens Subscriber = "tokens"

func RunTokensSubscriber(database *db.Instance, delivery amqp.Delivery) {
	tx := apm.DefaultTracer.StartTransaction("RunTokensSubscriber", "app")
	defer tx.End()

	ctx := apm.ContextWithTransaction(context.Background(), tx)
	event := make(map[string][]models.Asset)
	if err := json.Unmarshal(delivery.Body, &event); err != nil {
		if err := delivery.Ack(false); err != nil {
			log.Fatal(err, err)
		}
	}

	for address, assets := range event {
		if err := database.AddAssociationsForAddress(address, assets, ctx); err != nil {
			log.Error("Failed to AddAssociationsForAddress: " + err.Error())
		}
	}
	log.WithFields(log.Fields{"service": Tokens}).Info("Subscribed " + strconv.Itoa(len(event)))
	if err := delivery.Ack(false); err != nil {
		log.Fatal(err, err)
	}
	log.Info("------------------------------------------------------------")
}
