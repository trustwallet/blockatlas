package ripple

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strconv"
	"time"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("ripple.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.XRP]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	s, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}

	txs := make([]blockatlas.Tx, 0)
	for _, srcTx := range s {
		tx, ok := NormalizeTx(&srcTx)
		if !ok {
			continue
		}
		txs = append(txs, tx)
	}

	return txs, nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	if srcBlock, err := p.client.GetBlockByNumber(num); err == nil {
		txs := NormalizeTxs(srcBlock)
		return &blockatlas.Block{
			Number: num,
			Txs:    txs,
		}, nil
	} else {
		return nil, err
	}
}

func NormalizeTxs(srcTxs []Tx) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(&srcTx)
		if !ok || len(txs) >= blockatlas.TxPerPage {
			continue
		}
		txs = append(txs, tx)
	}
	return
}

// Normalize converts a Ripple transaction into the generic model
func NormalizeTx(srcTx *Tx) (blockatlas.Tx, bool) {

	date, err := time.Parse("2006-01-02T15:04:05-07:00", srcTx.Date)
	var unix int64
	if err != nil {
		unix = 0
	} else {
		unix = date.Unix()
	}

	v, vok := srcTx.Meta.DeliveredAmount.(string)
	if !vok || len(v) == 0 {
		return blockatlas.Tx{}, false
	}

	if srcTx.Payment.TransactionType != "Payment" {
		return blockatlas.Tx{}, false
	}

	result := blockatlas.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.XRP,
		Date:  unix,
		From:  srcTx.Payment.Account,
		To:    srcTx.Payment.Destination,
		Fee:   srcTx.Payment.Fee,
		Block: srcTx.LedgerIndex,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(v),
			Symbol:   coin.Coins[coin.XRP].Symbol,
			Decimals: coin.Coins[coin.XRP].Decimals,
		},
	}
	if srcTx.Payment.DestinationTag > 0 {
		result.Memo = strconv.FormatInt(srcTx.Payment.DestinationTag, 10)
	}
	return result, true
}
