// +build integration

package observer_test

import (
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"context"
	"log"
	"os"
	"testing"
)

var (
	rawTransactionsChannel, transactionsChannel, subscriptionChannel mq.MessageChannel
	database                                                         *db.Instance
)

func RunConsumerForChannelWithCancelAndDbConn(consumer mq.ConsumerWithDbConn, messageChannel mq.MessageChannel, database *db.Instance, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logger.Info("Consumer stopped")
			return
		case message := <-messageChannel:
			if message.Body == nil {
				continue
			}
			go consumer(database, message)
		}
	}
}

func TestMain(m *testing.M) {
	mq.InitService()
	setup.RunMQContainer()
	mqService := mq.GetService()

	database = setup.RunPgContainer()
	setup.RunMQContainer()
	if err := mqService.RawTransactions().Declare(); err != nil {
		log.Fatal(err)
	}
	if err := mqService.TxNotifications().Declare(); err != nil {
		log.Fatal(err)
	}
	if err := mqService.Subscriptions().Declare(); err != nil {
		log.Fatal(err)
	}
	rawTransactionsChannel = mqService.RawTransactions().GetMessageChannel()
	subscriptionChannel = mqService.Subscriptions().GetMessageChannel()
	transactionsChannel = mqService.TxNotifications().GetMessageChannel()

	code := m.Run()

	setup.StopMQContainer()
	setup.StopPgContainer()
	os.Exit(code)
}
