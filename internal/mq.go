package internal

import (
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/golibs/network/mq"
)

const (
	// End consumer of published transactions. Not consumed on blockatlas
	TxNotifications mq.Queue = "txNotifications"
	// Address:coin subscriptions
	Subscriptions       mq.Queue = "subscriptions"
	SubscriptionsTokens mq.Queue = "subscriptions_tokens"

	// Transactions to process, if match subscriptions, pushed to TxNotifications
	RawTransactions         mq.Queue    = "rawTransactions"
	RawTokens               mq.Queue    = "rawTokens"
	RawTransactionsExchange mq.Exchange = "raw_transactions"
)

type ConsumerDatabase struct {
	Database *db.Instance
	Delivery func(*db.Instance, amqp.Delivery) error
	Tag      string
}

func (c ConsumerDatabase) Callback(msg amqp.Delivery) error {
	return c.Delivery(c.Database, msg)
}
