package subscriber

import (
	"encoding/json"

	"github.com/trustwallet/blockatlas/internal"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Subscriber string

const (
	Notifications      Subscriber                       = "notifications"
	AddSubscription    blockatlas.SubscriptionOperation = "AddSubscription"
	DeleteSubscription blockatlas.SubscriptionOperation = "DeleteSubscription"
)

func RunSubscriber(database *db.Instance, delivery amqp.Delivery) error {
	var event blockatlas.SubscriptionEvent
	err := json.Unmarshal(delivery.Body, &event)
	if err != nil {
		log.WithFields(log.Fields{"service": Notifications, "body": string(delivery.Body)}).Error(err)
		return err
	}

	subscriptions := event.ParseSubscriptions(event.Subscriptions)
	switch event.Operation {
	case AddSubscription:
		err := database.CreateSubscriptions(subscriptions)
		if err != nil {
			log.WithFields(log.Fields{"service": Notifications, "operation": event.Operation, "subscriptions": subscriptions}).Error(err)
			return err
		}
		log.WithFields(log.Fields{"service": Notifications, "operation": event.Operation, "subscriptions": len(subscriptions)}).Info("Add subscriptions")
	case DeleteSubscription:
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
