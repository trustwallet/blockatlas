package tokensearcher

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/services/notifier"
	"strconv"
)

const TokenSearcher = "TokenSearcher"

type TokenSearcherConsumer struct {
	Database *db.Instance
}

func (c *TokenSearcherConsumer) GetQueue() string {
	return string(new_mq.RawTransactionsSearcher)
}

func (c *TokenSearcherConsumer) Callback(msg amqp.Delivery) error {
	txs, err := notifier.GetTransactionsFromDelivery(msg, TokenSearcher)
	if err != nil {
		log.Error("failed to get transactions", err)
		return err
	}
	if len(txs) == 0 {
		return nil
	}
	coinID := strconv.Itoa(int(txs[0].Coin))
	var addresses []string
	for _, tx := range txs {
		addresses = append(addresses, tx.GetAddresses()...)
	}
	for i := range addresses {
		addresses[i] = coinID + "_" + addresses[i]
	}

	associationsFromTransactions, err := c.Database.GetAssociationsByAddresses(notifier.ToUniqueAddresses(addresses))
	if err != nil {
		log.Error(err)
		return err
	}
	log.WithFields(log.Fields{"service": TokenSearcher}).
		Info("AssociationsFromTransactions " + strconv.Itoa(len(associationsFromTransactions)))

	associationsToAdd := associationsToAdd(fromModelToAssociation(associationsFromTransactions), assetsMap(txs, coinID))

	log.WithFields(log.Fields{"service": TokenSearcher}).
		Info("AssociationsToAdd " + strconv.Itoa(len(associationsToAdd)))
	err = c.Database.UpdateAssociationsForExistingAddresses(associationsToAdd)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("------------------------------------------------------------")
	return nil
}
