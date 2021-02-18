package kava

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/trustwallet/golibs/numbers"
	"github.com/trustwallet/golibs/types"
)

const kavaDenom = "ukava"

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	return p.GetTokenTxsByAddress(address, kavaDenom)
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (types.Txs, error) {
	tagsList := []string{"transfer.recipient", "message.sender"}
	var wg sync.WaitGroup
	out := make(chan []Tx, len(tagsList))
	wg.Add(len(tagsList))
	for _, t := range tagsList {
		go func(tag, addr string, wg *sync.WaitGroup) {
			defer wg.Done()
			page := 1
			txs, err := p.client.GetAddrTxs(addr, tag, page)
			if err != nil {
				return
			}
			// Condition when no more pages to paginate
			if txs.PageTotal == "1" || txs.PageTotal == "0" {
				out <- txs.Txs
				return
			}

			totalPages, err := strconv.Atoi(txs.PageTotal)
			if err != nil {
				return
			}
			// gaia does support sort option, paginate to get latest transactions by passing total pages page
			// https://github.com/cosmos/gaia/blob/f61b391aee5d04364d2b5539692bbb187ad9b946/docs/resources/gaiacli.md#query-transactions
			txs2, err := p.client.GetAddrTxs(addr, tag, totalPages)
			if err != nil {
				return
			}
			out <- txs2.Txs
		}(t, address, &wg)
	}
	wg.Wait()
	close(out)
	srcTxs := make([]Tx, 0)
	for r := range out {
		filteredTxs := p.FilterTxsByDenom(r, token)
		srcTxs = append(srcTxs, filteredTxs...)
	}
	return p.NormalizeTxs(srcTxs), nil
}

func (p *Platform) FilterTxsByDenom(txs []Tx, denom string) []Tx {
	filteredTxs := make([]Tx, 0)
	for _, tx := range txs {
		messages := tx.Data.Contents.Message
		if len(messages) == 0 {
			continue
		}
		var amount Amount
		switch messages[0].Value.(type) {
		case MessageValueTransfer:
			amount = messages[0].Value.(MessageValueTransfer).Amount[0]
		case MessageValueDelegate:
			amount = messages[0].Value.(MessageValueDelegate).Amount
		}
		if amount.Denom == denom {
			filteredTxs = append(filteredTxs, tx)
		}
	}
	return filteredTxs
}

// NormalizeTxs converts multiple Cosmos transactions
func (p *Platform) NormalizeTxs(srcTxs []Tx) types.Txs {
	txMap := make(map[string]bool)
	txs := make(types.Txs, 0)
	for _, srcTx := range srcTxs {
		_, ok := txMap[srcTx.ID]
		if ok {
			continue
		}
		normalisedInputTx, ok := p.Normalize(&srcTx)
		if ok {
			txMap[srcTx.ID] = true
			txs = append(txs, normalisedInputTx)
		}
	}
	return txs
}

// Normalize converts an Cosmos transaction into the generic model
func (p *Platform) Normalize(srcTx *Tx) (tx types.Tx, ok bool) {
	date, err := time.Parse("2006-01-02T15:04:05Z", srcTx.Date)
	if err != nil {
		return types.Tx{}, false
	}
	block, err := strconv.ParseUint(srcTx.Block, 10, 64)
	if err != nil {
		return types.Tx{}, false
	}
	// Sometimes fees can be null objects (in the case of no fees e.g. F044F91441C460EDCD90E0063A65356676B7B20684D94C731CF4FAB204035B41)
	fee := "0"
	if len(srcTx.Data.Contents.Fee.FeeAmount) > 0 {
		qty := srcTx.Data.Contents.Fee.FeeAmount[0].Quantity
		if len(qty) > 0 && qty != fee {
			fee, err = numbers.DecimalToSatoshis(srcTx.Data.Contents.Fee.FeeAmount[0].Quantity)
			if err != nil {
				return types.Tx{}, false
			}
		}
	}

	status := types.StatusCompleted
	// https://github.com/cosmos/cosmos-sdk/blob/95ddc242ad024ca78a359a13122dade6f14fd676/types/errors/errors.go#L19
	if srcTx.Code > 0 {
		status = types.StatusError
	}

	tx = types.Tx{
		ID:     srcTx.ID,
		Coin:   p.Coin().ID,
		Date:   date.Unix(),
		Status: status,
		Fee:    types.Amount(fee),
		Block:  block,
		Memo:   srcTx.Data.Contents.Memo,
	}

	if len(srcTx.Data.Contents.Message) == 0 {
		return tx, false
	}

	msg := srcTx.Data.Contents.Message[0]
	switch msg.Value.(type) {
	case MessageValueTransfer:
		transfer := msg.Value.(MessageValueTransfer)
		p.fillTransfer(&tx, transfer)
		return tx, true
	case MessageValueDelegate:
		delegate := msg.Value.(MessageValueDelegate)
		p.fillDelegate(&tx, delegate, srcTx.Events, msg.Type)
		return tx, true
	}
	return tx, false
}

func (p *Platform) fillTransfer(tx *types.Tx, transfer MessageValueTransfer) {
	if len(transfer.Amount) == 0 {
		return
	}
	amount := transfer.Amount[0]
	value, err := numbers.DecimalToSatoshis(amount.Quantity)
	if err != nil {
		return
	}
	tx.From = transfer.FromAddr
	tx.To = transfer.ToAddr
	tx.Type = types.TxTransfer
	tx.Meta = types.Transfer{
		Value:    types.Amount(value),
		Symbol:   p.Coin().Symbol,
		Decimals: p.Coin().Decimals,
	}
	switch {
	case amount.Denom == kavaDenom:
		tx.Type = types.TxTransfer
		tx.Meta = types.Transfer{
			Value:    types.Amount(value),
			Symbol:   p.Coin().Symbol,
			Decimals: p.Coin().Decimals,
		}
	default:
		tx.Type = types.TxNativeTokenTransfer
		tx.Meta = types.NativeTokenTransfer{
			Decimals: p.Coin().Decimals,
			From:     tx.From,
			Symbol:   strings.ToUpper(amount.Denom),
			Name:     amount.Denom,
			To:       tx.To,
			TokenID:  amount.Denom,
			Value:    types.Amount(value),
		}
	}
}

func (p *Platform) fillDelegate(tx *types.Tx, delegate MessageValueDelegate, events Events, msgType TxType) {
	value := ""
	if len(delegate.Amount.Quantity) > 0 {
		var err error
		value, err = numbers.DecimalToSatoshis(delegate.Amount.Quantity)
		if err != nil {
			return
		}
	}
	tx.From = delegate.DelegatorAddr
	tx.To = delegate.ValidatorAddr
	tx.Type = types.TxAnyAction

	key := types.KeyStakeDelegate
	title := types.KeyTitle("")
	switch msgType {
	case MsgDelegate:
		tx.Direction = types.DirectionOutgoing
		title = types.AnyActionDelegation
	case MsgUndelegate:
		tx.Direction = types.DirectionIncoming
		title = types.AnyActionUndelegation
	case MsgWithdrawDelegationReward:
		tx.Direction = types.DirectionIncoming
		title = types.AnyActionClaimRewards
		key = types.KeyStakeClaimRewards
		value = events.GetWithdrawRewardValue()
	}
	tx.Meta = types.AnyAction{
		Coin:     p.Coin().ID,
		Title:    title,
		Key:      key,
		Name:     p.Coin().Name,
		Symbol:   p.Coin().Symbol,
		Decimals: p.Coin().Decimals,
		Value:    types.Amount(value),
	}
}
