package fio

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"encoding/json"
	"strconv"
	"strings"
	"time"
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
			// convert to action-specific data
			var actionData ActionData
			dataJson, _ := json.Marshal(action.ActionTrace.Act.Data)
			dataErr := json.Unmarshal(dataJson, &actionData)
			if dataErr != nil {
				return blockatlas.Tx{}, errors.E("Unparseable Data")
			}
			from = actionData.From
			to = actionData.To
			amountNum, err := strconv.ParseFloat(strings.Split(actionData.Quantity, " ")[0], 64)
			if err == nil {
				amount = blockatlas.Amount(strconv.Itoa(int(amountNum * 1000000000)))
			}
		//}
		//if action.ActionTrace.Act.Name == "trnsfiopubky" {
		//	to = actionData.PayeePublicKey
		//	amount = blockatlas.Amount(string(actionData.Amount))
		//}
		date, _ := time.Parse("2006-01-02T15:04:05", action.BlockTime)
		tx := blockatlas.Tx{
			ID:     action.ActionTrace.TrxID,
			Coin:   p.Coin().ID,
			Date:   date.Unix(),
			From:   from,
			To:     to,
			Block:  action.BlockNum,
			Status: blockatlas.StatusCompleted,
			Fee:    "0", // trnsfiopubky: actionData.Fee
			Meta: blockatlas.Transfer{
				Value:    amount,
				Symbol:   p.Coin().Symbol,
				Decimals: p.Coin().Decimals,
			},
			Memo:   actionData.Memo,
			Type:   blockatlas.TxTransfer,
		}
		tx.Direction = tx.GetTransactionDirection(account)
		return tx, nil
	}
	return blockatlas.Tx{}, errors.E("Unknown action")
}
