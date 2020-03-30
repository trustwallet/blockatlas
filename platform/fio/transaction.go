package fio

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"strconv"
)

func (p *Platform) GetTxsByAddress(address string) (page blockatlas.TxPage, err error) {
	account := address // TODO this should be actor
	actions, err := p.client.getTransactions(account)
	if err != nil {
		return nil, err
	}
	var txs []blockatlas.Tx = make([]blockatlas.Tx, 0)
	for _, a := range actions {
		tx, err := p.Normalize(&a, account)
		if err != nil {
			continue
		}
		txs = append(txs, tx)
	}
	txPage := blockatlas.TxPage(txs)
	return txPage, nil
}

func (p *Platform) Normalize(action *Action, account string) (blockatlas.Tx, error) {
	var from string
	var to string
	var amount blockatlas.Amount

	// Action Act.Name == "trnsfiopubky" not handled
	if action.ActionTrace.Act.Account == "fio.token" &&
	   action.ActionTrace.Act.Name == "transfer" {
		//if action.ActionTrace.Act.Name == "transfer" {
			from = action.ActionTrace.Act.Data.From
			to = action.ActionTrace.Act.Data.To
			amountNum, err := strconv.Atoi(action.ActionTrace.Act.Data.Quantity)
			if err == nil {
				amount = blockatlas.Amount(strconv.Itoa(amountNum * 1000000000))
			}
		//}
		//if action.ActionTrace.Act.Name == "trnsfiopubky" {
		//	to = action.ActionTrace.Act.Data.PayeePublicKey
		//	amount = blockatlas.Amount(string(action.ActionTrace.Act.Data.Amount)) // TODO
		//}
		tx := blockatlas.Tx{
			ID:     action.ActionTrace.TrxID,
			Coin:   p.Coin().ID,
			//Date:   // TODO
			From:   from,
			To:     to,
			Block:  action.BlockNum,
			Status: blockatlas.StatusCompleted,
			Fee:    blockatlas.Amount(action.ActionTrace.Act.Data.Fee),
			Meta: blockatlas.Transfer{
				Value:    amount,
				Symbol:   p.Coin().Symbol,
				Decimals: p.Coin().Decimals,
			},
			Memo:   action.ActionTrace.Act.Data.Memo,
			Type:   blockatlas.TxTransfer,
			//Direction: // TODO
		}
		return tx, nil
	}
	return blockatlas.Tx{}, errors.E("Unknown action")
}
