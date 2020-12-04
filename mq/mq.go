package mq

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
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
	// End consumer of published transactions. Not consumed on blockatlas
	TxNotifications Queue = "txNotifications"
	// Address:coin subscriptions
	Subscriptions Queue = "subscriptions"
	// Transactions to process, if match subscriptions, pushed to TxNotifications
	RawTransactions Queue = "rawTransactions"
	// Token indexer for finding asset association with an address
	RawTransactionsSearcher Queue = "rawTransactionsSearcher"
	// Token indexer for finding new assets
	RawTransactionsTokenIndexer Queue = "rawTransactionsTokenIndexer"
	// Register new addresses to observers for token transfers
	TokensRegistration Queue = "tokensRegistration"
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
		log.Fatal("GetMessageChannel MQ issue "+err.Error(), string(q))
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

func FatalWorker(timeout time.Duration) {
	log.Info("Run MQ FatalWorker")
	for {
		if conn.IsClosed() {
			log.Fatal("MQ is not available now")
		}
		time.Sleep(timeout)
	}
}
