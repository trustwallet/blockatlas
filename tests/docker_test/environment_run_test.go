// +build integration

package docker_test

import (
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/tests/docker_test/setup"
	"log"
	"os"
	"testing"
)

var (
	rawTransactionsChannel,
	transactionsChannel, subscriptionChannel mq.MessageChannel
)

func TestMain(m *testing.M) {
	setup.RunMQContainer()
	if err := mq.RawTransactions.Declare(); err != nil {
		log.Fatal(err)
	}
	if err := mq.Transactions.Declare(); err != nil {
		log.Fatal(err)
	}
	if err := mq.Subscriptions.Declare(); err != nil {
		log.Fatal(err)
	}
	rawTransactionsChannel = mq.RawTransactions.GetMessageChannel()
	subscriptionChannel = mq.Subscriptions.GetMessageChannel()
	transactionsChannel = mq.Transactions.GetMessageChannel()

	setup.RunRedisContainer()
	code := m.Run()
	setup.StopMQContainer()
	setup.StopRedisContainer()
	os.Exit(code)
}
