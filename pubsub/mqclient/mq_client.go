package mqclient

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/pubsub"
	"go.uber.org/atomic"
	"time"
)

var (
	reconnectDelay = 5 * time.Second
	resendDelay    = 5 * time.Second
)

type Client struct {
	uri           string
	conn          *amqp.Connection
	channel       *amqp.Channel
	prefetchCount int
	streams       []*pubsub.Stream
	ctx           context.Context
	notifyClose   chan *amqp.Error
	notifyConfirm chan amqp.Confirmation
	isConnected   *atomic.Bool
	alive         *atomic.Bool
}

func New(uri string, prefetchCount int, ctx context.Context) (client pubsub.Client) {
	client = &Client{
		uri:           uri,
		ctx:           ctx,
		alive:         atomic.NewBool(false),
		isConnected:   atomic.NewBool(false),
		prefetchCount: prefetchCount,
		streams:       []*pubsub.Stream{},
	}
	return client
}

func (c *Client) Connect() error {
	conn, err := amqp.Dial(c.uri)
	if err != nil {
		log.Error("Client.connect MQ Dial issue " + err.Error())
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Error("Client.connect MQ Channel issue " + err.Error())
		return err
	}
	err = ch.Confirm(false)
	if err != nil {
		log.Error("Client.connect MQ Channel issue " + err.Error())
	}
	log.Debug("Interface: ", c)
	c.conn = conn
	c.channel = ch
	c.notifyClose = make(chan *amqp.Error)
	c.notifyConfirm = make(chan amqp.Confirmation)
	c.channel.NotifyClose(c.notifyClose)
	c.channel.NotifyPublish(c.notifyConfirm)
	c.isConnected.Store(true)
	return nil
}

func (c *Client) Run() error {
	log.Debug("Interface: ", c)
	if c.conn == nil {
		return errors.New("connect firstly")
	}
	go c.handleReconnect()
	err := c.channel.Qos(c.prefetchCount, 0, false)
	if err != nil {
		return errors.New("Client.connect MQ Qos issue " + err.Error())
	}
	for _, stream := range c.streams {
		if !(*stream).IsConnected() {
			go (*stream).Connect(c.ctx)
		}
	}
	return nil
}

func (c *Client) IsConnected() bool {
	return c.isConnected.Load()
}

func (c *Client) AddStream(consumer *pubsub.Consumer, isWriteOnly bool) error {
	var stream pubsub.Stream = &Stream{
		consumer:    consumer,
		client:      c,
		isConnected: atomic.NewBool(false),
		isWriteOnly: isWriteOnly,
		channel:     c.channel,
	}
	go stream.Connect(c.ctx) // Try connect, if client isn't run it will wait run
	c.streams = append(c.streams, &stream)
	return nil
}

func (c *Client) Push(queue string, data []byte) error {
	if !c.isConnected.Load() {
		// TODO: Is should wait connect to RabbitMQ or not?
		return errors.New("failed to push push: not connected")
	}

	//todo Add stream waiting
	for {
		err := c.PushUnsafe(queue, data)
		if err != nil {
			log.Error("Client.Push MQ issue " + err.Error())
			if err == pubsub.ErrDisconnected {
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

func (c *Client) PushUnsafe(queue string, data []byte) error {
	if !c.isConnected.Load() {
		return pubsub.ErrDisconnected
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

func (c *Client) Close() error {
	err := c.channel.Close()
	if err != nil {
		return errors.New("Client.Close MQ issue " + err.Error())
	}

	err = c.conn.Close()
	if err != nil {
		return errors.New("Client.Close MQ issue " + err.Error())
	}
	return nil
}

func (c *Client) handleReconnect() {
	for c.alive.Load() {
		log.Debug("Try connect after alive")
		c.isConnected.Store(false)
		for c.Connect() != nil {
			log.Debug("Try connect")
			if !c.alive.Load() {
				return
			}
			select {
			case <-c.ctx.Done():
				return
			case <-time.After(reconnectDelay + time.Second):
				// Add metric
			}
		}
		select {
		case <-c.ctx.Done():
			log.Fatal("Fucking end")
			return
		case <-c.notifyClose:
		}
	}
}
