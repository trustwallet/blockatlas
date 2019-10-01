package nebulas

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("nebulas.api"))}
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

	return NormalizeTxs(txs), nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLatestBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	txs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	return &blockatlas.Block{
		Number: num,
		Txs:    NormalizeTxs(txs),
	}, nil
}

func NormalizeTxs(txs []Transaction) []blockatlas.Tx {
	normalizeTxs := make([]blockatlas.Tx, 0)
	for _, srcTx := range txs {
		normalizeTxs = append(normalizeTxs, NormalizeTx(srcTx))
	}
	return normalizeTxs
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
