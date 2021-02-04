package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/trustwallet/golibs/asset"
	"github.com/trustwallet/golibs/types"
)

type Transaction struct {
	ID         string `gorm:"primary_key; autoIncrement:false; index"`
	Coin       uint   `gorm:"primary_key; autoIncrement:false; index"`
	AssetID    string
	From       string
	To         string
	FeeAssetID string
	Fee        string
	Date       time.Time
	Block      uint64
	Sequence   uint64
	Status     string
	Memo       string
	Type       string
	Metadata   postgres.Jsonb
}

func (tx Transaction) ToTx() (result types.Tx, err error) {
	coinId, _, err := asset.ParseID(tx.AssetID)
	if err != nil {
		return
	}

	bytes, err := tx.Metadata.MarshalJSON()
	if err != nil {
		return
	}

	result = types.Tx{
		ID:     tx.ID,
		Coin:   coinId,
		Date:   tx.Date.Unix(),
		From:   tx.From,
		To:     tx.To,
		Fee:    types.Amount(tx.Fee),
		Block:  tx.Block,
		Status: types.Status(tx.Status),
		Memo:   tx.Memo,
		Type:   types.TransactionType(tx.Type),
	}

	switch result.Type {
	case types.TxTransfer:
		var transfer types.Transfer
		err = json.Unmarshal(bytes, &transfer)
		result.Meta = transfer
	case types.TxTokenTransfer, types.TxNativeTokenTransfer:
		var transfer types.TokenTransfer
		err = json.Unmarshal(bytes, &transfer)
		result.Meta = transfer
	case types.TxContractCall:
		var call types.ContractCall
		err = json.Unmarshal(bytes, &call)
		result.Meta = call
	case types.TxAnyAction:
		var action types.AnyAction
		err = json.Unmarshal(bytes, &action)
		result.Meta = action
	default:
		err = fmt.Errorf("not supported metadata type: %s", tx.Type)
	}
	return
}

func ToTxPage(txs []Transaction) (types.TxPage, error) {
	page := make(types.TxPage, 0)
	for _, tx := range txs {
		t, err := tx.ToTx()
		if err != nil {
			return types.TxPage{}, err
		}
		page = append(page, t)
	}
	return page, nil
}
