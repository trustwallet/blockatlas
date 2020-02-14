package nano

import (
	"encoding/json"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	normalized := make([]blockatlas.Tx, 0)
	history, err := p.client.GetAccountHistory(address)
	if err != nil {
		return normalized, err
	}
	b, err := json.Marshal(history.History)
	if err != nil {
		return normalized, nil
	}
	var txs []Transaction
	err = json.Unmarshal(b, &txs)
	if err != nil {
		return normalized, nil
	}

	for _, srcTx := range txs {
		tx := p.Normalize(&srcTx, history.Account)
		normalized = append(normalized, tx)
	}

	return normalized, nil
}

func (p *Platform) Normalize(srcTx *Transaction, account string) (tx blockatlas.Tx) {
	var from string
	var to string

	if srcTx.Type == BlockTypeSend {
		from = account
		to = srcTx.Account
	} else if srcTx.Type == BlockTypeReceive {
		from = srcTx.Account
		to = account
	}

	status := blockatlas.StatusCompleted
	height, _ := strconv.ParseUint(srcTx.Height, 10, 64)
	if height == 0 {
		status = blockatlas.StatusPending
	}
	timestamp, _ := strconv.ParseInt(srcTx.LocalTimestamp, 10, 64)

	tx = blockatlas.Tx{
		ID:     srcTx.Hash,
		Coin:   p.Coin().ID,
		Date:   timestamp,
		From:   from,
		To:     to,
		Block:  height,
		Status: status,
		Fee:    "0",
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Amount),
			Symbol:   p.Coin().Symbol,
			Decimals: p.Coin().Decimals,
		},
	}
	return tx
}
