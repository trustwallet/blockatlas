package subscriber

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/mq"
	"go.elastic.co/apm"
)

const Tokens Subscriber = "tokens"

type TokenSubscriberConsumer struct {
	Database *db.Instance
}

func (c *TokenSubscriberConsumer) GetQueue() string {
	return string(new_mq.TokensRegistration)
}

func (c *TokenSubscriberConsumer) Callback(msg amqp.Delivery) error {
	tx := apm.DefaultTracer.StartTransaction("RunTokensSubscriber", "app")
	defer tx.End()

	ctx := apm.ContextWithTransaction(context.Background(), tx)
	event := make(map[string][]models.Asset)
	if err := json.Unmarshal(msg.Body, &event); err != nil {
		return err
	}

	for address, assets := range event {
		if err := c.Database.AddAssociationsForAddress(address, assets, ctx); err != nil {
			log.Error("Failed to AddAssociationsForAddress: " + err.Error())
			return err
		}
	}
	log.WithFields(log.Fields{"service": Tokens, "count": len(event)}).Info("Subscribed")
	log.Info("------------------------------------------------------------")
	return nil
}
