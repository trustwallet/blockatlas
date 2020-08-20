package notifier

import (
	"context"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"go.elastic.co/apm"
)

type TransactionNotification struct {
	Action blockatlas.TransactionType `json:"action"`
	Result blockatlas.Tx              `json:"result"`
}

func getNotificationBatches(notifications []TransactionNotification, sizeUint uint, ctx context.Context) [][]TransactionNotification {
	span, _ := apm.StartSpan(ctx, "getNotificationBatches", "app")
	defer span.End()
	size := int(sizeUint)
	resultLength := (len(notifications) + size - 1) / size
	result := make([][]TransactionNotification, resultLength)
	lo, hi := 0, size
	for i := range result {
		if hi > len(notifications) {
			hi = len(notifications)
		}
		result[i] = notifications[lo:hi:hi]
		lo, hi = hi, hi+size
	}
	return result
}

func buildNotificationsByAddress(address string, txs blockatlas.Txs, ctx context.Context) []TransactionNotification {
	span, _ := apm.StartSpan(ctx, "buildNotification", "app")
	defer span.End()

	transactionsByAddress := toUniqueTransactions(findTransactionsByAddress(txs, address))

	result := make([]TransactionNotification, 0, len(transactionsByAddress))
	for _, tx := range transactionsByAddress {
		tx.Direction = tx.GetTransactionDirection(address)
		tx.InferUtxoValue(address, tx.Coin)
		result = append(result, TransactionNotification{Action: tx.Type, Result: tx})
	}

	return result
}

func ToUniqueAddresses(addresses []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range addresses {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func toUniqueTransactions(txs []blockatlas.Tx) []blockatlas.Tx {
	keys := make(map[string]bool)
	var list []blockatlas.Tx
	for _, entry := range txs {
		key := entry.ID + string(entry.Direction)
		if _, value := keys[key]; !value {
			keys[key] = true
			list = append(list, entry)
		}
	}
	return list
}

func findTransactionsByAddress(txs blockatlas.Txs, address string) []blockatlas.Tx {
	result := make([]blockatlas.Tx, 0)
	for _, tx := range txs {
		if containsAddress(tx, address) {
			result = append(result, tx)
		}
	}
	return result
}

func containsAddress(tx blockatlas.Tx, address string) bool {
	allAddresses := tx.GetAddresses()
	txAddresses := ToUniqueAddresses(allAddresses)
	for _, a := range txAddresses {
		if a == address {
			return true
		}
	}
	return false
}
