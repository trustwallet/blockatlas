package subscriber

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"go.elastic.co/apm"
	"strconv"
)

const (
	Notifications      Subscriber                       = "notifications"
	AddSubscription    blockatlas.SubscriptionOperation = "AddSubscription"
	DeleteSubscription blockatlas.SubscriptionOperation = "DeleteSubscription"
	UpdateSubscription blockatlas.SubscriptionOperation = "UpdateSubscription"
)

func RunTransactionsSubscriber(database *db.Instance, delivery amqp.Delivery) {
	tx := apm.DefaultTracer.StartTransaction("RunTransactionsSubscriber", "app")
	defer tx.End()

	ctx := apm.ContextWithTransaction(context.Background(), tx)

	var event blockatlas.SubscriptionEvent
	err := json.Unmarshal(delivery.Body, &event)
	if err != nil {
		errAck := delivery.Ack(false)
		logger.Fatal(err, errAck)
	}

	subscriptions := event.ParseSubscriptions(event.Subscriptions)
	params := logger.Params{"operation": event.Operation, "subscriptions_len": len(subscriptions)}

	switch event.Operation {
	case AddSubscription, UpdateSubscription:
		err = database.AddSubscriptionsForNotifications(ToSubscriptionData(subscriptions), ctx)
		if err != nil {
			logger.Error(err, params)
		}
		logger.Info("Added", params)
	case DeleteSubscription:
		err := database.DeleteSubscriptionsForNotifications(ToSubscriptionData(subscriptions), ctx)
		if err != nil {
			logger.Error(err, params)
		}
		logger.Info("Deleted", params)
	}

	err = delivery.Ack(false)
	if err != nil {
		logger.Error(err, params)
	}
}

func ToSubscriptionData(sub []blockatlas.Subscription) []string {
	data := make([]string, 0, len(sub))
	for _, s := range sub {
		coinStr := strconv.FormatUint(uint64(s.Coin), 10)
		address := coinStr + "_" + s.Address
		data = append(data, address)
	}
	return data
}
