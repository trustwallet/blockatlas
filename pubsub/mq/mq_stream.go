package mq

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/pubsub"
	"time"
)

type Stream struct {
	consumer    *pubsub.Consumer
	client      *pubsub.Client
	channel     *amqp.Channel
	isConnected bool
	isWriteOnly bool
}

func (s Stream) Connect(cancelCtx context.Context) {
	s.isConnected = true
	for {
		if (*s.client).IsConnected() {
			break
		}
		time.Sleep(1 * time.Second)
	}
	s.declareQueue()
	if !s.isWriteOnly {
		return
	}
	messageChannel, err := s.channel.Consume(
		(*s.consumer).GetQueue(),
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		s.isConnected = false
		log.Fatal("GetMessageChannel MQ issue "+err.Error(), (*s.consumer).GetQueue())
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
func (s Stream) GetConsumer() *pubsub.Consumer {
	return s.consumer
}

func (s Stream) GetClient() *pubsub.Client {
	return s.client
}

func (s Stream) IsConnected() bool {
	return s.isConnected
}

func (s Stream) IsWriteOnly() bool {
	return s.isWriteOnly
}

func (s *Stream) declareQueue() {
	_, err := s.channel.QueueDeclare((*s.consumer).GetQueue(), true, false, false, false, nil)
	if err != nil {
		log.Fatal("Stream.Init MQ issue "+err.Error(), (*s.consumer).GetQueue())
	}
}

func (s *Stream) delivery(msg amqp.Delivery) {
	if (*s.consumer).Callback(msg) == nil {
		ack((*s.consumer).GetQueue(), msg)
	}
}

func ack(queue string, msg amqp.Delivery) {
	err := msg.Ack(false)
	if err != nil {
		log.Error("Stream Ack MQ issue on queue: ", queue, err)
	}
}
