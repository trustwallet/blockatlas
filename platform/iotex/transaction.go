package iotex

import (
	"strconv"
	"time"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	txs := make([]txtype.Tx, 0)
	var start int64

	totalTrx, err := p.client.GetAddressTotalTransactions(address)
	if err != nil {
		return nil, err
	}

	if totalTrx >= txtype.TxPerPage {
		start = totalTrx - txtype.TxPerPage
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
func Normalize(trx *ActionInfo) *txtype.Tx {
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
	return &txtype.Tx{
		ID:       trx.ActHash,
		Coin:     coin.IOTX,
		From:     trx.Sender,
		To:       trx.Action.Core.Transfer.Recipient,
		Fee:      txtype.Amount(trx.GasFee),
		Date:     date.Unix(),
		Block:    uint64(height),
		Status:   txtype.StatusCompleted,
		Sequence: uint64(nonce),
		Type:     txtype.TxTransfer,
		Meta: txtype.Transfer{
			Value:    trx.Action.Core.Transfer.Amount,
			Symbol:   coin.Coins[coin.IOTX].Symbol,
			Decimals: coin.Coins[coin.IOTX].Decimals,
		},
	}
}
