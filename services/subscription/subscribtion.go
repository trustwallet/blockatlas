package subscription

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/storage"
)

const (
	addSubscription    blockatlas.SubscriptionOperation = "AddSubscription"
	deleteSubscription blockatlas.SubscriptionOperation = "DeleteSubscription"
)

func Consume(message []byte, storage storage.Addresses) bool {
	var event blockatlas.SubscriptionEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		logger.Fatal(err)
	}
	subscription := event.ParseSubscriptions()

	switch event.Operation {
	case addSubscription:
		err := storage.AddSubscriptions(subscription)
		if err != nil {
			logger.Fatal(err, logger.Params{"operation": event.Operation, "guid": event.GUID})
		}
		return true
	case deleteSubscription:
		err := storage.DeleteSubscriptions(nil)
		if err != nil {
			logger.Fatal(logger.Params{"operation": event.Operation, "guid": event.GUID})
		}
		return true
	}

	return false
}
