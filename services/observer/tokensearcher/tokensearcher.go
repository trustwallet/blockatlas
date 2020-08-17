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

	defer func() {
		if err := delivery.Ack(false); err != nil {
			logger.Error(err)
		}
	}()

	txs, err := notifier.GetTransactionsFromDelivery(delivery, ctx)
	if err != nil {
		logger.Error("failed to get transactions", err)
	}
	coin := txs[0].Coin

	allAddresses := make([]string, 0)
	for _, tx := range txs {
		allAddresses = append(allAddresses, tx.GetAddresses()...)
	}

	addresses := notifier.ToUniqueAddresses(allAddresses)
	associationsCurrent, err := database.GetTokensByAddresses(addresses, ctx)
	txsMap := makeTransactionsMap(coin, txs)
	associationsToAdd := allAssociationsToAdd(database, associationsCurrent, txsMap)

	err = database.AddAssociations(associationsToAdd, ctx)
	if err != nil {
		logger.Error(err)
	}
}
