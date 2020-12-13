package observer_test

import (
	"context"
	"fmt"
	"github.com/ory/dockertest"
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pubsub"
	"github.com/trustwallet/blockatlas/pubsub/mqclient"
	"os"
	"testing"
)

var (
	mqResource *dockertest.Resource
	mqClient   pubsub.Client
	ctx        context.Context
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	if runMQContainer() != nil {
		log.Fatal("container doesn't start")
	}

	code := m.Run()

	log.Error(stopMQContainer())
	log.Error(stopMQContainer())
	os.Exit(code)
}

func TestMqConnect(t *testing.T) {

}

func runMQContainer() error {
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	mqResource, err = pool.Run("rabbitmq", "latest", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		uri := fmt.Sprintf("amqp://localhost:%s", mqResource.GetPort("5672/tcp"))
		mqClient = mqclient.New(uri, 10, ctx)
		return mqClient.Connect()
	}); err != nil {
		return err
	}
	return nil
}

func stopMQContainer() error {
	mq.Close()
	return mqResource.Close()
}
