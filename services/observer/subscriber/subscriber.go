package subscriber

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

const (
	AddSubscription    blockatlas.SubscriptionOperation = "AddSubscription"
	DeleteSubscription blockatlas.SubscriptionOperation = "DeleteSubscription"
	UpdateSubscription blockatlas.SubscriptionOperation = "UpdateSubscription"
)

func RunSubscriber(delivery amqp.Delivery) {
	var event blockatlas.SubscriptionEvent
	err := json.Unmarshal(delivery.Body, &event)
	if err != nil {
		logger.Fatal(err)
	}

	newSubscriptions := event.ParseSubscriptions(event.NewSubscriptions)
	oldSubscriptions := event.ParseSubscriptions(event.OldSubscriptions)

	params := logger.Params{"operation": event.Operation, "guid": event.GUID, "new_subscriptions_len": len(newSubscriptions), "old_subscriptions_len": len(oldSubscriptions)}

	guid := event.GUID

	switch event.Operation {
	case UpdateSubscription:
		err := db.AddToExistingSubscription(guid, ToSubscriptionData(newSubscriptions))
		if err != nil {
			logger.Error(err, params)
		}
		logger.Info("Updated", params)
	case AddSubscription:
		err = db.AddSubscriptions(guid, ToSubscriptionData(newSubscriptions))
		if err != nil {
			logger.Error(err, params)
		}
		logger.Info("Added", params)
	case DeleteSubscription:
		err := db.DeleteSubscriptions(ToSubscriptionData(oldSubscriptions))
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

func ToSubscriptionData(sub []blockatlas.Subscription) []models.SubscriptionData {
	data := make([]models.SubscriptionData, 0, len(sub))
	for _, s := range sub {
		data = append(data, models.SubscriptionData{Coin: s.Coin, Address: s.Address, SubscriptionId: s.GUID})
	}
	return data
}
