// +build integration

package observer_test

import (
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"log"
	"os"
	"testing"
)

var (
	rawTransactionsChannel, transactionsChannel, subscriptionChannel mq.MessageChannel
	database                                                         *db.Instance
)

func TestMain(m *testing.M) {
	database = setup.RunPgContainer()
	setup.RunMQContainer()
	if err := mq.RawTransactions.Declare(); err != nil {
		log.Fatal(err)
	}
	if err := mq.TxNotifications.Declare(); err != nil {
		log.Fatal(err)
	}
	if err := mq.Subscriptions.Declare(); err != nil {
		log.Fatal(err)
	}
	rawTransactionsChannel = mq.RawTransactions.GetMessageChannel()
	subscriptionChannel = mq.Subscriptions.GetMessageChannel()
	transactionsChannel = mq.TxNotifications.GetMessageChannel()

	code := m.Run()

	setup.StopMQContainer()
	setup.StopPgContainer()
	os.Exit(code)
}
