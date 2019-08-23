package nebulas

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"math/big"
	"net/http"
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

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLatestIrreversibleBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	if block, err := p.client.GetBlockByNumber(num); err == nil {
		var normalizeTxs []blockatlas.Tx
		for _, srcTx := range block.TxnList {
			normalizeTxs = append(normalizeTxs, NormalizeNasTx(srcTx, block))
		}

		return &blockatlas.Block{
			Number: num,
			Txs:    normalizeTxs,
		}, nil
	} else {
		return nil, err
	}
}

func NormalizeTx(srcTx Transaction) blockatlas.Tx {
	var status string = blockatlas.StatusCompleted
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

func NormalizeNasTx(srcTx NasTransaction, block NasBlock) blockatlas.Tx {
	var status string = blockatlas.StatusCompleted
	if srcTx.Status == 0 {
		status = blockatlas.StatusFailed
	}
	//calculate fee
	fee := calcFee(srcTx.GasPrice, srcTx.GasUsed)

	return blockatlas.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.NAS,
		From:     srcTx.From,
		To:       srcTx.To,
		Fee:      blockatlas.Amount(fee),
		Date:     srcTx.Timestamp,
		Block:    block.Height,
		Status:   status,
		Sequence: srcTx.Nonce,
		Meta: blockatlas.Transfer{
			Value:    blockatlas.Amount(srcTx.Value),
			Symbol:   coin.Coins[coin.NAS].Symbol,
			Decimals: coin.Coins[coin.NAS].Decimals,
		},
	}
}

//How to calculate fees can be found here http://wiki.nebulas.io/en/latest/go-nebulas/design-overview/gas.html
func calcFee(gasPrice string, gasUsed string) string {
	var gasPriceBig, gasUsedBig, feeBig big.Int

	gasPriceBig.SetString(gasPrice, 10)
	gasUsedBig.SetString(gasUsed, 10)

	feeBig.Mul(&gasPriceBig, &gasUsedBig)

	return feeBig.String()
}
