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

func RunSubscriber(database *db.Instance, delivery amqp.Delivery) {
	var event blockatlas.SubscriptionEvent
	err := json.Unmarshal(delivery.Body, &event)
	if err != nil {
		logger.Fatal(err)
	}

	subscriptions := event.ParseSubscriptions(event.Subscriptions)

	params := logger.Params{"operation": event.Operation, "id": event.Id, "subscriptions_len": len(subscriptions)}

	id := event.Id

	switch event.Operation {
	case AddSubscription, UpdateSubscription:
		err = database.AddSubscriptions(id, ToSubscriptionData(subscriptions))
		if err != nil {
			logger.Error(err, params)
		}
		logger.Info("Added", params)
	case DeleteSubscription:
		err := database.DeleteAllSubscriptions(id)
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
		data = append(data, models.SubscriptionData{Coin: s.Coin, Address: s.Address, SubscriptionId: s.Id})
	}
	return data
}
