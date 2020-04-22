package setup

import (
	"fmt"
	"github.com/ory/dockertest"
	"github.com/trustwallet/blockatlas/mq"
	"log"
)

var (
	mqResource *dockertest.Resource
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
		return mq.GetService().Init(fmt.Sprintf("amqp://localhost:%s", mqResource.GetPort("5672/tcp")), 500)
	}); err != nil {
		return err
	}
	return nil
}

func stopMQContainer() error {
	mq.GetService().Close()
	return mqResource.Close()
}
