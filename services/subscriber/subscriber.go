package subscriber

import (
	"encoding/json"

	"github.com/trustwallet/blockatlas/internal"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/golibs/types"
)

func RunSubscriber(database *db.Instance, delivery amqp.Delivery) error {
	var event types.SubscriptionEvent
	err := json.Unmarshal(delivery.Body, &event)
	if err != nil {
		log.WithFields(log.Fields{"service": types.Notifications, "body": string(delivery.Body), "error": err}).Error("Unable to unmarshal MQ Message")
		return nil
	}

	subscriptions := event.ParseSubscriptions(event.Subscriptions)
	switch event.Operation {
	case types.AddSubscription:
		err := database.CreateSubscriptions(subscriptions)
		if err != nil {
			log.WithFields(log.Fields{"service": types.Notifications, "operation": event.Operation, "subscriptions": subscriptions}).Error(err)
			return err
		}
		log.WithFields(log.Fields{"service": types.Notifications, "operation": event.Operation, "subscriptions": len(subscriptions)}).Info("Add subscriptions")
	case types.DeleteSubscription:
		subscriptionsIds := make([]string, 0)
		for _, subscription := range subscriptions {
			subscriptionsIds = append(subscriptionsIds, subscription.AddressID())
		}
		return database.DeleteSubscriptions(subscriptionsIds)
	}

	// Pass over subscribed addresses to find all associated tokens to such addresses
	err = internal.SubscriptionsTokens.Publish(delivery.Body)
	if err != nil {
		log.Error(err)
		return nil
	}

	return nil
}
