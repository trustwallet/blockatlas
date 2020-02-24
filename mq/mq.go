package mq

import (
	"github.com/streadway/amqp"
)

var (
	amqpChan *amqp.Channel
	conn     *amqp.Connection
	queue    amqp.Queue
)

type QueueName string

const Notifications QueueName = "transactions"

func Init(uri string) (err error) {
	conn, err = amqp.Dial(uri)
	if err != nil {
		return
	}
	amqpChan, err = conn.Channel()
	if err != nil {
		return
	}
	queue, err = amqpChan.QueueDeclare(string(Notifications), true, false, false, false, nil)
	return
}

func Close() {
	amqpChan.Close()
	conn.Close()
}

func Publish(body []byte) error {
	return amqpChan.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
}
