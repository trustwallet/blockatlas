package tokensearcher

import (
	"context"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/observer/notifier"
	"go.elastic.co/apm"
)

func Run(database *db.Instance, delivery amqp.Delivery) {
	tx := apm.DefaultTracer.StartTransaction("RunNotifier", "app")
	defer tx.End()
	ctx := apm.ContextWithTransaction(context.Background(), tx)

	txs, err := notifier.GetTransactionsFromDelivery(delivery, ctx)
	if err != nil {
		logger.Error("failed to get transactions", err)
		if err := delivery.Ack(false); err != nil {
			logger.Error(err)
		}
	}

	addresses := make([]string, 0)
	for _, tx := range txs {
		addresses = append(addresses, tx.GetAddresses()...)
	}

	associationsFromTransactions, err := database.GetAssociationsByAddresses(notifier.ToUniqueAddresses(addresses), ctx)
	if err != nil {
		logger.Error(err)
		return
	}

	associationsToAdd := associationsToAdd(fromModelToAssociation(associationsFromTransactions), assetsMap(txs))
	err = database.UpdateAssociationsForExistingAddresses(associationsToAdd, ctx)
	if err != nil {
		logger.Error(err)
		return
	}

	if err := delivery.Ack(false); err != nil {
		logger.Error(err)
	}
}
