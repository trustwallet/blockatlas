package subscriber

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"go.elastic.co/apm"
)

const Tokens Subscriber = "tokens"

func RunTokensSubscriber(database *db.Instance, delivery amqp.Delivery) {
	tx := apm.DefaultTracer.StartTransaction("RunTokensSubscriber", "app")
	defer tx.End()

	ctx := apm.ContextWithTransaction(context.Background(), tx)

	event := make(map[string][]string)
	if err := json.Unmarshal(delivery.Body, &event); err != nil {
		if err := delivery.Ack(false); err != nil {
			logger.Fatal(err, err)
		}
	}

	for address, assets := range event {
		if err := database.AddAssociationsForAddress(address, assets, ctx); err != nil {
			logger.Error("Failed to AddAssociationsForAddress: " + err.Error())
		}
	}

	if err := delivery.Ack(false); err != nil {
		logger.Fatal(err, err)
	}
}
