package cosmos

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/util"
	"strconv"
	"sync"
	"time"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("cosmos.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ATOM]
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	txs := NormalizeTxs(srcTxs.Txs)
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	srcTxs := make([]Tx, 0)
	tagsList := []string{"transfer.recipient", "message.sender"}

	var wg sync.WaitGroup
	for _, t := range tagsList {
		wg.Add(1)
		go func(tag, addr string) {
			defer wg.Done()
			txs, _ := p.client.GetAddrTxs(addr, tag)
			srcTxs = append(srcTxs, txs.Txs...)
		}(t, address)
	}
	wg.Wait()
	return NormalizeTxs(srcTxs), nil
}

// NormalizeTxs converts multiple Cosmos transactions
func NormalizeTxs(srcTxs []Tx) blockatlas.TxPage {
	txMap := make(map[string]bool)
	txs := make(blockatlas.TxPage, 0)
	for _, srcTx := range srcTxs {
		_, ok := txMap[srcTx.ID]
		if ok {
			continue
		}
		normalisedInputTx, ok := Normalize(&srcTx)
		if ok {
			txMap[srcTx.ID] = true
			txs = append(txs, normalisedInputTx)
		}
	}
	return txs
}

// Normalize converts an Cosmos transaction into the generic model
func Normalize(srcTx *Tx) (tx blockatlas.Tx, ok bool) {
	date, err := time.Parse("2006-01-02T15:04:05Z", srcTx.Date)
	if err != nil {
		return blockatlas.Tx{}, false
	}
	block, err := strconv.ParseUint(srcTx.Block, 10, 64)
	if err != nil {
		return blockatlas.Tx{}, false
	}
	// Sometimes fees can be null objects (in the case of no fees e.g. F044F91441C460EDCD90E0063A65356676B7B20684D94C731CF4FAB204035B41)
	fee := "0"
	if len(srcTx.Data.Contents.Fee.FeeAmount) > 0 {
		fee, err = util.DecimalToSatoshis(srcTx.Data.Contents.Fee.FeeAmount[0].Quantity)
		if err != nil {
			return blockatlas.Tx{}, false
		}
	}

	tx = blockatlas.Tx{
		ID:     srcTx.ID,
		Coin:   coin.ATOM,
		Date:   date.Unix(),
		Status: blockatlas.StatusCompleted,
		Fee:    blockatlas.Amount(fee),
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
		fillTransfer(&tx, transfer)
		return tx, true
	case MessageValueDelegate:
		delegate := msg.Value.(MessageValueDelegate)
		fillDelegate(&tx, delegate, msg.Type)
		return tx, true
	}
	return tx, false
}

func fillTransfer(tx *blockatlas.Tx, transfer MessageValueTransfer) {
	if len(transfer.Amount) == 0 {
		return
	}
	value, err := util.DecimalToSatoshis(transfer.Amount[0].Quantity)
	if err != nil {
		return
	}
	tx.From = transfer.FromAddr
	tx.To = transfer.ToAddr
	tx.Type = blockatlas.TxTransfer

	tx.Meta = blockatlas.Transfer{
		Value:    blockatlas.Amount(value),
		Symbol:   coin.Coins[coin.ATOM].Symbol,
		Decimals: coin.Coins[coin.ATOM].Decimals,
	}
}

func fillDelegate(tx *blockatlas.Tx, delegate MessageValueDelegate, msgType string) {
	value, err := util.DecimalToSatoshis(delegate.Amount.Quantity)
	if err != nil {
		return
	}
	tx.From = delegate.DelegatorAddr
	tx.To = delegate.ValidatorAddr
	tx.Type = blockatlas.TxAnyAction

	title := blockatlas.KeyTitle("")
	switch msgType {
	case MsgDelegate:
		title = blockatlas.AnyActionDelegation
	case MsgUndelegate:
		title = blockatlas.AnyActionUndelegation
	}
	tx.Meta = blockatlas.AnyAction{
		Coin:     coin.ATOM,
		Title:    title,
		Key:      blockatlas.KeyStakeDelegate,
		Name:     "ATOM",
		Symbol:   coin.Coins[coin.ATOM].Symbol,
		Decimals: coin.Coins[coin.ATOM].Decimals,
		Value:    blockatlas.Amount(value),
	}
}
