package aion

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/util"
	"net/http"
	"strconv"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.BaseURL = viper.GetString("aion.api")
	p.client.HTTPClient = http.DefaultClient
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
func NormalizeTx(srcTx *Tx) blockatlas.Tx {
	fee := strconv.Itoa(srcTx.NrgConsumed)
	value := util.DecimalExp(string(srcTx.Value), 18)

	return blockatlas.Tx{
		ID:    srcTx.TransactionHash,
		Coin:  coin.AION,
		Date:  srcTx.BlockTimestamp,
		From:  "0x" + srcTx.FromAddr,
		To:    "0x" + srcTx.ToAddr,
		Fee:   blockatlas.Amount(fee),
		Block: srcTx.BlockNumber,
		Meta:  blockatlas.Transfer{
			Value: blockatlas.Amount(value),
		},
	}
}

// NormalizeTxs converts multiple Aion transactions
func NormalizeTxs(srcTxs []Tx) []blockatlas.Tx {
	txs := make([]blockatlas.Tx, len(srcTxs))
	for i, srcTx := range srcTxs {
		txs[i] = NormalizeTx(&srcTx)
	}
	return txs
}
