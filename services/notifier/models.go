package notifier

import (
	"github.com/trustwallet/golibs/txtype"
)

type TransactionNotification struct {
	Action txtype.TransactionType `json:"action"`
	Result txtype.Tx              `json:"result"`
}

func buildNotificationsByAddress(address string, txs txtype.Txs) []TransactionNotification {
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

func toUniqueTransactions(txs []txtype.Tx) []txtype.Tx {
	keys := make(map[string]bool)
	var list []txtype.Tx
	for _, entry := range txs {
		key := entry.ID + string(entry.Direction)
		if _, value := keys[key]; !value {
			keys[key] = true
			list = append(list, entry)
		}
	}
	return list
}

func findTransactionsByAddress(txs txtype.Txs, address string) []txtype.Tx {
	result := make([]txtype.Tx, 0)
	for _, tx := range txs {
		if containsAddress(tx, address) {
			result = append(result, tx)
		}
	}
	return result
}

func containsAddress(tx txtype.Tx, address string) bool {
	allAddresses := tx.GetAddresses()
	txAddresses := ToUniqueAddresses(allAddresses)
	for _, a := range txAddresses {
		if a == address {
			return true
		}
	}
	return false
}
