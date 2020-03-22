package mq

import (
	"context"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

var (
	PrefetchCount int
	amqpChan      *amqp.Channel
	conn          *amqp.Connection
)

type (
	Queue          string
	Consumer       func(amqp.Delivery, storage.Addresses)
	MessageChannel <-chan amqp.Delivery
)

const (
	Transactions         Queue = "transactions"
	Subscriptions        Queue = "subscriptions"
	RawTransactions      Queue = "rawTransactions"
	defaultPrefetchCount       = 5
	minPrefetchCount           = 1
)

func Init(uri string) (err error) {
	conn, err = amqp.Dial(uri)
	if err != nil {
		return
	}
	amqpChan, err = conn.Channel()
	if err != nil {
		return
	}
	return
}

func Close() {
	amqpChan.Close()
	conn.Close()
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

func (q Queue) RunConsumerForChannelWithCancel(consumer Consumer, messageChannel MessageChannel, cache storage.Addresses, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logger.Info("Consumer stopped")
			return
		case message := <-messageChannel:
			if message.Body == nil {
				continue
			}
			go consumer(message, cache)
		}
	}
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
		logger.Fatal("MQ issue " + err.Error())
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

func (q Queue) RunConsumerWithCancel(consumer Consumer, cache storage.Addresses, ctx context.Context) {
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
		logger.Fatal("MQ issue " + err.Error())
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

	for {
		select {
		case <-ctx.Done():
			logger.Info("Consumer stopped")
			return
		case message := <-messageChannel:
			if message.Body == nil {
				continue
			}
			go consumer(message, cache)
		}
	}
}

func (q Queue) GetMessageChannel() MessageChannel {
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
		logger.Fatal("MQ issue " + err.Error())
	}

	return messageChannel
}

func (mc MessageChannel) GetMessage() amqp.Delivery {
	return <-mc
}

func RestoreConnectionWorker(uri string, queue Queue, timeout time.Duration) {
	logger.Info("Run MQ RestoreConnectionWorker")
	for {
		if conn.IsClosed() {
			for {
				logger.Warn("MQ is not available now")
				logger.Warn("Trying to connect to MQ...")
				if err := Init(uri); err != nil {
					logger.Warn("MQ is still unavailable")
					time.Sleep(timeout)
					continue
				}
				if err := queue.Declare(); err != nil {
					logger.Warn("Can't declare queues:", queue)
					time.Sleep(timeout)
					continue
				} else {
					logger.Info("MQ connection restored")
					break
				}
			}
		}
		time.Sleep(timeout)
	}
}

func FatalWorker(timeout time.Duration) {
	logger.Info("Run FatalWorker")
	for {
		if conn.IsClosed() {
			logger.Fatal("MQ is not available now")
		}
		time.Sleep(timeout)
	}
}
