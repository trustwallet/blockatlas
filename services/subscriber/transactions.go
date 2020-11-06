package subscriber

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"go.elastic.co/apm"
	"strconv"
)

type Subscriber string

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
		return
	}

	subscriptions := event.ParseSubscriptions(event.Subscriptions)
	switch event.Operation {
	case AddSubscription, UpdateSubscription:
		err = database.AddSubscriptionsForNotifications(ToSubscriptionData(subscriptions), ctx)
		if err != nil {
			log.WithFields(
				log.Fields{"service": Notifications,
					"operation":         event.Operation,
					"subscriptions_len": len(subscriptions),
				},
			).Error(err)
		}
		log.WithFields(
			log.Fields{"service": Notifications,
				"operation":         event.Operation,
				"subscriptions_len": len(subscriptions),
			},
		).Info("Added")
	case DeleteSubscription:
		err := database.DeleteSubscriptionsForNotifications(ToSubscriptionData(subscriptions), ctx)
		if err != nil {
			log.WithFields(
				log.Fields{"service": Notifications,
					"operation":         event.Operation,
					"subscriptions_len": len(subscriptions),
				},
			).Error(err)
		}
		log.WithFields(
			log.Fields{"service": Notifications,
				"operation":         event.Operation,
				"subscriptions_len": len(subscriptions),
			},
		).Info("Added")
	}

	defer func() {
		err = delivery.Ack(false)
		if err != nil {
			log.WithFields(
				log.Fields{"service": Notifications,
					"operation":         event.Operation,
					"subscriptions_len": len(subscriptions),
				},
			).Error(err)
		}
	}()
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
