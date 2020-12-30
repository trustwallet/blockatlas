// +build integration

package observer_test

import (
	"log"
	"os"
	"testing"

	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"github.com/trustwallet/golibs/network/mq"
)

var (
	rawTransactionsChannel, transactionsChannel, subscriptionChannel mq.MessageChannel
	database                                                         *db.Instance
)

func TestMain(m *testing.M) {
	database = setup.RunPgContainer()
	setup.RunMQContainer()
	if err := internal.RawTransactions.Declare(); err != nil {
		log.Fatal(err)
	}
	if err := internal.TxNotifications.Declare(); err != nil {
		log.Fatal(err)
	}
	if err := internal.Subscriptions.Declare(); err != nil {
		log.Fatal(err)
	}
	rawTransactionsChannel = internal.RawTransactions.GetMessageChannel()
	subscriptionChannel = internal.Subscriptions.GetMessageChannel()
	transactionsChannel = internal.TxNotifications.GetMessageChannel()

	code := m.Run()

	setup.StopMQContainer()
	setup.StopPgContainer()
	os.Exit(code)
}
