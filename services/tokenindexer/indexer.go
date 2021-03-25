package tokenindexer

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/services/notifier"
	"github.com/trustwallet/golibs/types"
)

const (
	TokenIndexer              = "TokenIndexer"
	SubscriptionsTokenIndexer = "SubscriptionsTokenIndexer"
)

func RunTokenIndexer(database *db.Instance, delivery amqp.Delivery) error {
	transactions, err := notifier.GetTransactionsFromDelivery(delivery, TokenIndexer)
	if err != nil {
		log.WithFields(log.Fields{"service": TokenIndexer, "body": string(delivery.Body), "error": err}).Error("Unable to unmarshal MQ Message")
		return nil
	}

	assetsTxs := transactions.FilterTransactionsByType([]types.TransactionType{
		types.TxContractCall,
		types.TxTokenTransfer,
		types.TxNativeTokenTransfer,
	})

	if len(assetsTxs) == 0 {
		return nil
	}

	// Add new assets to db
	assets := GetAssetsFromTransactions(assetsTxs)
	err = database.AddNewAssets(assets)
	if err != nil {
		log.WithFields(log.Fields{"service": TokenIndexer}).Error("Failed to add new assets", err)
		return err
	}

	// Add asset <> address association
	addressAssetsMap := assetsMap(assetsTxs)

	return CreateAssociations(database, addressAssetsMap)
}

func CreateAssociations(database *db.Instance, addressAssetsMap map[string][]string) error {
	associations, err := calculateSubscriptionAssetAssociations(database, addressAssetsMap)
	if err != nil {
		return err
	}
	return database.CreateSubscriptionsAssets(associations)
}

func calculateSubscriptionAssetAssociations(database *db.Instance, addressAssetsMap map[string][]string) ([]models.SubscriptionsAssetAssociation, error) {
	associations := make([]models.SubscriptionsAssetAssociation, 0)

	addressIds := make([]string, 0)
	assetIds := make([]string, 0)
	for addressId, assets := range addressAssetsMap {
		addressIds = append(addressIds, addressId)

		assetIds = append(assetIds, assets...)
	}

	if len(addressIds) == 0 || len(assetIds) == 0 {
		return associations, nil
	}

	subscriptions, err := database.GetSubscriptions(addressIds)
	if err != nil {
		return associations, err
	}

	assets, err := database.GetAssetsByIDs(assetIds)
	if err != nil {
		return associations, err
	}

	assetsMap := map[string]models.Asset{}
	for _, asset := range assets {
		assetsMap[asset.Asset] = asset
	}

	subscriptionsMap := map[string]models.Subscription{}
	for _, subscription := range subscriptions {
		subscriptionsMap[subscription.Address] = subscription
	}

	uniqueMap := map[string]bool{}
	for addressId, assets := range addressAssetsMap {
		subscription, ok := subscriptionsMap[addressId]
		if !ok {
			continue
		}

		for _, assetId := range assets {
			asset, ok := assetsMap[assetId]
			if !ok {
				continue
			}
			subscriptionKey := strconv.Itoa(int(asset.ID)) + "_" + strconv.Itoa(int(subscription.ID))
			if _, ok := uniqueMap[subscriptionKey]; !ok {
				association := models.SubscriptionsAssetAssociation{
					SubscriptionId: subscription.ID,
					AssetId:        asset.ID,
				}
				associations = append(associations, association)
				uniqueMap[subscriptionKey] = true
			}
		}
	}

	return associations, nil
}

func GetAssetsFromTransactions(txs types.Txs) []models.Asset {
	var result []models.Asset
	for _, tx := range txs {
		assets := models.AssetsFrom(tx)
		result = append(result, assets...)
	}
	return result
}

func assetsMap(txs types.Txs) map[string][]string {
	result := make(map[string][]string)
	for _, tx := range txs {
		prefix := strconv.Itoa(int(tx.Coin)) + "_"
		addresses := tx.GetAddresses()
		assets := models.AssetsFrom(tx)

		for _, asset := range assets {
			for _, address := range addresses {
				assetId := prefix + address
				assetIDs := result[assetId]
				result[assetId] = append(assetIDs, asset.Asset)
			}
		}
	}
	return result
}
