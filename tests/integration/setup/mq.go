package setup

import (
	"fmt"
	"github.com/ory/dockertest"
	"github.com/trustwallet/blockatlas/mq"
	"github.com/trustwallet/blockatlas/pkg/servicerepo"
	"log"
)

var (
	mqResource *dockertest.Resource
	mqService  mq.MQServiceIface
)

func runMQContainer(serviceRepo *servicerepo.ServiceRepo) error {
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	mqResource, err = pool.Run("rabbitmq", "latest", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	mqService = mq.GetService(serviceRepo)
	if err = pool.Retry(func() error {
		return mqService.Init(fmt.Sprintf("amqp://localhost:%s", mqResource.GetPort("5672/tcp")), 500)
	}); err != nil {
		return err
	}
	return nil
}

func stopMQContainer() error {
	mqService.Close()
	return mqResource.Close()
}
