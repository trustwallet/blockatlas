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
	srcBlock, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	block := NormalizeBlock(srcBlock)
	return &block, nil
}

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	srcTxs, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTxs(srcTxs), err
}

// NormalizeTx converts a Nimiq transaction into the generic model
func NormalizeTx(srcTx *Tx) blockatlas.Tx {
	date, _ := srcTx.Timestamp.Int64()
	return blockatlas.Tx{
		ID:    srcTx.Hash,
		Coin:  coin.NIM,
		Date:  date,
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
