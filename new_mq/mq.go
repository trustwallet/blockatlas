package new_mq

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

type (
	Queue string
)

const (
	reconnectDelay = 5 * time.Second
	resendDelay    = 5 * time.Second
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

var (
	ErrDisconnected = errors.New("disconnected from rabbitmq, trying to reconnect")
)

type Client struct {
	conn          *amqp.Connection
	channel       *amqp.Channel
	prefetchCount int
	streams       []*Stream
	ctx           context.Context
	notifyClose   chan *amqp.Error
	notifyConfirm chan amqp.Confirmation
	isConnected   bool
	alive         bool
}

type Stream struct {
	consumer    Consumer
	client      *Client
	isConnected bool
	isStream    bool
}

type Consumer interface {
	GetQueue() string
	Callback(msg amqp.Delivery) error
}

func New(uri string, prefetchCount int, ctx context.Context) (client *Client, err error) {
	client = &Client{
		ctx:           ctx,
		alive:         true,
		prefetchCount: prefetchCount,
	}
	go client.handleReconnect(uri)
	return client, err
}

func (c *Client) handleReconnect(addr string) {
	for c.alive {
		c.isConnected = false
		var retryCount int
		for !c.connect(addr) {
			if !c.alive {
				return
			}
			select {
			case <-c.ctx.Done():
				return
			case <-time.After(reconnectDelay + /*time.Duration(retryCount)**/ time.Second):
				// Add metric
				retryCount++
			}
		}
		select {
		case <-c.ctx.Done():
			return
		case <-c.notifyClose:
		}
	}
}

func (c *Client) connect(addr string) bool {
	conn, err := amqp.Dial(addr)
	if err != nil {
		log.Error("Client.connect MQ Dial issue " + err.Error())
		return false
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Error("Client.connect MQ Channel issue " + err.Error())
		return false
	}
	err = ch.Confirm(false)
	if err != nil {
		log.Info("Client.connect MQ Channel issue " + err.Error())
	}
	c.changeConnection(conn, ch)
	c.isConnected = true
	return true
}

func (c *Client) changeConnection(connection *amqp.Connection, channel *amqp.Channel) {
	c.conn = connection
	c.channel = channel
	c.notifyClose = make(chan *amqp.Error)
	c.notifyConfirm = make(chan amqp.Confirmation)
	c.channel.NotifyClose(c.notifyClose)
	c.channel.NotifyPublish(c.notifyConfirm)
	c.run()
}

func (c *Client) AddStream(consumer Consumer) {
	stream := c.consume(&consumer, true)
	go stream.Connect(c.ctx)
	c.streams = append(c.streams, stream)
}

func (c *Client) AddPublish(consumer Consumer) {
	stream := c.consume(&consumer, false)
	go stream.Connect(c.ctx)
	c.streams = append(c.streams, stream)
}

func (c *Client) DeclareQueue(consumer Consumer) {
	_, err := c.channel.QueueDeclare(consumer.GetQueue(), true, false, false, false, nil)
	if err != nil {
		log.Fatal("Stream.Init MQ issue "+err.Error(), consumer.GetQueue())
	}
}

func (c *Client) Close() {
	err := c.channel.Close()
	if err != nil {
		log.Error("Client.Close MQ issue " + err.Error())
	}

	err = c.conn.Close()
	if err != nil {
		log.Error("Client.Close MQ issue " + err.Error())
	}
}

func (c *Client) consume(consumer *Consumer, isStream bool) *Stream {
	return &Stream{
		consumer:    *consumer,
		client:      c,
		isConnected: false,
		isStream:    isStream,
	}
}

func (c *Client) run() {
	err := c.channel.Qos(c.prefetchCount, 0, false)
	if err != nil {
		log.Fatal("Client.connect MQ Qos issue " + err.Error())
	}
	for _, stream := range c.streams {
		if !stream.isConnected {
			go stream.Connect(c.ctx)
		}
	}
}

func (c *Client) Push(queue Queue, data []byte) error {
	if !c.isConnected {
		// TODO: Is should wait connect to RabbitMQ or not?
		return errors.New("failed to push push: not connected")
	}
	for {
		err := c.UnsafePush(queue, data)
		if err != nil {
			log.Error("Client.Push MQ issue " + err.Error())
			if err == ErrDisconnected {
				continue
			}
			return err
		}
		select {
		case confirm := <-c.notifyConfirm:
			if confirm.Ack {
				return nil
			}
		case <-time.After(resendDelay):
		}
	}
}

func (c *Client) UnsafePush(queue Queue, data []byte) error {
	if !c.isConnected {
		return ErrDisconnected
	}
	return c.channel.Publish(
		"",            // Exchange
		string(queue), // Routing key
		false,         // Mandatory
		false,         // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
}

func (s *Stream) Connect(cancelCtx context.Context) {
	s.isConnected = true
	for {
		if s.client.isConnected {
			break
		}
		time.Sleep(1 * time.Second)
	}
	s.client.DeclareQueue(s.consumer)
	if !s.isStream {
		return
	}
	messageChannel, err := s.client.channel.Consume(
		s.consumer.GetQueue(),
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		s.isConnected = false
		log.Fatal("GetMessageChannel MQ issue "+err.Error(), s.consumer.GetQueue())
	}
	for {
		select {
		case <-cancelCtx.Done():
			log.Info("Consumer stopped")
			return
		case msg, ok := <-messageChannel:
			if !ok {
				s.isConnected = false
				return
			}
			if msg.Body != nil {
				s.delivery(msg)
			}
		}
	}
}

func (s *Stream) delivery(msg amqp.Delivery) {
	if s.consumer.Callback(msg) == nil {
		ack(s.consumer.GetQueue(), msg)
	}
}

func ack(queue string, msg amqp.Delivery) {
	err := msg.Ack(false)
	if err != nil {
		log.Error("Stream Ack MQ issue on queue: ", queue, err)
	}
}
