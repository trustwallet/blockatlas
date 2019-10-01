package nimiq

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("nimiq.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.NIM]
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	if srcBlock, err := p.client.GetBlockByNumber(num); err == nil {
		block := NormalizeBlock(srcBlock)
		return &block, nil
	} else {
		return nil, err
	}
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	if srcTxs, err := p.client.GetTxsOfAddress(address, blockatlas.TxPerPage); err == nil {
		return NormalizeTxs(srcTxs), err
	} else {
		return nil, err
	}
}

// NormalizeTx converts a Nimiq transaction into the generic model
func NormalizeTx(srcTx *Tx) blockatlas.Tx {
	return blockatlas.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.NIM,
		Date:  srcTx.Timestamp,
		From:  srcTx.FromAddress,
		To:    srcTx.ToAddress,
		Fee:   srcTx.Fee,
		Block: srcTx.BlockNumber,
		Meta: blockatlas.Transfer{
			Value:    srcTx.Value,
			Symbol:   coin.Coins[coin.NIM].Symbol,
			Decimals: coin.Coins[coin.NIM].Decimals,
		},
	}
}

// NormalizeTxs converts multiple Nimiq transactions
func NormalizeTxs(srcTxs []Tx) []blockatlas.Tx {
	txs := make([]blockatlas.Tx, len(srcTxs))
	for i, srcTx := range srcTxs {
		txs[i] = NormalizeTx(&srcTx)
	}
	return txs
}

// NormalizeBlock converts a Nimiq block into the generic model
func NormalizeBlock(srcBlock *Block) blockatlas.Block {
	return blockatlas.Block{
		Number: srcBlock.Number,
		ID:     srcBlock.Hash,
		Txs:    NormalizeTxs(srcBlock.Txs),
	}
}
