package nebulas

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"net/http"
	"strconv"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client.URL = viper.GetString("nebulas.api")
	p.client.HTTPClient = http.DefaultClient
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.NAS]
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	txs, err := p.client.GetTxs(address, 1)
	if err != nil {
		return nil, err
	}

	var normalizeTxs []blockatlas.Tx
	for _, srcTx := range txs {
		normalizeTxs = append(normalizeTxs, NormalizeTx(srcTx))
	}
	return normalizeTxs, nil
}

func NormalizeTx(srcTx Transaction) blockatlas.Tx {
	var status = blockatlas.StatusCompleted
	if srcTx.Status == 0 {
		status = blockatlas.StatusFailed
	}
	return blockatlas.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.NAS,
		From:     srcTx.From.Hash,
		To:       srcTx.To.Hash,
		Fee:      blockatlas.Amount(srcTx.TxFee),
		Date:     int64(srcTx.Timestamp) / 1000,
		Block:    srcTx.Block.Height,
		Status:   status,
		Sequence: srcTx.Nonce,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Value),
			Symbol:   coin.Coins[coin.NAS].Symbol,
			Decimals: coin.Coins[coin.NAS].Decimals,
		},
	}
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	result, err := p.client.GetLatestIrreversibleBlock()
	if err != nil {
		return 0, err
	}

	height, err := strconv.ParseInt(result.Height, 10, 64)
	if err != nil {
		return 0, err
	}
	return height, nil
}
