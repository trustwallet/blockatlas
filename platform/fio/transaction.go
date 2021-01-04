package fio

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"errors"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

const (
	contractToken        = "fio.token"
	contractTreasury     = "fio.treasury"
	actionTransfer       = "transfer"
	actionTransferPubkey = "trnsfiopubky"
)

func (p *Platform) GetTxsByAddress(address string) (page blockatlas.TxPage, err error) {
	// take actor from address
	account := actorFromPublicKeyOrActor(address)
	actions, err := p.client.getTransactions(account)
	if err != nil {
		return nil, err
	}
	txs := make([]blockatlas.Tx, 0)
	for _, a := range actions {
		tx, err := p.Normalize(&a, account)
		if err != nil {
			continue
		}
		txs = append(txs, tx)
	}
	txs = unique(txs)
	txPage := blockatlas.TxPage(txs)
	sort.Sort(txPage)
	return txPage, nil
}

func (p *Platform) Normalize(action *Action, account string) (blockatlas.Tx, error) {
	var (
		to, from, memo string
		amount, fee    blockatlas.Amount
		sequence       uint64
	)
	const dateFormat string = "2006-01-02T15:04:05"

	if action.ActionTrace.Act.Account == contractToken &&
		(action.ActionTrace.Act.Name == actionTransfer || action.ActionTrace.Act.Name == actionTransferPubkey) {
		// convert to action-specific data
		dataJSON, err := json.Marshal(action.ActionTrace.Act.Data)
		if err != nil {
			return blockatlas.Tx{}, errors.New("Unparseable Data field")
		}
		switch action.ActionTrace.Act.Name {
		case actionTransfer:
			var actionData ActionDataTransfer
			if json.Unmarshal(dataJSON, &actionData) != nil {
				return blockatlas.Tx{}, errors.New("Unparseable Data field")
			}
			if actionData.To == "fio.treasury" {
				return blockatlas.Tx{}, errors.New("Skip tx sent to treasury, usually fee")
			}
			from = actionData.From
			to = actionData.To
			amountNum, err := strconv.ParseFloat(strings.Split(actionData.Quantity, " ")[0], 64)
			if err == nil {
				amount = blockatlas.Amount(strconv.Itoa(int(amountNum * 1000000000)))
			}
			// fee unknown
			memo = actionData.Memo
			sequence = action.ActionSeq
		case actionTransferPubkey:
			var actionData ActionDataTrnsfiopubky
			if json.Unmarshal(dataJSON, &actionData) != nil {
				return blockatlas.Tx{}, errors.New("Unparseable Data field")
			}
			from = actionData.Actor
			to = actorFromPublicKeyOrActor(actionData.PayeePublicKey)
			amount = blockatlas.Amount(strconv.FormatInt(actionData.Amount, 10))
			fee = blockatlas.Amount(strconv.FormatInt(actionData.MaxFee, 10))
			// not set sequence because it might be duplicated
		}
		date, _ := time.Parse(dateFormat, action.BlockTime)
		tx := blockatlas.Tx{
			ID:       action.ActionTrace.TrxID,
			Coin:     p.Coin().ID,
			Date:     date.Unix(),
			From:     from,
			To:       to,
			Block:    action.BlockNum,
			Sequence: sequence,
			Status:   blockatlas.StatusCompleted,
			Fee:      fee,
			Meta: blockatlas.Transfer{
				Value:    amount,
				Symbol:   p.Coin().Symbol,
				Decimals: p.Coin().Decimals,
			},
			Memo: memo,
			Type: blockatlas.TxTransfer,
		}
		tx.Direction = tx.GetTransactionDirection(account)
		return tx, nil
	}
	return blockatlas.Tx{}, errors.New("Unknown action")
}

func unique(txs []blockatlas.Tx) []blockatlas.Tx {
	set := make(map[string]struct{})
	var result []blockatlas.Tx
	for _, tx := range txs {
		id := fmt.Sprintf("%s-%d", tx.ID, tx.Sequence)
		if _, ok := set[id]; ok {
			continue
		} else {
			set[id] = struct{}{}
			result = append(result, tx)
		}
	}
	return result
}
