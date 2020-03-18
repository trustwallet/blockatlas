package subscriber

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
)

const (
	AddSubscription    blockatlas.SubscriptionOperation = "AddSubscription"
	DeleteSubscription blockatlas.SubscriptionOperation = "DeleteSubscription"
	UpdateSubscription blockatlas.SubscriptionOperation = "UpdateSubscription"
)

func RunSubscriber(delivery amqp.Delivery, storage storage.Addresses) {
	var event blockatlas.SubscriptionEvent
	err := json.Unmarshal(delivery.Body, &event)
	if err != nil {
		logger.Fatal(err)
	}
	newSubscriptions := event.ParseSubscriptions(event.NewSubscriptions)
	oldSubscriptions := event.ParseSubscriptions(event.OldSubscriptions)

	params := logger.Params{"operation": event.Operation, "guid": event.GUID, "new_subscriptions_len": len(newSubscriptions), "old_subscriptions_len": len(oldSubscriptions)}

	switch event.Operation {
	case UpdateSubscription:
		err := storage.DeleteSubscriptions(oldSubscriptions)
		if err != nil {
			logger.Error(err, params)
		}
		err = storage.AddSubscriptions(newSubscriptions)
		if err != nil {
			logger.Error(err, params)
		}
		err = delivery.Ack(false)
		if err != nil {
			logger.Error(err, params)
		}
		logger.Info("Updated", params)
	case AddSubscription:
		err = storage.AddSubscriptions(newSubscriptions)
		if err != nil {
			logger.Error(err, params)
		}
		err = delivery.Ack(false)
		if err != nil {
			logger.Error(err, params)
		}
		logger.Info("Added", params)
	case DeleteSubscription:
		err := storage.DeleteSubscriptions(oldSubscriptions)
		if err != nil {
			logger.Error(err, params)
		}
		err = delivery.Ack(false)
		if err != nil {
			logger.Error(err, params)
		}
		logger.Info("Deleted", params)
	}
}
