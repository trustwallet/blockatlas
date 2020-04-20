package mq

import (
	"context"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/servicerepo"
	"time"
)

type MQServiceIface interface {
	Init(uri string, prefetchCount int) (err error)
	Close()
	RestoreConnectionWorker(uri string, queue *Queue, timeout time.Duration)
	FatalWorker(timeout time.Duration)

	TxNotifications() *Queue
	Subscriptions() *Queue
	RawTransactions() *Queue
}

type mqService struct {
	prefetchCount   int
	amqpChan        *amqp.Channel
	conn            *amqp.Connection
	txNotifications *Queue
	subscriptions   *Queue
	rawTransactions *Queue
}

// InitService Adds new mq.mqService instance
func InitService(serviceRepo *servicerepo.ServiceRepo) {
	serviceRepo.Add(new(mqService))
}

func GetService(s *servicerepo.ServiceRepo) MQServiceIface {
	return s.Get("mq.mqService").(MQServiceIface)
}

type Queue struct {
	name          string
	channel       *amqp.Channel
	prefetchCount int
}

func NewQueue(name string, channel *amqp.Channel, prefetchCount int) *Queue {
	return &Queue{name: name, channel: channel, prefetchCount: prefetchCount}
}

type (
	Consumer           func(amqp.Delivery)
	ConsumerWithDbConn func(*db.Instance, amqp.Delivery)
	MessageChannel     <-chan amqp.Delivery
)

func (m *mqService) Init(uri string, prefetchCount int) (err error) {
	m.conn, err = amqp.Dial(uri)
	if err != nil {
		return err
	}
	m.amqpChan, err = m.conn.Channel()
	if err != nil {
		return err
	}
	m.prefetchCount = prefetchCount

	m.txNotifications = NewQueue("txNotifications", m.amqpChan, m.prefetchCount)
	m.subscriptions = NewQueue("subscriptions", m.amqpChan, m.prefetchCount)
	m.rawTransactions = NewQueue("rawTransactions", m.amqpChan, m.prefetchCount)

	return nil
}

func (m *mqService) TxNotifications() *Queue { return m.txNotifications }

func (m *mqService) Subscriptions() *Queue { return m.subscriptions }

func (m *mqService) RawTransactions() *Queue { return m.rawTransactions }

func (m *mqService) Close() {
	err := m.amqpChan.Close()
	if err != nil {
		logger.Error(err)
	}

	err = m.conn.Close()
	if err != nil {
		logger.Error(err)
	}
}

func (mc MessageChannel) GetMessage() amqp.Delivery {
	return <-mc
}

func (q Queue) Declare() error {
	_, err := q.channel.QueueDeclare(q.name, true, false, false, false, nil)
	return err
}

func (q Queue) Publish(body []byte) error {
	return q.channel.Publish("", q.name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
}

func RunConsumerForChannelWithCancelAndDbConn(consumer ConsumerWithDbConn, messageChannel MessageChannel, database *db.Instance, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logger.Info("Consumer stopped")
			return
		case message := <-messageChannel:
			if message.Body == nil {
				continue
			}
			go consumer(database, message)
		}
	}
}

func (q Queue) GetMessageChannel() MessageChannel {
	messageChannel, err := q.channel.Consume(
		q.name,
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

	err = q.channel.Qos(
		q.prefetchCount,
		0,
		true,
	)
	if err != nil {
		logger.Fatal("No qos limit ", err)
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
			logger.Info("Consumer stopped")
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
			logger.Info("Consumer stopped")
			return
		case message := <-messageChannel:
			if message.Body == nil {
				continue
			}
			go consumer(database, message)
		}
	}
}

func (m *mqService) RestoreConnectionWorker(uri string, queue *Queue, timeout time.Duration) {
	logger.Info("Run MQ RestoreConnectionWorker")
	for {
		if m.conn.IsClosed() {
			for {
				logger.Warn("MQ is not available now")
				logger.Warn("Trying to connect to MQ...")
				if err := m.Init(uri, m.prefetchCount); err != nil {
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

func (m *mqService) FatalWorker(timeout time.Duration) {
	logger.Info("Run MQ FatalWorker")
	for {
		if m.conn.IsClosed() {
			logger.Fatal("MQ is not available now")
		}
		time.Sleep(timeout)
	}
}
