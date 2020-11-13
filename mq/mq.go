package mq

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"time"
)

var (
	PrefetchCount int
	amqpChan      *amqp.Channel
	conn          *amqp.Connection
)

type (
	Queue              string
	Consumer           func(amqp.Delivery)
	ConsumerWithDbConn func(*db.Instance, amqp.Delivery)
	MessageChannel     <-chan amqp.Delivery
)

const (
	TxNotifications             Queue = "txNotifications"
	Subscriptions               Queue = "subscriptions"
	RawTransactions             Queue = "rawTransactions"
	RawTransactionsSearcher     Queue = "rawTransactionsSearcher"
	RawTransactionsTokenIndexer Queue = "rawTransactionsTokenIndexer"
	TokensRegistration          Queue = "tokensRegistration"
)

func Init(uri string) (err error) {
	conn, err = amqp.Dial(uri)
	if err != nil {
		return err
	}
	amqpChan, err = conn.Channel()
	return err
}

func Close() {
	err := amqpChan.Close()
	if err != nil {
		log.Error(err)
	}

	err = conn.Close()
	if err != nil {
		log.Error(err)
	}
}

func (mc MessageChannel) GetMessage() amqp.Delivery {
	return <-mc
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

func RunConsumerForChannelWithCancelAndDbConn(consumer ConsumerWithDbConn, messageChannel MessageChannel, database *db.Instance, concurrent bool, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Info("Consumer stopped")
			return
		case message := <-messageChannel:
			if message.Body == nil {
				continue
			}
			if concurrent {
				go consumer(database, message)
			} else {
				consumer(database, message)
			}

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
		log.Fatal("MQ issue " + err.Error())
	}

	err = amqpChan.Qos(
		PrefetchCount,
		0,
		true,
	)
	if err != nil {
		log.Fatal("No qos limit ", err)
	}

	return messageChannel
}

func (q Queue) RunConsumer(consumer Consumer) {
	messageChannel := q.GetMessageChannel()
	for data := range messageChannel {
		go consumer(data)
	}
}

func (q Queue) RunConsumerWithCancel(consumer Consumer, ctx context.Context) {
	messageChannel := q.GetMessageChannel()
	for {
		select {
		case <-ctx.Done():
			log.Info("Consumer stopped")
			return
		case message := <-messageChannel:
			if message.Body == nil {
				continue
			}
			go consumer(message)
		}
	}
}

func (q Queue) RunConsumerWithCancelAndDbConn(consumer ConsumerWithDbConn, database *db.Instance, ctx context.Context) {
	messageChannel := q.GetMessageChannel()
	for {
		select {
		case <-ctx.Done():
			log.Info("Consumer stopped")
			return
		case message := <-messageChannel:
			if message.Body == nil {
				continue
			}
			consumer(database, message)
		}
	}
}

func (q Queue) RunConsumerWithCancelAndDbConnConcurrent(consumer ConsumerWithDbConn, database *db.Instance, ctx context.Context) {
	messageChannel := q.GetMessageChannel()
	for {
		select {
		case <-ctx.Done():
			log.Info("Consumer stopped")
			return
		case message := <-messageChannel:
			if message.Body == nil {
				continue
			}
			go consumer(database, message)
		}
	}
}

func FatalWorker(timeout time.Duration) {
	log.Info("Run MQ FatalWorker")
	for {
		if conn.IsClosed() {
			log.Fatal("MQ is not available now")
		}
		time.Sleep(timeout)
	}
}
