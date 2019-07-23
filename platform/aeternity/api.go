package aeternity

import (
	"net/http"
	"strconv"

	"github.com/trustwallet/blockatlas/coin"

	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.URL = viper.GetString("aeternity.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.AE]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	addressTxs, err := p.client.GetTxs(address, 25)
	if err != nil {
		return nil, err
	}

	var txs []blockatlas.Tx
	for _, srcTx := range addressTxs {
		txs = append(txs, NormalizeTx(&srcTx, p.Coin().ID))
	}

	return txs, nil
}

func NormalizeTx(srcTx *Transaction, coinIndex uint) blockatlas.Tx {
	return blockatlas.Tx{
		ID:     srcTx.Hash,
		Coin:   coinIndex,
		From:   srcTx.TxValue.Sender,
		To:     srcTx.TxValue.Recipient,
		Fee:    blockatlas.Amount(strconv.Itoa(int(srcTx.TxValue.Fee))),
		Date:   int64(srcTx.Timestamp) / 1000,
		Block:  srcTx.TxValue.BlockHeight,
		Memo:   "",
		Status: blockatlas.StatusCompleted,
		Meta: blockatlas.Transfer{
			Value: blockatlas.Amount(strconv.Itoa(int(srcTx.TxValue.Amount))),
		},
	}
}
