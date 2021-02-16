package near

import (
	"encoding/json"
	"errors"

	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	normalized := make(types.Txs, 0)
	return normalized, nil
}

func (p *Platform) GetTransactionsByAccount(account, token string, limit int, database *db.Instance) (page types.Txs, err error) {
	txs, err := database.GetTransactionsByAccount(account, p.Coin().ID, limit)
	if err != nil {
		return
	}
	return models.ToTxs(txs)
}

func NormalizeChunk(chunk ChunkDetail) types.Txs {
	normalized := make(types.Txs, 0)
	for _, tx := range chunk.Transactions {
		if len(tx.Actions) != 1 {
			continue
		}

		transfer, err := mapTransfer(tx.Actions[0])
		if err != nil {
			continue
		}

		normalized = append(normalized, types.Tx{
			ID:       tx.Hash,
			Coin:     coin.NEAR,
			From:     tx.SignerID,
			To:       tx.ReceiverID,
			Fee:      "0",
			Date:     int64(chunk.Header.Timestamp),
			Block:    chunk.Header.Height,
			Status:   types.StatusCompleted,
			Sequence: uint64(tx.Nonce),
			Type:     types.TxTransfer,
			Meta: types.Transfer{
				Value:    types.Amount(transfer.Transfer.Deposit),
				Symbol:   coin.Near().Name,
				Decimals: coin.Near().Decimals,
			},
		})
	}

	return normalized
}

func mapTransfer(i interface{}) (action TransferAction, err error) {
	bytes, err := json.Marshal(i)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &action)
	if err != nil {
		return
	}

	if action.Transfer.Deposit == "" {
		err = errors.New("unable marshalling to transfer actoin struct")
	}
	return
}
