package mq

import (
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
)

var (
	PrefetchCount int
	amqpChan      *amqp.Channel
	Conn          *amqp.Connection
	queue         amqp.Queue
)

type (
	Queue    string
	Consumer func(amqp.Delivery, storage.Addresses)
)

const (
	Transactions         Queue = "transactions"
	Subscriptions        Queue = "subscriptions"
	defaultPrefetchCount       = 5
	minPrefetchCount           = 1
)

func Init(uri string) (err error) {
	Conn, err = amqp.Dial(uri)
	if err != nil {
		return
	}
	amqpChan, err = Conn.Channel()
	if err != nil {
		return
	}
	return
}

func Close() {
	amqpChan.Close()
	Conn.Close()
}

func (q Queue) Declare() error {
	_, err := amqpChan.QueueDeclare(string(q), true, false, false, false, nil)
	return err
}

func (q Queue) Publish(body []byte) error {
	return amqpChan.Publish("", string(q), false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
}

func (q Queue) RunConsumer(consumer Consumer, cache storage.Addresses) {
	messageChannel, err := amqpChan.Consume(
		string(q),
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Error(err)
		return
	}

	if PrefetchCount < minPrefetchCount {
		logger.Info("Change prefetch count to default")
		PrefetchCount = defaultPrefetchCount
	}

	err = amqpChan.Qos(
		PrefetchCount,
		0,
		true,
	)

	if err != nil {
		logger.Error("no qos limit ", err)
	}

	for data := range messageChannel {
		go consumer(data, cache)
	}
}
