package iotex

import (
	"strconv"
	"time"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	txs := make(types.Txs, 0)
	var start int64

	totalTrx, err := p.client.GetAddressTotalTransactions(address)
	if err != nil {
		return nil, err
	}

	if totalTrx >= types.TxPerPage {
		start = totalTrx - types.TxPerPage
	}

	actions, err := p.client.GetTxsOfAddress(address, start)
	if err != nil {
		return nil, err
	}

	for _, srcTx := range actions.ActionInfo {
		tx := Normalize(srcTx)
		if tx != nil {
			txs = append(txs, *tx)
		}
	}

	return txs, nil
}

// Normalize converts an Iotex transaction into the generic model
func Normalize(trx *ActionInfo) *types.Tx {
	if trx.Action == nil {
		return nil
	}
	if trx.Action.Core == nil {
		return nil
	}
	if trx.Action.Core.Transfer == nil {
		return nil
	}

	date, err := time.Parse(time.RFC3339, trx.Timestamp)
	if err != nil {
		return nil
	}
	height, err := strconv.ParseInt(trx.BlkHeight, 10, 64)
	if err != nil {
		return nil
	}
	if height <= 0 {
		return nil
	}
	nonce, err := strconv.ParseInt(trx.Action.Core.Nonce, 10, 64)
	if err != nil {
		return nil
	}
	if trx.GasFee == "" {
		trx.GasFee = "0"
	}
	return &types.Tx{
		ID:       trx.ActHash,
		Coin:     coin.IOTEX,
		From:     trx.Sender,
		To:       trx.Action.Core.Transfer.Recipient,
		Fee:      types.Amount(trx.GasFee),
		Date:     date.Unix(),
		Block:    uint64(height),
		Status:   types.StatusCompleted,
		Sequence: uint64(nonce),
		Type:     types.TxTransfer,
		Meta: types.Transfer{
			Value:    trx.Action.Core.Transfer.Amount,
			Symbol:   coin.Iotex().Symbol,
			Decimals: coin.Iotex().Decimals,
		},
	}
}
