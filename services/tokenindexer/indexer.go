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
		log.WithFields(log.Fields{"service": TokenIndexer, "body": string(delivery.Body)}).Error("failed to get transactions: ", err)
		return err
	}
	transactions = transactions.FilterTransactionsByType([]types.TransactionType{
		types.TxTokenTransfer,
		types.TxNativeTokenTransfer,
	})

	if len(transactions) == 0 {
		return nil
	}

	// Add new assets to db

	assets := GetAssetsFromTransactions(transactions)
	err = database.AddNewAssets(assets)
	if err != nil {
		log.WithFields(log.Fields{"service": TokenIndexer}).Error("Failed to add new assets", err)
		return err
	}

	// Add asset <> address association
	addressAssetsMap := assetsMap(transactions)

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
			association := models.SubscriptionsAssetAssociation{
				SubscriptionId: subscription.ID,
				AssetId:        asset.ID,
			}
			associations = append(associations, association)
		}
	}

	return associations, nil
}

func GetAssetsFromTransactions(txs []types.Tx) []models.Asset {
	var assets []models.Asset
	for _, tx := range txs {
		asset, ok := models.AssetFrom(tx)
		if !ok {
			continue
		}
		assets = append(assets, asset)
	}
	return assets
}

func assetsMap(txs types.Txs) map[string][]string {
	result := make(map[string][]string)
	for _, tx := range txs {
		prefix := strconv.Itoa(int(tx.Coin)) + "_"
		addresses := tx.GetAddresses()
		asset, ok := models.AssetFrom(tx)
		if !ok {
			continue
		}
		for _, address := range addresses {
			assetIDs := result[prefix+address]
			result[prefix+address] = append(assetIDs, asset.Asset)
		}
	}
	return result
}
