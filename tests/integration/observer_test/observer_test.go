// +build integration

package observer_test

import (
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/servicerepo"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"log"
	"os"
	"testing"
)

var (
	database                                                         *db.Instance
)

var serviceRepo *servicerepo.ServiceRepo;

func TestMain(m *testing.M) {
	serviceRepo = servicerepo.New()
	mq.InitService(serviceRepo)
	setup.RunMQContainer(serviceRepo)
	mqService := mq.GetService(serviceRepo)

	database = setup.RunPgContainer()
	setup.RunMQContainer(serviceRepo)
	if err := mqService.RawTransactions().Declare(); err != nil {
		log.Fatal(err)
	}
	if err := mqService.TxNotifications().Declare(); err != nil {
		log.Fatal(err)
	}
	if err := mqService.Subscriptions().Declare(); err != nil {
		log.Fatal(err)
	}

	code := m.Run()

	setup.StopMQContainer()
	setup.StopPgContainer()
	os.Exit(code)
}
