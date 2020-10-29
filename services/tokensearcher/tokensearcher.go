package tokensearcher

import (
	"context"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/notifier"
	"go.elastic.co/apm"
	"strconv"
)

const TokenSearcher = "TokenSearcher"

func Run(database *db.Instance, delivery amqp.Delivery) {
	tx := apm.DefaultTracer.StartTransaction("RunNotifier", "app")
	defer tx.End()
	ctx := apm.ContextWithTransaction(context.Background(), tx)

	txs, err := notifier.GetTransactionsFromDelivery(delivery, TokenSearcher, ctx)
	if err != nil {
		logger.Error("failed to get transactions", err)
		if err := delivery.Ack(false); err != nil {
			logger.Error(err)
		}
	}
	if len(txs) == 0 {
		return
	}
	coinID := strconv.Itoa(int(txs[0].Coin))
	var addresses []string
	for _, tx := range txs {
		addresses = append(addresses, tx.GetAddresses()...)
	}
	for i := range addresses {
		addresses[i] = coinID + "_" + addresses[i]
	}

	associationsFromTransactions, err := database.GetAssociationsByAddresses(notifier.ToUniqueAddresses(addresses), ctx)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("AssociationsFromTransactions "+strconv.Itoa(len(associationsFromTransactions)), logger.Params{"service": TokenSearcher})

	associationsToAdd := associationsToAdd(fromModelToAssociation(associationsFromTransactions), assetsMap(txs, coinID))

	logger.Info("AssociationsToAdd "+strconv.Itoa(len(associationsToAdd)), logger.Params{"service": TokenSearcher})
	err = database.UpdateAssociationsForExistingAddresses(associationsToAdd, ctx)
	if err != nil {
		logger.Error(err)
		return
	}

	if err := delivery.Ack(false); err != nil {
		logger.Error(err)
	}
	logger.Info("------------------------------------------------------------")
}
