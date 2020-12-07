package setup

import (
	"fmt"
	"github.com/ory/dockertest"
	"github.com/trustwallet/blockatlas/mq"
	"log"
)

var (
	mqResource *dockertest.Resource
	mqClient   *new_mq.Client
)

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
		mqClient, err = new_mq.New(fmt.Sprintf("amqp://localhost:%s", mqResource.GetPort("5672/tcp")), 1, nil)
		return err
	}); err != nil {
		return err
	}
	return nil
}

func stopMQContainer() error {
	mqClient.Close()
	return mqResource.Close()
}
