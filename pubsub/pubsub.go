package pubsub

import (
	"context"
	"errors"
	"github.com/streadway/amqp"
)

var (
	ErrDisconnected = errors.New("disconnected from rabbitmq, trying to reconnect")
)

type Client interface {
	Connect() error
	Run() error
	IsConnected() bool
	AddStream(consumer *Consumer, isWriteOnly bool) error
	Push(queue string, data []byte) error
	PushUnsafe(queue string, data []byte) error
	Close() error
}

type Stream interface {
	Connect(cancelCtx context.Context)
	GetConsumer() *Consumer
	GetClient() *Client
	IsConnected() bool
	IsWriteOnly() bool
}

type Consumer interface {
	GetQueue() string
	Callback(msg amqp.Delivery) error
}
