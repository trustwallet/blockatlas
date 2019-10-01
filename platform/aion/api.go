package aion

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/util"
	"strconv"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("aion.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.AION]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	if srcTxs, err := p.client.GetTxsOfAddress(address, blockatlas.TxPerPage); err == nil {
		return NormalizeTxs(srcTxs.Content), err
	} else {
		return nil, err
	}
}

// NormalizeTx converts an Aion transaction into the generic model
func NormalizeTx(srcTx *Tx) (tx blockatlas.Tx, ok bool) {
	fee := strconv.Itoa(srcTx.NrgConsumed)
	value := util.DecimalExp(string(srcTx.Value), 18)
	value, ok = util.CutZeroFractional(value)
	if !ok {
		return tx, false
	}

	return blockatlas.Tx{
		ID:    srcTx.TransactionHash,
		Coin:  coin.AION,
		Date:  srcTx.BlockTimestamp,
		From:  "0x" + srcTx.FromAddr,
		To:    "0x" + srcTx.ToAddr,
		Fee:   blockatlas.Amount(fee),
		Block: srcTx.BlockNumber,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(value),
			Symbol:   coin.Coins[coin.AION].Symbol,
			Decimals: coin.Coins[coin.AION].Decimals,
		},
	}, true
}

// NormalizeTxs converts multiple Aion transactions
func NormalizeTxs(srcTxs []Tx) []blockatlas.Tx {
	var txs []blockatlas.Tx
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(&srcTx)
		if ok {
			txs = append(txs, tx)
		}
	}
	return txs
}
