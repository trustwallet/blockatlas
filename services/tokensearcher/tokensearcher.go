package tokensearcher

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
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
		log.Error("failed to get transactions", err)
		if err := delivery.Ack(false); err != nil {
			log.Error(err)
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
		log.Error(err)
		return
	}
	log.WithFields(log.Fields{"service": TokenSearcher}).
		Info("AssociationsFromTransactions " + strconv.Itoa(len(associationsFromTransactions)))

	associationsToAdd := associationsToAdd(fromModelToAssociation(associationsFromTransactions), assetsMap(txs, coinID))

	log.WithFields(log.Fields{"service": TokenSearcher}).
		Info("AssociationsToAdd " + strconv.Itoa(len(associationsToAdd)))
	err = database.UpdateAssociationsForExistingAddresses(associationsToAdd, ctx)
	if err != nil {
		log.Error(err)
		return
	}

	if err := delivery.Ack(false); err != nil {
		log.Error(err)
	}
	log.Info("------------------------------------------------------------")
}
