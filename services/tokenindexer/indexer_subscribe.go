package tokenindexer

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/types"
)

type ConsumerIndexer struct {
	Database   *db.Instance
	TokensAPIs map[uint]blockatlas.TokensAPI
	Delivery   func(*db.Instance, map[uint]blockatlas.TokensAPI, amqp.Delivery) error
	Tag        string
}

func (c ConsumerIndexer) Callback(msg amqp.Delivery) error {
	return c.Delivery(c.Database, c.TokensAPIs, msg)
}

func RunTokenIndexerSubscribe(database *db.Instance, apis map[uint]blockatlas.TokensAPI, delivery amqp.Delivery) error {
	var event types.SubscriptionEvent
	err := json.Unmarshal(delivery.Body, &event)
	if err != nil {
		log.WithFields(log.Fields{"service": SubscriptionsTokenIndexer, "body": string(delivery.Body), "error": err}).Error("Unable to unmarshal MQ Message")
		return nil
	}

	log.WithFields(log.Fields{"service": TokenIndexer, "event": event.Operation, "subscriptions": len(event.Subscriptions)}).Info("Processing")

	subscriptions := event.ParseSubscriptions(event.Subscriptions)
	switch event.Operation {
	case types.AddSubscription:
		addressAssetsMap := map[string][]string{}

		for _, coinAddress := range subscriptions {
			api, ok := apis[coinAddress.Coin]
			if !ok {
				continue
			}
			assetIds, err := api.GetTokenListIdsByAddress(coinAddress.Address)
			if err != nil {
				continue
			}
			addressAssetsMap[coinAddress.AddressID()] = assetIds
		}
		return CreateAssociations(database, addressAssetsMap)
	case types.DeleteSubscription:
		//No action is needed
		return nil
	}

	return nil
}
