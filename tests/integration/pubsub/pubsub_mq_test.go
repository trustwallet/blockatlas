// +build integration

package pubsub_test

import (
	"context"
	"fmt"
	"github.com/ory/dockertest"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pubsub"
	"github.com/trustwallet/blockatlas/pubsub/mqclient"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	mqResource *dockertest.Resource
	mqClient   pubsub.Client
	ctx        context.Context
)

type RawTransactionConsumer struct {
	TestFunc func(delivery amqp.Delivery)
}

func (c RawTransactionConsumer) GetQueue() string {
	return string("test_queue")
}

func (c RawTransactionConsumer) Callback(msg amqp.Delivery) error {
	c.TestFunc(msg)
	return nil
}

func TestMain(m *testing.M) {
	ctx = context.Background()
	if err := runMQContainer(); err != nil {
		log.Fatal("container doesn't start: ", err)
	}

	code := m.Run()

	log.Error(stopMQContainer())
	os.Exit(code)
}

func TestMqConnect(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	var consumer pubsub.Consumer
	consumer = RawTransactionConsumer{
		TestFunc: func(msg amqp.Delivery) {
			defer wg.Done()
			assert.Equal(t, string(msg.Body), `{"message":"test"}`)
		},
	}
	assert.Nil(t, mqClient.AddStream(&consumer, false))
	time.Sleep(10 * time.Second)
	assert.Nil(t, mqClient.Push(consumer.GetQueue(), []byte(`{"message":"test"}`)))
	wg.Wait()
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
	//time.Sleep(10 * time.Second)
	if err = pool.Retry(func() error {
		uri := fmt.Sprintf("amqp://localhost:%s", mqResource.GetPort("5672/tcp"))
		mqClient = mqclient.New(uri, 10, ctx)
		return mqClient.Connect()
	}); err != nil {
		stopMQContainer()
		return err
	}

	return mqClient.Run()
}

func stopMQContainer() error {
	mqClient.Close()
	return mqResource.Close()
}
