package subscription

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
)

func Consume(delivery amqp.Delivery, storage storage.Addresses) {
	var event blockatlas.SubscriptionEvent
	err := json.Unmarshal(delivery.Body, &event)
	if err != nil {
		logger.Fatal(err)
	}
	subscriptions := event.ParseSubscriptions()

	switch event.Operation {
	case AddSubscription:
		err := storage.AddSubscriptions(subscriptions)
		if err != nil {
			logger.Fatal(err, logger.Params{"operation": event.Operation, "guid": event.GUID})
		}
		logger.Info("Success", logger.Params{"operation": event.Operation, "guid": event.GUID, "subscriptions_len": len(subscriptions)})
	case DeleteSubscription:
		err := storage.DeleteSubscriptions(nil)
		if err != nil {
			logger.Fatal(err, logger.Params{"operation": event.Operation, "guid": event.GUID})
		}
		logger.Info("Success", logger.Params{"operation": event.Operation, "guid": event.GUID, "subscriptions_len": len(subscriptions)})
	}
}
