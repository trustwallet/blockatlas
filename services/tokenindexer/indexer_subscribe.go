package tokenindexer

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/subscriber"
)

type ConsumerIndexer struct {
	Database   *db.Instance
	TokensAPIs map[uint]blockatlas.TokensAPI
	Delivery   func(*db.Instance, map[uint]blockatlas.TokensAPI, amqp.Delivery) error
}

func (c ConsumerIndexer) Callback(msg amqp.Delivery) error {
	return c.Delivery(c.Database, c.TokensAPIs, msg)
}

func RunTokenIndexerSubscribe(database *db.Instance, apis map[uint]blockatlas.TokensAPI, delivery amqp.Delivery) error {
	var event blockatlas.SubscriptionEvent
	err := json.Unmarshal(delivery.Body, &event)
	if err != nil {
		log.WithFields(log.Fields{"service": SubscriptionsTokenIndexer, "event": event}).Error(err)
		return err
	}

	subscriptions := event.ParseSubscriptions(event.Subscriptions)
	switch event.Operation {
	case subscriber.AddSubscription:
		addressAssetsMap := map[string][]string{}

		for _, coinAddress := range subscriptions {
			api, ok := apis[coinAddress.Coin]
			if !ok {
				continue
			}
			assets, err := api.GetTokenListByAddress(coinAddress.Address)
			if err != nil {
				continue
			}
			assetIds := make([]string, 0)
			for _, asset := range assets {
				assetIds = append(assetIds, asset.AssetId())
			}
			addressAssetsMap[coinAddress.AddressID()] = assetIds
		}
		return CreateAssociations(database, addressAssetsMap)
	case subscriber.DeleteSubscription:
		//No action is needed
		break
	}

	return nil
}
